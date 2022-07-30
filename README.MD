# EXAM Raya OJT

## Requirement 

Installed Go 
```bash
$ go version
go version go1.18.4 linux/amd64
```

Create a postgresql database
```sql
CREATE DATABASE bank
```

> Import Query in Directory Database Table 

## Setup Project

Clone Project 
```bash
git clone https://gitlab.com/ojt-bank-raya/golang/exam/hendro-susilo-exam
cd hendro-susilo-exam
```

Run Install Depedency
```bash
go mod tidy
```

Run Rest Api Project (Keep Running) (Automated Migrate Database and Seed)
```bash
go run .
```

Run Service Update Status Transaction (Keep Running)
```bash
go run services/main.go
```