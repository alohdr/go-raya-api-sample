package auth

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go-rest-api/utils"
)

type (
	Repository struct {
		db  *sql.DB
		rdb *redis.Client
	}
)

func (r Repository) GetUserByIdentity(ctx context.Context, identity string) (*UserModel, error) {
	qsyntax := fmt.Sprintf(`SELECT 
			user_id, 
			user_name, 
			user_pin_bank, 
			user_email, 
			user_password, 
			user_balance, 
			created_at, 
			updated_at, 
			deleted_at
			FROM %s
			where user_name = $1 or user_email = $1;`, utils.Table_Users)

	rows, err := r.db.QueryContext(ctx, qsyntax, identity)
	if err != nil {
		return nil, err
	}

	var rowsResult UserModel

	for rows.Next() {
		var rowsScan UserModel
		err = rows.Scan(&rowsScan.UserID, &rowsScan.UserName, &rowsScan.UserPinkBank, &rowsScan.UserEmail, &rowsScan.UserPass, &rowsScan.UserBalance, &rowsScan.CreatedAt, &rowsScan.UpdatedAt, &rowsScan.DeletedAt)
		if err != nil {
			return nil, err
		}

		rowsResult = rowsScan
	}

	return &rowsResult, nil

}
