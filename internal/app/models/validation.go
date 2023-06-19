package models

// LoginPasswordValidate - модель для валидации логина и пароля при регистрации.
type LoginPasswordValidate struct {
	Login    string `json:"login" validate:"required,alphanum,gte=1"` //не короче 1 символа, из цифр и английского алфавита
	Password string `json:"password" validate:"required,gte=1"`       //не короче 1 символа
}
