package internal

// Code outside the parent directory tree cannot import packages from internal.
// If another project (or even another module) tries to import project/internal/domain,
// it will result in a compile-time error.
//
// In the original tutorial this file sat in a sub-directory called models, so you had
// ./internal/models/models.go. Given we only have one models.go file that contains all the models
// I've decided to put it in its own file without the directory so ./internal/models.go.
//
// The implication is that the package name has to be consistant as you can't have two go files in the
// same directory with different package names. Hence the package name of internal.

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id          int
	Txn_date    time.Time
	Who         string
	Description string
	Payee       string
	Amount      decimal.Decimal
	Category    string
}
