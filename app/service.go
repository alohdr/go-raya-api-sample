package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go-rest-api/db"
	"go-rest-api/redis"
	"go-rest-api/utils"
	"os"
	"time"
)

type (
	Service struct {
		repository Repository
	}
)

func NewService() Service {
	return Service{
		repository: Repository{
			db:  db.Postgres(),
			rdb: redis.NewRedisClient(os.Getenv("REDIS_HOST"), ""),
		},
	}
}

// USER
func (s Service) GetAllUser(ctx context.Context) (*UserResponses, error) {

	dest := UserResponses{}

	get, err := s.repository.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	for _, val := range get {
		dest = append(dest, UserResponse{
			UserID:      val.UserID,
			UserName:    val.UserName,
			UserPin:     val.UserPin,
			UserEmail:   val.UserEmail,
			UserPass:    val.UserPass,
			UserBalance: val.UserBalance,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			DeletedAt:   val.DeletedAt,
		})
	}

	return &dest, nil
}
func (s Service) InsertUser(ctx context.Context, req []byte) (*UserInsertResponse, error) {

	var user User
	err := json.Unmarshal(req, &user)
	if err != nil {
		return nil, err
	}

	if user.UserPin == "" {
		return nil, errors.New("insert pin")
	}

	id := uuid.New().String()

	CreatedUser := User{
		UserID:    id,
		UserName:  user.UserName,
		UserPin:   user.UserPin,
		UserEmail: user.UserEmail,
		UserPass:  user.UserPass,
	}

	err = s.repository.InsertUser(ctx, CreatedUser)
	if err != nil {
		return nil, err
	}

	return &UserInsertResponse{
		Status: "success",
		Result: ResultUser{
			UserID:   CreatedUser.UserID,
			UserName: CreatedUser.UserName,
		},
	}, nil
}
func (s Service) GetSaldo(ctx context.Context, userId string, userPin string) (*SaldoResponse, error) {

	get, err := s.repository.GetSaldo(ctx, userId, userPin)
	if err != nil {
		return nil, err
	}

	dest := SaldoResponse{
		UserID:      get.UserID,
		UserName:    get.UserName,
		UserBalance: get.UserBalance,
	}

	return &dest, nil

}

//BANK
func (s Service) GetAllBank(ctx context.Context, req string) (*BankResponses, error) {

	dest := BankResponses{}
	if req != "" {
		getName, err := s.repository.GetBankByName(ctx, req)
		if err != nil {
			return nil, err
		}

		for _, val := range getName {
			dest = append(dest, BankResponse{
				BankID:       val.BankID,
				BankCode:     val.BankCode,
				BankName:     val.BankName,
				BankAdminFee: val.BankAdminFee,
				BankIcon:     val.BankIcon,
				// CreatedAt:    val.CreatedAt.Format(utils.LayoutDateTime),
				CreatedAt: val.CreatedAt,
				UpdatedAt: val.UpdatedAt,
				DeletedAt: val.DeletedAt,
			})
		}

		return &dest, nil
	}

	key := "bank"
	ttl := time.Duration(30) * time.Second

	resp, errResp := s.repository.RedisGet(ctx, key)
	fmt.Println(resp)
	if errResp != nil {
		get, err := s.repository.GetAllBank(ctx)
		if err != nil {
			return nil, err
		}

		for _, val := range get {
			dest = append(dest, BankResponse{
				BankID:       val.BankID,
				BankCode:     val.BankCode,
				BankName:     val.BankName,
				BankAdminFee: val.BankAdminFee,
				BankIcon:     val.BankIcon,
				// CreatedAt:    val.CreatedAt.Format(utils.LayoutDateTime),
				CreatedAt: val.CreatedAt,
				UpdatedAt: val.UpdatedAt,
				DeletedAt: val.DeletedAt,
			})
		}

		errSetRedis := s.repository.RedisSet(ctx, key, &dest, ttl)
		if errSetRedis != nil {
			return nil, errors.New("CANT SET REDIS DB")
		}

		fmt.Println("masuk disini")

		return &dest, nil
	}

	for _, val := range resp {
		dest = append(dest, BankResponse{
			BankID:       val.BankID,
			BankCode:     val.BankCode,
			BankName:     val.BankName,
			BankAdminFee: val.BankAdminFee,
			BankIcon:     val.BankIcon,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
			DeletedAt:    val.DeletedAt,
		})
	}

	fmt.Println("dariredis")
	return &dest, nil

}
func (s Service) InsertBank(ctx context.Context, req BankCollect) (*BankCollect, error) {

	CreatedBank := BankCollect{}
	CreatedBank = append(CreatedBank, req...)

	err := s.repository.InsertBank(ctx, CreatedBank)
	if err != nil {
		return nil, err
	}

	return &CreatedBank, nil
}

