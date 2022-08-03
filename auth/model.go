package auth

import "time"

type (
	UserModel struct {
		UserID       string     `db:"user_id"`
		UserName     string     `db:"user_name"`
		UserPinkBank string     `db:"user_pin_bank"`
		UserEmail    string     `db:"user_email"`
		UserPass     string     `db:"user_passsword"`
		UserBalance  int        `db:"user_balance"`
		CreatedAt    time.Time  `db:"created_at"`
		UpdatedAt    *time.Time `db:"updated_at"`
		DeletedAt    *time.Time `db:"deleted_at"`
	}
)
