package response

import "github.com/Gym-Apps/gym-backend/models"

type UserLoginDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	AccountName string `json:"account_name" gorm:"-"`
	Token       string `json:"token"`
}

func (u *UserLoginDTO) Convert(user models.User) {
	u.ID = user.ID
	u.Name = user.Name
	u.Phone = user.Phone
	u.Surname = user.Surname
	u.Email = user.Email
	u.Gender = user.GenderName
	u.AccountName = user.AccountName
}
