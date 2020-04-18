package routes

import (
	gin "github.com/gin-gonic/gin"

	manager "muskdaily.com/manager/abstractions"
	accountViewModel "muskdaily.com/viewModel/account"
)

type Account struct {
	Base
	SignUpManager manager.SignUpManager
}

func (this Account) New() {
	this.R.POST("signup", this.SignUp)
	this.R.POST("activate", this.Activate)
}

// @Summary Signup an account
// @Description It will register new users if the post data provided properly.
// @ID signup-new-account
// @Router /account/signup [post]
// @Accept  json
// @Param firstName body string true "FirstName"
// @Param lastname 	body string true "LastName"
// @Param email 	body string true "Email"
// @Param password 	body string true "Password"
// @Success  201	"Created a new account. Not active yet, an activation email will be sent to the given email."
// @Failure  409	"Account was created before but not activated. A new activation email will be sent to the given email."
// @Failure  406	"An account already exists with the given email."
// @Failure  400	"The server cannot or will not process the request due to an apparent client error."
func (this Account) SignUp(c *gin.Context) {
	var signUpViewModel accountViewModel.SignUpViewModel
	c.ShouldBind(&signUpViewModel)

	if signUpViewModel.FirstName == "" || signUpViewModel.LastName == "" || signUpViewModel.Email == "" || signUpViewModel.Password == "" {
		c.Writer.WriteHeader(400)
		return
	}

	c.Writer.WriteHeader(this.SignUpManager.SignUp(signUpViewModel))
}

// @Summary Activate an account
// @Description It will activate the account if proper code is provided. The code will be sent to user through email.
// @ID activate-new-account
// @Router /account/activate [post]
// @Accept  json
// @Param Email 			body string true "Email"
// @Param ActivationCode	body string true "ActivationCode"
// @Success  200	"The account is activated."
// @Failure  400	"The server cannot or will not process the request due to an apparent client error."
// @Failure  401	"Wrong activation code is given."
// @Failure  406	"The account is already activated."
// @Failure  409	"The account is not registered yet."
func (this Account) Activate(c *gin.Context) {
	var activationViewModel accountViewModel.ActivationViewModel
	c.ShouldBind(&activationViewModel)

	if activationViewModel.Email == "" || activationViewModel.ActivationCode == "" {
		c.Writer.WriteHeader(400)
		return
	}

	c.Writer.WriteHeader(this.SignUpManager.Activate(activationViewModel))
}
