// Filename: internal/data/deposit.go
package data

type Deposit struct {
	UserID        int64   `json:"user_id"`
	BankNumber    int64   `json:"bank_number"`
	DepositAmount float64 `json:"deposit_amount"`
}