//ACCOUNT
func (s Service) GetAllAccount(ctx context.Context) (*AccountResponses, error) {

	dest := AccountResponses{}

	get, err := s.repository.GetAllAccount(ctx)
	if err != nil {
		return nil, err
	}

	for _, val := range get {
		dest = append(dest, AccountResponse{
			AccountID:      val.AccountID,
			BankCode:       val.BankCode,
			AccountName:    val.AccountName,
			AccountBankNum: val.AccountBankNum,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
			DeletedAt:      val.DeletedAt,
		})
	}

	return &dest, nil
}
func (s Service) CheckAccountNumber(ctx context.Context, req string) (*CheckAccountNumberResponses, error) {
	dest := CheckAccountNumberResponses{}

	get, err := s.repository.CheckAccountNumber(ctx, req)
	if err != nil {
		return nil, err
	}

	for _, val := range get {
		dest = append(dest, CheckAccountNumberResponse{
			AccountID:      val.AccountID,
			AccountName:    val.AccountName,
			AccountBankNum: val.AccountBankNum,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
			DeletedAt:      val.DeletedAt,
		})
	}

	return &dest, nil
}

//TRANSACTION
func (s Service) GetLastTransaction(ctx context.Context, reqUser string, reqBank string) (*LastTransactionResponses, error) {
	dest := LastTransactionResponses{}

	if reqBank == "" {
		get, err := s.repository.GetLastTransaction(ctx, reqUser)
		if err != nil {
			return nil, err
		}

		for _, val := range get {
			dest = append(dest, LastTransactionResponse{
				TransactionID:  val.TransactionID,
				BankCode:       val.BankCode,
				BankName:       val.BankName,
				AccountName:    val.AccountName,
				AccountBankNum: val.AccountBankNum,
				CreatedAt:      val.CreatedAt,
			})
		}
		return &dest, nil
	}

	get, err := s.repository.GetLastTransactionByParams(ctx, reqUser, reqBank)
	if err != nil {
		return nil, err
	}

	for _, val := range get {
		dest = append(dest, LastTransactionResponse{
			TransactionID:  val.TransactionID,
			BankCode:       val.BankCode,
			BankName:       val.BankName,
			AccountName:    val.AccountName,
			AccountBankNum: val.AccountBankNum,
			CreatedAt:      val.CreatedAt,
		})
	}
	return &dest, nil
}
func (s Service) GetFavoriteTransaction(ctx context.Context, reqUser string) (*FavoriteTransactionResponses, error) {
	dest := FavoriteTransactionResponses{}

	get, err := s.repository.GetFavoriteTransaction(ctx, reqUser)
	if err != nil {
		return nil, err
	}

	for _, val := range get {
		dest = append(dest, FavoriteTransactionResponse{
			TransactionID:  val.TransactionID,
			BankCode:       val.BankCode,
			BankName:       val.BankName,
			AccountName:    val.AccountName,
			AccountBankNum: val.AccountBankNum,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
		})
	}

	return &dest, nil
}
func (s Service) GetTransactionDetail(ctx context.Context, id string) (*TransactionDetailResponse, error) {

	get, err := s.repository.GetTransactionDetail(ctx, id)
	if err != nil {
		return nil, err
	}

	dest := TransactionDetailResponse{
		TransactionID:       get.TransactionID,
		UserID:              get.UserID,
		UserName:            get.UserName,
		AccountID:           get.AccountID,
		AccountName:         get.AccountName,
		AccountBankNum:      get.AccountBankNum,
		BankID:              get.BankID,
		BankCode:            get.BankCode,
		BankIcon:            get.BankIcon,
		TransactionAdminFee: get.TransactionAdminFee,
		TransactionAmount:   get.TransactionAmount,
		TransactionType:     get.TransactionType,
		TransactionDesc:     get.TransactionDesc,
		TransactionStatus:   get.TransactionStatus,
		CreatedAt:           get.CreatedAt,
	}

	return &dest, nil

}
func (s Service) UpdateStatusTransaction(ctx context.Context, req []byte) (*TransactionDetailResponse, error) {
	var status TransactionStatus
	err := json.Unmarshal(req, &status)
	if err != nil {
		return nil, err
	}

	if status.TransactionID == "" {
		return nil, errors.New("transaction id unknown")
	}

	//id := uuid.New().String()

	UpdateTransactionStatus := TransactionStatus{
		TransactionID:     status.TransactionID,
		TransactionStatus: status.TransactionStatus,
	}

	err = s.repository.UpdateStatusTransaction(ctx, UpdateTransactionStatus)
	if err != nil {
		return nil, err
	}

	get, errGet := s.repository.GetTransactionDetail(ctx, status.TransactionID)
	if errGet != nil {
		return nil, errGet
	}

	resp := TransactionDetailResponse{
		TransactionID:       get.TransactionID,
		UserID:              get.UserID,
		UserName:            get.UserName,
		AccountID:           get.AccountID,
		AccountName:         get.AccountName,
		AccountBankNum:      get.AccountBankNum,
		BankID:              get.BankID,
		BankCode:            get.BankCode,
		BankIcon:            get.BankIcon,
		TransactionAdminFee: get.TransactionAdminFee,
		TransactionAmount:   get.TransactionAmount,
		TransactionType:     get.TransactionType,
		TransactionDesc:     get.TransactionDesc,
		TransactionStatus:   get.TransactionStatus,
		CreatedAt:           get.CreatedAt,
	}

	return &resp, nil
}
func (s Service) InsertTransaction(ctx context.Context, req []byte) (*TransactionDetailResponse, error) {
	var transaction Transaction
	err := json.Unmarshal(req, &transaction)
	if err != nil {
		return nil, err
	}
	if transaction.UserID == "" && transaction.AccountID == "" && transaction.TransactionAmount == 0 && transaction.TransactionType == "" {
		return nil, errors.New("complete your form")
	}

	id := uuid.New().String()

	createTransaction := Transaction{
		TransactionID:       id,
		UserID:              transaction.UserID,
		AccountID:           transaction.AccountID,
		BankID:              transaction.BankID,
		TransactionAdminFee: transaction.TransactionAdminFee,
		TransactionAmount:   transaction.TransactionAmount,
		TransactionType:     transaction.TransactionType,
		TransactionDesc:     transaction.TransactionDesc,
		IsFavorite:          transaction.IsFavorite,
	}

	err = s.repository.InsertTransaction(ctx, createTransaction)
	if err != nil {
		return nil, err
	}

	go utils.Publisher(id)

	get, errGet := s.repository.GetTransactionDetail(ctx, id)
	if errGet != nil {
		return nil, errGet
	}

	createdTransaction := TransactionDetailResponse{
		TransactionID:       get.TransactionID,
		UserID:              get.UserID,
		UserName:            get.UserName,
		AccountID:           get.AccountID,
		AccountName:         get.AccountName,
		AccountBankNum:      get.AccountBankNum,
		BankID:              get.BankID,
		BankCode:            get.BankCode,
		BankIcon:            get.BankIcon,
		TransactionAdminFee: get.TransactionAdminFee,
		TransactionAmount:   get.TransactionAmount,
		TransactionType:     get.TransactionType,
		TransactionDesc:     get.TransactionDesc,
		TransactionStatus:   get.TransactionStatus,
		CreatedAt:           get.CreatedAt,
	}

	return &createdTransaction, nil
}
