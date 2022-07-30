-- bank.accounts definition

-- Drop table

-- DROP TABLE bank.accounts;

CREATE TABLE bank.accounts (
                               account_id varchar(100) NOT NULL,
                               bank_id varchar(100) NOT NULL,
                               account_name varchar(100) NOT NULL,
                               account_bank_number varchar(100) NOT NULL,
                               created_at timestamp NULL DEFAULT now(),
                               updated_at timestamp NULL,
                               deleted_at timestamp NULL,
                               CONSTRAINT accounts_pkey PRIMARY KEY (account_id)
);


-- bank.accounts foreign keys

ALTER TABLE bank.accounts ADD CONSTRAINT fk_bank_id FOREIGN KEY (bank_id) REFERENCES bank.bank(bank_id);