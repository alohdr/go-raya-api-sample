-- bank.bank definition

-- Drop table

-- DROP TABLE bank.bank;

CREATE TABLE bank.bank (
                           bank_id varchar(100) NOT NULL,
                           bank_code varchar(50) NULL,
                           bank_name varchar(100) NOT NULL,
                           bank_admin_fee int4 NULL DEFAULT 0,
                           bank_icon varchar(50) NULL DEFAULT ''::character varying,
                           created_at timestamp NULL DEFAULT now(),
                           updated_at timestamp NULL,
                           deleted_at timestamp NULL,
                           CONSTRAINT bank_bank_code_key UNIQUE (bank_code),
                           CONSTRAINT bank_pkey PRIMARY KEY (bank_id)
);