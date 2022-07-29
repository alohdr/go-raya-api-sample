package app

type (
	User struct {
		UserID    string `json:"user_id" form:"user_id, omitempty"`
		UserName  string `json:"user_name" form:"user_name"`
		UserPin   string `json:"user_pin_bank" form:"user_pin_bank"`
		UserEmail string `json:"user_email" form:"user_email"`
		UserPass  string `json:"user_password" form:"user_passsword"`
	}

	Bank struct {
		BankCode     string `json:"bank_code" form:"bank_code"`
		BankName     string `json:"bank_name" form:"bank_name"`
		BankAdminFee int    `json:"bank_admin_fee" form:"bank_admin_fee, omitempty"`
		BankIcon     string `json:"bank_icon" form:"bank_icon, omitempty"`
	}

	BankCollect []Bank

	TransactionStatus struct {
		TransactionID     string `json:"transaction_id" form:"transaction_id"`
		TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	}

	Transaction struct {
		TransactionID       string `json:"transaction_id" form:"transaction_id omitempty"`
		UserID              string `json:"user_id" form:"user_id"`
		AccountID           string `json:"account_id" form:"account_id"`
		BankID              string `json:"bank_id" form:"bank_id"`
		TransactionAdminFee int    `json:"transaction_admin_fee" form:"transaction_admin_fee omitempty"`
		TransactionAmount   int    `json:"transaction_amount" form:"transaction_amount"`
		TransactionType     string `json:"transaction_type" form:"transaction_type"`
		TransactionDesc     string `json:"transaction_desc" form:"transaction_desc omitempty"`
		IsFavorite          bool   `json:"is_favorite" form:"is_favorite omitempty"`
	}
)
