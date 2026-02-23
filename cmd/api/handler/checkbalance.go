// Filename: cmd/api/checkbalance.go
package handler

import (
  // "encoding/json"
  "fmt"
  "net/http"
  // import the data package which contains the definition for Comment
  "github.com/spector-asael/lab4-crud/internal/data"
)
                 
func (a *ApplicationDependencies) checkBalanceHandler(
    w http.ResponseWriter,
    r *http.Request,
) {

    var incomingData struct {
        UserID     int64 `json:"user_id"`
        BankNumber int64 `json:"bank_number"`
    }

    err := a.readJSON(w, r, &incomingData)
    if err != nil {
        a.badRequestResponse(w, r, err)
        return
    }

    amount, ok := validateBankAccount(incomingData.UserID, incomingData.BankNumber)
    if !ok {
        a.badRequestResponse(w, r, fmt.Errorf("invalid user_id or bank_number"))
        return
    }

    // Create your data.Balance struct
    balance := data.Balance{
        UserID: incomingData.UserID,
        Amount: amount,
    }

    err = a.writeJSON(
        w,
        http.StatusOK,
        envelope{"balance": balance},
        nil,
    )
    if err != nil {
        a.serverErrorResponse(w, r, err)
    }
}