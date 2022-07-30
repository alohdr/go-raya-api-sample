-- bank.transactions definition

-- Drop table

-- DROP TABLE bank.transactions;

CREATE TABLE bank.transactions (
                                   transaction_id varchar(100) NOT NULL,
                                   user_id varchar(100) NOT NULL,
                                   account_id varchar(100) NOT NULL,
                                   bank_id varchar(100) NOT NULL,
                                   transaction_admin_fee int4 NULL,
                                   transaction_amount int4 NOT NULL,
                                   transaction_type varchar(50) NOT NULL,
                                   transaction_desc text NULL,
                                   transaction_status public."statustr" NULL DEFAULT 'pending'::statustr,
                                   is_favorite bool NULL DEFAULT false,
                                   created_at timestamp NULL DEFAULT now(),
                                   updated_at timestamp NULL,
                                   deleted_at timestamp NULL,
                                   CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id)
);


-- bank.transactions foreign keys

ALTER TABLE bank.transactions ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES bank.accounts(account_id);
ALTER TABLE bank.transactions ADD CONSTRAINT fk_bank_id FOREIGN KEY (bank_id) REFERENCES bank.bank(bank_id);
ALTER TABLE bank.transactions ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES bank.users(user_id);