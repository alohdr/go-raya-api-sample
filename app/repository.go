package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go-rest-api/utils"
	"time"
)

type (
	Repository struct {
		db  *sql.DB
		rdb *redis.Client
	}
)

//USER
func (r Repository) GetAllUser(ctx context.Context) ([]UserModel, error) {
	qsyntax := fmt.Sprintf(`SELECT * FROM %s`, utils.Table_Users)

	rows, err := r.db.QueryContext(ctx, qsyntax)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []UserModel

	//Fetch data to struct
	for rows.Next() {
		var rowsScan UserModel
		err := rows.Scan(&rowsScan.UserID, &rowsScan.UserName, &rowsScan.UserPin, &rowsScan.UserEmail, &rowsScan.UserPass, &rowsScan.UserBalance, &rowsScan.CreatedAt, &rowsScan.UpdatedAt, &rowsScan.DeletedAt)

		if err != nil {
			return nil, err
		}

		// Append for every next row
		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}
func (r Repository) InsertUser(ctx context.Context, req User) error {
	qsyntax := fmt.Sprintf(`INSERT INTO %s (user_id, user_name, user_pin_bank, user_email, user_password) VALUES ($1, $2, $3, $4, $5)`, utils.Table_Users)

	_, err := r.db.ExecContext(ctx, qsyntax,
		req.UserID,
		req.UserName,
		req.UserPin,
		req.UserEmail,
		req.UserPass)

	if err != nil {
		return err
	}

	return nil
}

//ACCOUNT
func (r Repository) GetAllAccount(ctx context.Context) ([]AccountModel, error) {
	qsyntax := fmt.Sprintf(`SELECT * FROM %s`, utils.Table_Accounts)

	rows, err := r.db.QueryContext(ctx, qsyntax)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []AccountModel

	for rows.Next() {
		var rowscan AccountModel
		err := rows.Scan(&rowscan.AccountID, &rowscan.BankCode, &rowscan.AccountName, &rowscan.AccountBankNum, &rowscan.CreatedAt, &rowscan.UpdatedAt, &rowscan.DeletedAt)

		if err != nil {
			return nil, err
		}

		rowsScanArr = append(rowsScanArr, rowscan)
	}

	return rowsScanArr, nil
}

//BANK
func (r Repository) GetAllBank(ctx context.Context) ([]BankModel, error) {
	qsyntax := fmt.Sprintf(`SELECT * FROM %s`, utils.Table_Bank)

	rows, err := r.db.QueryContext(ctx, qsyntax)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []BankModel

	//Fetch data to struct
	for rows.Next() {
		var rowsScan BankModel
		err := rows.Scan(&rowsScan.BankID, &rowsScan.BankCode, &rowsScan.BankName, &rowsScan.BankAdminFee, &rowsScan.BankIcon, &rowsScan.CreatedAt, &rowsScan.UpdatedAt, &rowsScan.DeletedAt)

		if err != nil {
			return nil, err
		}

		// Append for every next row
		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}
func (r Repository) GetBankByName(ctx context.Context, req string) ([]BankModel, error) {
	qsyntax := fmt.Sprintf(`SELECT
				bank_code,
				bank_name,
				bank_admin_fee,
				bank_icon,
				created_at ,
				updated_at
				FROM %s
				WHERE bank_code is not null AND bank_name ILIKE $1
				limit 10;`, utils.Table_Bank)

	req = fmt.Sprint("%" + req + "%")

	rows, err := r.db.QueryContext(ctx, qsyntax, req)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []BankModel

	//Fetch data to struct
	for rows.Next() {
		var rowsScan BankModel
		err := rows.Scan(&rowsScan.BankCode, &rowsScan.BankName, &rowsScan.BankAdminFee, &rowsScan.BankIcon, &rowsScan.CreatedAt, &rowsScan.UpdatedAt)

		if err != nil {
			return nil, err
		}

		// Append for every next row
		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}
func (r Repository) InsertBank(ctx context.Context, req BankCollect) error {
	qsyntax := fmt.Sprintf(`INSERT INTO %s (bank_code, bank_name, bank_admin_fee, bank_icon) VALUES ($1, $2, $3, $4)`, utils.Table_Bank)

	for _, v := range req {

		_, err := r.db.ExecContext(ctx, qsyntax, v.BankCode, v.BankName, v.BankAdminFee, v.BankIcon)

		if err != nil {
			return err
		}
	}

	return nil
}

//TRANSACTION
func (r Repository) GetLastTransaction(ctx context.Context, reqUser string) ([]LastTransactionModel, error) {
	qsyntax := fmt.Sprintf(`SELECT 
				distinct on (%s.account_id)
				transactions.transaction_id, 
				bank.bank_code, 
				bank.bank_name, 
				accounts.account_name, 
				accounts.account_bank_number, 
				transactions.created_at 
				FROM %s
				join %s
				on bank.accounts.account_id = bank.transactions.account_id
				join %s
				on bank.bank.bank_id = bank.transactions.bank_id 
				where user_id = $1
				order by bank.transactions.account_id, bank.transactions.created_at  desc
				limit 10`, utils.Table_Transactions, utils.Table_Transactions, utils.Table_Accounts, utils.Table_Bank)

	rows, err := r.db.QueryContext(ctx, qsyntax, reqUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []LastTransactionModel

	for rows.Next() {
		var rowsScan LastTransactionModel
		err := rows.Scan(&rowsScan.TransactionID, &rowsScan.BankCode, &rowsScan.BankName, &rowsScan.AccountName, &rowsScan.AccountBankNum, &rowsScan.CreatedAt)
		if err != nil {
			return nil, err
		}

		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}
func (r Repository) GetLastTransactionByParams(ctx context.Context, reqUser string, reqBank string) ([]LastTransactionModel, error) {
	qsyntax := fmt.Sprintf(`SELECT 
				distinct on (%s.account_id)
				transactions.transaction_id, 
				bank.bank_code, 
				bank.bank_name, 
				accounts.account_name, 
				accounts.account_bank_number, 
				transactions.created_at 
				FROM %s
				join %s
				on bank.accounts.account_id = bank.transactions.account_id
				join %s
				on bank.bank.bank_id = bank.transactions.bank_id 
				where user_id = $1 and (bank.bank_name ilike $2 or account_name ilike $2)
				order by bank.transactions.account_id, bank.transactions.created_at  desc
				limit 10`, utils.Table_Transactions, utils.Table_Transactions, utils.Table_Accounts, utils.Table_Bank)

	rows, err := r.db.QueryContext(ctx, qsyntax, reqUser, reqBank)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []LastTransactionModel

	for rows.Next() {
		var rowsScan LastTransactionModel
		err := rows.Scan(&rowsScan.TransactionID, &rowsScan.BankCode, &rowsScan.BankName, &rowsScan.AccountName, &rowsScan.AccountBankNum, &rowsScan.CreatedAt)
		if err != nil {
			return nil, err
		}

		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}
func (r Repository) GetFavoriteTransaction(ctx context.Context) ([]FavoriteTransactionModel, error) {
	qsyntax := fmt.Sprintf(`SELECT 
				transactions.transaction_id, 
				bank.bank_code, 
				bank.bank_name, 
				accounts.account_name, 
				accounts.account_bank_number, 
				transactions.created_at, 
				transactions.updated_at 
				FROM %s
				join %s
				on bank.accounts.account_id = bank.transactions.account_id
				join %s
				on bank.bank.bank_id = bank.transactions.bank_id 
				where bank.transactions.is_favorite = true
				order by created_at
				limit 10`, utils.Table_Transactions, utils.Table_Accounts, utils.Table_Bank)

	rows, err := r.db.QueryContext(ctx, qsyntax)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []FavoriteTransactionModel

	for rows.Next() {
		var rowsScan FavoriteTransactionModel
		err := rows.Scan(&rowsScan.TransactionID, &rowsScan.BankCode, &rowsScan.BankName, &rowsScan.AccountName, &rowsScan.AccountBankNum, &rowsScan.CreatedAt, &rowsScan.UpdatedAt)
		if err != nil {
			return nil, err
		}

		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}
func (r Repository) CheckAccountNumber(ctx context.Context, req string) ([]CheckAccountNumberModel, error) {
	qsyntax := fmt.Sprintf(`select 
				account_id, 
				account_name, 
				account_bank_number, 
				created_at, 
				updated_at, 
				deleted_at 
				FROM %s
				WHERE deleted_at is null AND account_bank_number ILIKE $1 or account_name ilike $1 
				limit 10;`, utils.Table_Accounts)

	req = fmt.Sprint("%" + req + "%")
	rows, err := r.db.QueryContext(ctx, qsyntax, req)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []CheckAccountNumberModel

	for rows.Next() {
		var rowsScan CheckAccountNumberModel
		err := rows.Scan(&rowsScan.AccountID, &rowsScan.AccountName, &rowsScan.AccountBankNum, &rowsScan.CreatedAt, &rowsScan.UpdatedAt, &rowsScan.DeletedAt)
		if err != nil {
			return nil, err
		}

		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}
func (r Repository) GetSaldo(ctx context.Context, userId string, userPin string) (*SaldoModel, error) {
	qsyntax := fmt.Sprintf(`SELECT 
				user_id, user_name, user_balance 
				FROM %s
				where users.user_id = $1 and user_pin_bank = $2`, utils.Table_Users)

	rows, err := r.db.QueryContext(ctx, qsyntax, userId, userPin)
	if err != nil {
		return nil, err
	}

	var rowsResult SaldoModel

	for rows.Next() {
		var rowsScan SaldoModel
		err := rows.Scan(&rowsScan.UserID, &rowsScan.UserName, &rowsScan.UserBalance)
		if err != nil {
			return nil, err
		}

		rowsResult = rowsScan
	}

	return &rowsResult, nil
}
func (r Repository) GetTransactionDetail(ctx context.Context, id string) (*TransactionDetailModel, error) {
	qsyntax := fmt.Sprintf(`SELECT 
				transactions.transaction_id, 
				users.user_id, 
				users.user_name, 
				accounts.account_id, 
				accounts.account_name, 
				accounts.account_bank_number, 
				bank.bank_id, 
				bank.bank_code,
				bank.bank_icon,
				transaction_admin_fee, 
				transaction_amount, 
				transaction_type, 
				transaction_desc, 
				transaction_status,
				transactions.created_at
				FROM %s
				join %s
				on users.user_id = transactions.user_id
				join %s
				on accounts.account_id = transactions.account_id
				join %s
				on bank.bank_id = transactions.bank_id
				where transaction_id = $1;`, utils.Table_Transactions, utils.Table_Users, utils.Table_Accounts, utils.Table_Bank)

	rows, err := r.db.QueryContext(ctx, qsyntax, id)
	if err != nil {
		return nil, err
	}

	var rowsResult TransactionDetailModel

	for rows.Next() {
		var rowsScan TransactionDetailModel
		err := rows.Scan(&rowsScan.TransactionID, &rowsScan.UserID, &rowsScan.UserName, &rowsScan.AccountID, &rowsScan.AccountName, &rowsScan.AccountBankNum, &rowsScan.BankID, &rowsScan.BankCode, &rowsScan.BankIcon, &rowsScan.TransactionAdminFee, &rowsScan.TransactionAmount, &rowsScan.TransactionType, &rowsScan.TransactionDesc, &rowsScan.TransactionStatus, &rowsScan.CreatedAt)
		if err != nil {
			return nil, err
		}

		rowsResult = rowsScan
	}

	return &rowsResult, nil

}
func (r Repository) UpdateStatusTransaction(ctx context.Context, req TransactionStatus) error {
	qsyntax := fmt.Sprintf(`UPDATE %s
				SET transaction_status=$1::statustr, updated_at=now()
				WHERE transaction_id=$2;`, utils.Table_Transactions)

	_, err := r.db.ExecContext(ctx, qsyntax, req.TransactionStatus, req.TransactionID)
	if err != nil {
		return err
	}

	return nil

}
func (r Repository) InsertTransaction(ctx context.Context, req Transaction) error {
	qsyntax := fmt.Sprintf(`INSERT INTO %s
				(transaction_id, user_id, account_id, bank_id, transaction_admin_fee, transaction_amount, transaction_type, transaction_desc, transaction_status, is_favorite, created_at)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, 'pending'::statustr, $9, now());
				`, utils.Table_Transactions)

	_, err := r.db.ExecContext(ctx, qsyntax,
		req.TransactionID,
		req.UserID,
		req.AccountID,
		req.BankID,
		req.TransactionAdminFee,
		req.TransactionAmount,
		req.TransactionType,
		req.TransactionDesc,
		req.IsFavorite)

	if err != nil {
		return err
	}

	return nil
}
func (r Repository) RedisSet(ctx context.Context, key string, data interface{}, ttl time.Duration) error {
	val, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return err
	}

	r.rdb.Set(ctx, key, string(val), ttl)
	return nil
}
func (r Repository) RedisGet(ctx context.Context, key string) (BankResponses, error) {
	dest := BankResponses{}
	res, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	errUnmarshal := json.Unmarshal([]byte(res), &dest)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	fmt.Println("ini get redis di repo")
	return dest, nil
}
