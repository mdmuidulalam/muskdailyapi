package manager

import (
	accountViewModel "muskdaily.com/viewModel/account"
)

type SignUpManager interface {
	SignUp(signUpViewModel accountViewModel.SignUpViewModel) int
	Activate(accountViewModel accountViewModel.ActivationViewModel) int
}
