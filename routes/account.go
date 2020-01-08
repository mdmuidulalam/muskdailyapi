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
}

func (this Account) SignUp(c *gin.Context) {
	var signUpViewModel accountViewModel.SignUpViewModel
	c.ShouldBind(&signUpViewModel)

	this.SignUpManager.SignUp(signUpViewModel)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
