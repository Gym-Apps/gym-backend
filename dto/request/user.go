package request

type UserLoginDTO struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterDTO struct {
	Name        string `json:"name" validate:"required"`
	Surname     string `json:"surname" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Birthday    string `json:"birthday" validate:"required,datetime=02.01.2006"`
	AccountType uint8  `json:"account_type" validate:"required"`
	Gender      uint8  `json:"gender" validate:"required"`
}

type UserResetPasswordDTO struct {
	NewPassword string `json:"new_password" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
}
