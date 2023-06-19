package models

// TextData - модель хранения текстовых данных.
type TextData struct {
	Description  string `redis:"description"`
	Data         string `redis:"data"`
	LastModified int64  `redis:"last_modified"`
}

// LoginPasswordData - модель хранения данных логина и пароля.
type LoginPasswordData struct {
	Login        string `redis:"login"`
	Password     string `redis:"password"`
	Description  string `redis:"description"`
	LastModified int64  `redis:"last_modified"`
}

// BinaryData - модель храниения бинарных данных.
type BinaryData struct {
	Description  string `redis:"description"`
	Data         []byte `redis:"data"`
	LastModified int64  `redis:"last_modified"`
}

// BankCardData - модель хранения банковских данных.
type BankCardData struct {
	Description  string `redis:"description"`
	Number       string `redis:"number"`
	ValidThru    string `redis:"valid_thru"`
	CVV          string `redis:"cvv"`
	LastModified int64  `redis:"last_modified"`
}
