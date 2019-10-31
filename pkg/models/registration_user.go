package models

import (
	db "blog/pkg/sqlite"
	"strings"
	"time"

	"github.com/jinzhu/copier"
)

type RegistrationRequestBody struct {
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	SecretQuestion string `json:"secret_question"`
	SecretAnswer   string `json:"secret_answer"`
}

type UserRegistrationResponse struct {
	Result string `json:"result"`
}

type User struct {
	Email          string `gorm:"column:email"`
	Password       string `gorm:"column:password"`
	Phone          string `gorm:"column:phone"`
	SecretQuestion string `gorm:"column:secret_question"`
	SecretAnswer   string `gorm:"column:secret_answer"`
	CreatedAt      string `gorm:"column:created_at"`
	UpdatedAt      string `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}

func (u *UserRegistrationResponse) RegistrationUser(registrationRequestBody RegistrationRequestBody) {
	var checkEmail GetUserInformation
	var checkPhone GetUserInformation

	db.GetDb().
		Raw("select * from user where email = $1", registrationRequestBody.Email).
		Find(&checkEmail)

	db.GetDb().
		Raw("select * from user where phone = $1", registrationRequestBody.Phone).
		Find(&checkPhone)

	if strings.EqualFold(checkEmail.Email, registrationRequestBody.Email) || strings.EqualFold(checkPhone.Phone, registrationRequestBody.Phone) {
		u.Result = "Your email or phone was registered earlier"
	} else {
		var user User
		copier.Copy(&user, &registrationRequestBody)
		t := time.Now().UTC().String()
		time := strings.Split(t, " +0000 UTC")

		user.CreatedAt = time[0]
		user.UpdatedAt = time[0]

		db.GetDb().
			Create(&user)

		var getUserInformation GetUserInformation
		db.GetDb().
			Raw("select * from user where email = $1", registrationRequestBody.Email).
			Find(&getUserInformation)

		u.Result = "success"

		if !(strings.EqualFold(getUserInformation.Email, registrationRequestBody.Email) && strings.EqualFold(getUserInformation.Password, registrationRequestBody.Password)) {
			u.Result = "insert wasn't success"
		}

	}

}
