package app

import "time"

type (
	BankResponse struct {
		BankID       string     `json:"bank_id"`
		BankCode     string     `json:"bank_code"`
		BankName     string     `json:"bank_name"`
		BankAdminFee int        `json:"bank_admin_fee"`
		BankIcon     string     `json:"bank_icon"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    *time.Time `json:"updated_at"`
		DeletedAt    *time.Time `json:"deleted_at"`
	}

	BankResponses []BankResponse

	UserResponse struct {
		UserID      string     `json:"user_id"`
		UserName    string     `json:"user_name"`
		UserPin     string     `json:"user_pin_bank"`
		UserEmail   string     `json:"user_email"`
		UserPass    string     `json:"user_passsword"`
		UserBalance int        `json:"user_balance"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
	}

	UserResponses []UserResponse

	ResultUser struct {
		UserID   string `json:"user_id"`
		UserName string `json:"user_name"`
	}

	UserInsertResponse struct {
		Status string `json:"status"`
		Result ResultUser
	}

	AccountResponse struct {
		AccountID      string     `json:"account_id"`
		BankCode       string     `json:"bank_code"`
		AccountName    string     `json:"account_name"`
		AccountBankNum string     `json:"account_bank_number"`
		CreatedAt      time.Time  `json:"created_at"`
		UpdatedAt      *time.Time `json:"updated_at"`
		DeletedAt      *time.Time `json:"deleted_at"`
	}

	AccountResponses []AccountResponse

	LastTransactionResponse struct {
		TransactionID  string    `json:"transaction_id"`
		BankCode       string    `json:"bank_code"`
		BankName       string    `json:"bank_name"`
		AccountName    string    `json:"account_name"`
		AccountBankNum string    `json:"account_bank_number"`
		CreatedAt      time.Time `json:"created_at"`
	}

	LastTransactionResponses []LastTransactionResponse

	FavoriteTransactionResponse struct {
		TransactionID  string     `json:"transaction_id"`
		BankCode       string     `json:"bank_code"`
		BankName       string     `json:"bank_name"`
		AccountName    string     `json:"account_name"`
		AccountBankNum string     `json:"account_bank_number"`
		CreatedAt      time.Time  `json:"created_at"`
		UpdatedAt      *time.Time `json:"updated_at"`
	}

	FavoriteTransactionResponses []FavoriteTransactionResponse

	CheckAccountNumberResponse struct {
		AccountID      string     `json:"account_id"`
		AccountName    string     `json:"account_name"`
		AccountBankNum string     `json:"account_bank_number"`
		CreatedAt      time.Time  `json:"created_at"`
		UpdatedAt      *time.Time `json:"updated_at"`
		DeletedAt      *time.Time `json:"deleted_at"`
	}

	CheckAccountNumberResponses []CheckAccountNumberResponse

	SaldoResponse struct {
		UserID      string `json:"user_id"`
		UserName    string `json:"user_name"`
		UserBalance int    `json:"user_balance"`
	}

	SaldoResponses []SaldoResponse

	TransactionDetailResponse struct {
		TransactionID       string    `json:"transaction_id"`
		UserID              string    `json:"user_id"`
		UserName            string    `json:"user_name"`
		AccountID           string    `json:"account_id"`
		AccountName         string    `json:"account_name"`
		AccountBankNum      string    `json:"account_bank_number"`
		BankID              string    `json:"bank_id"`
		BankCode            string    `json:"bank_code"`
		BankIcon            string    `json:"bank_icon"`
		TransactionAdminFee int       `json:"transaction_admin_fee"`
		TransactionAmount   int       `json:"transaction_amount"`
		TransactionType     string    `json:"transaction_type"`
		TransactionDesc     string    `json:"transaction_desc"`
		TransactionStatus   string    `json:"transaction_status"`
		CreatedAt           time.Time `json:"created_at"`
	}

	TransactionDetailResponses []TransactionDetailResponse
)
