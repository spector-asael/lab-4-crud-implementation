// Filename: internal/data/balance.go
package data

type Balance struct {
	UserID int64   `json:"user_id"` // the user to which the balance belongs
	Amount float64 `json:"amount"`  // the amount of the balance
}
