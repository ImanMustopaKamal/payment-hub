package entities

type Contact struct {
	ID         string `dynamodbav:"contact_id"`
	AccountID  string `dynamodbav:"account_id"`
	Name       string `dynamodbav:"contact_name"`
	CardNumber int32  `dynamodbav:"contact_number"`
	BankName   string `dynamodbav:"contact_bank_name"`
}

type ContactCreateDto struct {
	AccountID  string `json:"account_id" validate:"required"`
	Name       string `json:"contact_name"`
	CardNumber int32  `json:"contact_number"`
	BankName   string `json:"contact_bank_name"`
}
