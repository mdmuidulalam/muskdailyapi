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

func (this Account) SignUp(c *gin.Context) {
	var signUpViewModel accountViewModel.SignUpViewModel
	c.ShouldBind(&signUpViewModel)

	if signUpViewModel.FirstName == "" || signUpViewModel.LastName == "" || signUpViewModel.Email == "" || signUpViewModel.Password == "" {
		c.JSON(400, gin.H{})
		return
	}

	c.JSON(this.SignUpManager.SignUp(signUpViewModel), gin.H{})
}

func (this Account) Activate(c *gin.Context) {
	var activationViewModel accountViewModel.ActivationViewModel
	c.ShouldBind(&activationViewModel)

	if activationViewModel.Email == "" || activationViewModel.ActivationCode == "" {
		c.JSON(400, gin.H{})
		return
	}

	c.JSON(this.SignUpManager.Activate(activationViewModel), gin.H{})
}
