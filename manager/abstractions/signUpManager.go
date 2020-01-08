package manager

import (
	viewModel "muskdaily.com/viewModel/account"
)

type SignUpManager interface {
	SignUp(signUpViewModel viewModel.SignUpViewModel) (bool, string)
}
