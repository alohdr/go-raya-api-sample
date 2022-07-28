create schema bank;
create table bank.bank (
	bank_code VARCHAR(100) PRIMARY KEY,
	bank_name VARCHAR(100) NOT null,
	bank_admin_fee int not null,
	bank_icon VARCHAR(50),
	created_at TIMESTAMP default now(),
	updated_at TIMESTAMP default now(),
	deleted_at TIMESTAMP 
)