// Filename: cmd/api/comments.go
package handler

import (
  // "encoding/json"
  "net/http"
  "time"
  "math/rand"
  // import the data package which contains the definition for Comment
  "github.com/spector-asael/lab4-crud/internal/data"
)
                 
func (a *ApplicationDependencies) createCommentHandler(
    w http.ResponseWriter,
    r *http.Request,
) {

    var incomingData data.Comment

    err := a.readJSON(w, r, &incomingData)
    if err != nil {
        a.badRequestResponse(w, r, err)
        return
    }

    // Set server-controlled fields AFTER decoding
    incomingData.ID = rand.Int63() % 99999
    incomingData.CreatedAt = time.Now().UTC()
    incomingData.Version = 1

    err = a.writeJSON(
        w,
        http.StatusCreated,
        envelope{"comment": incomingData},
        nil,
    )
    if err != nil {
        a.serverErrorResponse(w, r, err)
    }
}
