package manager

import (
	bcrypt "golang.org/x/crypto/bcrypt"

	data "muskdaily.com/data/abstractions"
	model "muskdaily.com/model"
	viewModel "muskdaily.com/viewModel/account"
)

type AccountManager struct {
	AccountData data.AccountData
}

func (this AccountManager) SignUp(signUpViewModel viewModel.SignUpViewModel) (bool, string) {

	password := []byte(signUpViewModel.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	this.AccountData.Connect()
	this.AccountData.AddAccount(model.Account{
		FirstName:      signUpViewModel.FirstName,
		LastName:       signUpViewModel.LastName,
		Email:          signUpViewModel.Email,
		HashedPassword: hashedPassword,
		Active:         signUpViewModel.Active,
	})
	this.AccountData.Disconnect()

	return true, "User is registered"
}
