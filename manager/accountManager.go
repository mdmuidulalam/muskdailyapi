package manager

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	bcrypt "golang.org/x/crypto/bcrypt"

	data "muskdaily.com/data/abstractions"
	model "muskdaily.com/model"
	emailService "muskdaily.com/services/emailService"
	randomService "muskdaily.com/services/randomService"
	accountViewodel "muskdaily.com/viewModel/account"
)

type AccountManager struct {
	AccountData data.AccountData
}

func (this AccountManager) SignUp(signUpViewModel accountViewodel.SignUpViewModel) int {
	this.AccountData.Connect()
	defer this.AccountData.Disconnect()

	accounts := this.AccountData.SelectAccounts(bson.D{{"email", signUpViewModel.Email}})
	if len(accounts) == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpViewModel.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		activationCode := randomService.RandString(randomService.RandIntRange(5, 30))
		hashedActivationCode, err := bcrypt.GenerateFromPassword([]byte(activationCode), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		this.AccountData.InsertAccount(model.Account{
			FirstName:      signUpViewModel.FirstName,
			LastName:       signUpViewModel.LastName,
			Email:          signUpViewModel.Email,
			HashedPassword: hashedPassword,
			Active:         false,
			CreatedOn:      time.Now().UTC(),
			ModifiedOn:     nil,
			HashedCode:     hashedActivationCode,
			CodeSentTime:   time.Now().UTC(),
		})

		go emailService.SendMail([]string{signUpViewModel.Email}, "",
			"Musk Daily Account Activation Email",
			"Hello from Musk Daily,\r\nYour activation link for Musk Daily is: "+
				activationCode+"\r\nPlease, clcik on the link to activate link to activate you account.\r\n\r\nIf you didn't register for Musk Daily ignore this email.")

		return 201
	} else if !accounts[0].Active {
		activationCode := randomService.RandString(randomService.RandIntRange(5, 30))
		hashedActivationCode, err := bcrypt.GenerateFromPassword([]byte(activationCode), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		this.AccountData.UpdateAccounts(bson.D{{"email", signUpViewModel.Email}}, bson.D{{"$set", bson.D{{"hashedcode", hashedActivationCode}, {"codesenttime", time.Now().UTC()}}}})

		go emailService.SendMail([]string{signUpViewModel.Email}, "",
			"Musk Daily Account Activation Email",
			"Hello from Musk Daily,\r\n"+
				"You have already created an account for Musk Daily, but didn't activate the account.\r\n"+
				"Your new activation link for Musk Daily is: "+
				activationCode+"\r\nPlease, clcik on the link to activate link to activate you account.\r\n\r\nIf you didn't register for Musk Daily ignore this email.")

		return 409
	} else {
		return 406
	}
}

func (this AccountManager) Activate(accountViewModel accountViewodel.ActivationViewModel) int {
	this.AccountData.Connect()
	defer this.AccountData.Disconnect()
	accountFilter := bson.D{{"email", accountViewModel.Email}}
	accounts := this.AccountData.SelectAccounts(accountFilter)
	if len(accounts) != 0 && !accounts[0].Active {
		account := accounts[0]
		err := bcrypt.CompareHashAndPassword(account.HashedCode, []byte(accountViewModel.ActivationCode))
		if err == nil {
			this.AccountData.UpdateAccounts(accountFilter, bson.D{{"$set", bson.D{{"active", true}, {"hashedcode", nil}, {"codesenttime", nil}}}})
			return 200
		} else {
			return 401
		}
	} else if accounts[0].Active {
		return 409
	} else {
		return 406
	}
}
