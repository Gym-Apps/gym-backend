package request

type UserLoginDTO struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResetPasswordDTO struct {
	NewPassword string `json:"new_password" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
}
