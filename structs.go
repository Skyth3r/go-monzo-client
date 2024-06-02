package main

type AccountsResp struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID                string         `json:"id"`
	Closed            bool           `json:"closed"`
	Created           string         `json:"created"`
	Description       string         `json:"description"`
	Type              string         `json:"type"`
	OwnerType         string         `json:"owner_type"`
	IsFlex            bool           `json:"is_flex"`
	Currency          string         `json:"currency"`
	LegalEntity       string         `json:"legal_entity"`
	CountryCode       string         `json:"country_code"`
	CountryCodeAlpha3 string         `json:"country_code_alpha3"`
	Owners            []Owner        `json:"owners"`
	LinkedAccounts    []string       `json:"linked_accounts"`
	BusinessID        string         `json:"business_id"`
	AccountNumber     string         `json:"account_number"`
	SortCode          string         `json:"sort_code"`
	PaymentDetails    PaymentDetails `json:"payment_details"`
}

type Owner struct {
	UserID             string `json:"user_id"`
	PreferredName      string `json:"preferred_name"`
	PreferredFirstName string `json:"preferred_first_name"`
}

type PaymentDetails struct {
	UK   PaymentDetailsUK   `json:"locale_uk"`
	IBAN PaymentDetailsIBAN `json:"iban"`
}

type PaymentDetailsUK struct {
	AccountNumber string `json:"account_number"`
	SortCode      string `json:"sort_code"`
}

type PaymentDetailsIBAN struct {
	Unformatted string `json:"unformatted"`
	Formatted   string `json:"formatted"`
	BIC         string `json:"bic"`
}

type Balance struct {
	Balance      int64  `json:"balance"`
	TotalBalance int64  `json:"total_balance"`
	Currency     string `json:"currency"`
	SpendToday   int64  `json:"spend_today"`
}

type Pots struct {
	Pots []Pot `json:"pots"`
}

type Pot struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Style    string `json:"style"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	Deleted  bool   `json:"deleted"`
}
