package models

type TextData struct {
	Description string
	Data        string
}

type LoginPasswordData struct {
	Login       string
	Password    string
	Description string
}

type BinaryData struct {
	Description string
	Data        []byte
	HexData     string
}

type BankCardData struct {
	Description string
	Number      string
	ValidThru   string
	CVV         string
}

type Data[T DataType] struct {
	Data T
}

type DataType interface {
	LoginPasswordData | TextData | BinaryData | BankCardData
}
