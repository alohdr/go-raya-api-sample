package app

import (
	"fmt"
	"strings"
	"time"
)

type (
	CustomTime struct {
		time.Time
	}

	BankModel struct {
		BankID       string     `db:"bank_id"`
		BankCode     string     `db:"bank_code"`
		BankName     string     `db:"bank_name"`
		BankAdminFee int        `db:"bank_admin_fee"`
		BankIcon     string     `db:"bank_icon"`
		CreatedAt    time.Time  `db:"created_at"`
		UpdatedAt    *time.Time `db:"updated_at"`
		DeletedAt    *time.Time `db:"deleted_at"`
	}

	UserModel struct {
		UserID      string     `db:"user_id"`
		UserName    string     `db:"user_name"`
		UserPin     string     `db:"user_pin_bank"`
		UserEmail   string     `db:"user_email"`
		UserPass    string     `db:"user_passsword"`
		UserBalance int        `db:"user_balance"`
		CreatedAt   time.Time  `db:"created_at"`
		UpdatedAt   *time.Time `db:"updated_at"`
		DeletedAt   *time.Time `db:"deleted_at"`
	}

	AccountModel struct {
		AccountID      string     `db:"account_id"`
		BankCode       string     `db:"bank_code"`
		AccountName    string     `db:"account_name"`
		AccountBankNum string     `db:"account_bank_number"`
		CreatedAt      time.Time  `db:"created_at"`
		UpdatedAt      *time.Time `db:"updated_at"`
		DeletedAt      *time.Time `db:"deleted_at"`
	}

	CheckAccountNumberModel struct {
		AccountID      string     `db:"account_id"`
		AccountName    string     `db:"account_name"`
		AccountBankNum string     `db:"account_bank_number"`
		CreatedAt      time.Time  `db:"created_at"`
		UpdatedAt      *time.Time `db:"updated_at"`
		DeletedAt      *time.Time `db:"deleted_at"`
	}

	LastTransactionModel struct {
		TransactionID  string    `db:"transaction_id"`
		BankCode       string    `db:"bank_code"`
		BankName       string    `db:"bank_name"`
		AccountName    string    `db:"account_name"`
		AccountBankNum string    `db:"account_bank_number"`
		CreatedAt      time.Time `db:"created_at"`
	}

	FavoriteTransactionModel struct {
		TransactionID  string     `db:"transaction_id"`
		BankCode       string     `db:"bank_code"`
		BankName       string     `db:"bank_name"`
		AccountName    string     `db:"account_name"`
		AccountBankNum string     `db:"account_bank_number"`
		CreatedAt      time.Time  `db:"created_at"`
		UpdatedAt      *time.Time `db:"updated_at"`
	}

	SaldoModel struct {
		UserID      string `db:"user_id"`
		UserName    string `db:"user_name"`
		UserBalance int    `db:"user_balance"`
	}

	TransactionDetailModel struct {
		TransactionID       string    `db:"transaction_id"`
		UserID              string    `db:"user_id"`
		UserName            string    `db:"user_name"`
		AccountID           string    `db:"account_id"`
		AccountName         string    `db:"account_name"`
		AccountBankNum      string    `db:"account_bank_number"`
		BankID              string    `db:"bank_id"`
		BankCode            string    `db:"bank_code"`
		BankIcon            string    `db:"bank_icon"`
		TransactionAdminFee int       `db:"transaction_admin_fee"`
		TransactionAmount   int       `db:"transaction_amount"`
		TransactionType     string    `db:"transaction_type"`
		TransactionDesc     string    `db:"transaction_desc"`
		TransactionStatus   string    `db:"transaction_status"`
		CreatedAt           time.Time `db:"created_at"`
	}
)

const ctLayout = "01-02-2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}
