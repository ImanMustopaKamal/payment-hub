package contact

type ContactDto struct {
	ID         string `dynamodbav:"id"`
	Name       string `dynamodbav:"contact_name"`
	CardNumber int16  `dynamodbav:"contact_number"`
	BankName   string `dynamodbav:"contact_bank_name"`
	// Add more fields as needed
}
