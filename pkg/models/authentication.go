package models

import (
	db "blog/pkg/sqlite"
	"strings"
)

type AuthRequestBody struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	Result string `json:"result"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
}

type GetUserInformation struct {
	Password string `gorm:"column:password"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
}

type ErrorResponse struct {
	Result string `json:"result"`
	Data   struct {
		ErrorInformation string `json:"error_information"`
		ErrorMessage     string `json:"error_message"`
	} `json:"data"`
}

func (a *AuthenticationResponse) AuthWithEmail(authRequestBody AuthRequestBody) {
	var getUserInformation GetUserInformation

	db.GetDb().
		Raw("select * from user where email = $1 and password = $2", authRequestBody.Email, authRequestBody.Password).
		Find(&getUserInformation)

	if strings.EqualFold(getUserInformation.Password, authRequestBody.Password) && strings.EqualFold(getUserInformation.Email, authRequestBody.Email) {
		a.Result = "success"
		a.Data.Token = "token"
	} else if strings.EqualFold(getUserInformation.Password, authRequestBody.Password) || strings.EqualFold(getUserInformation.Email, authRequestBody.Email) {
		a.Result = "unsuccess"
		a.Data.Token = "coudn't generate token"
	}

}

func (a *AuthenticationResponse) AuthWithPhone(authRequestBody AuthRequestBody) {
	var getUserInformation GetUserInformation

	db.GetDb().
		Raw("select * from user where phone = $1 and password = $2", authRequestBody.Phone, authRequestBody.Password).
		Find(&getUserInformation)

	if strings.EqualFold(getUserInformation.Password, authRequestBody.Password) && strings.EqualFold(getUserInformation.Phone, authRequestBody.Phone) {
		a.Result = "success"
		a.Data.Token = "token"
	} else if strings.EqualFold(getUserInformation.Password, authRequestBody.Password) || strings.EqualFold(getUserInformation.Phone, authRequestBody.Phone) {
		a.Result = "unsuccess"
		a.Data.Token = "coudn't generate token"
	}

}

func (e *ErrorResponse) EmptyPassword(authRequestBody AuthRequestBody) {
	e.Result = "unsuccess"
	e.Data.ErrorInformation = "User can not have empty password"
	e.Data.ErrorMessage = "empty_password"
}

func (e *ErrorResponse) EmptyLogin(authRequestBody AuthRequestBody) {
	e.Result = "unsuccess"
	e.Data.ErrorInformation = "User can not have empty phone or email"
	e.Data.ErrorMessage = "empty_login"
}
