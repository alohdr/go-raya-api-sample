-- bank.users definition

-- Drop table

-- DROP TABLE bank.users;

CREATE TABLE bank.users (
                            user_id varchar(100) NOT NULL,
                            user_name varchar(100) NOT NULL,
                            user_pin_bank varchar(6) NOT NULL,
                            user_email varchar(50) NOT NULL,
                            user_password varchar(50) NOT NULL,
                            user_balance int4 NULL DEFAULT 0,
                            created_at timestamp NULL DEFAULT now(),
                            updated_at timestamp NULL,
                            deleted_at timestamp NULL,
                            CONSTRAINT users_pkey PRIMARY KEY (user_id)
);