migrateup:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/bank?sslmode=disable" -verbose up 
migratedown:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/bank?sslmode=disable" -verbose down 

migrate-schema:
	migrate create -ext sql -dir migrations -seq bank_schema