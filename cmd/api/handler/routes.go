package handler

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func (a *ApplicationDependencies)Routes() http.Handler  {

   // setup a new router
   router := httprouter.New()
   router.NotFound = http.HandlerFunc(a.notFoundResponse)
   router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedResponse)
   // setup routes
   router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthcheckHandler)
   router.HandlerFunc(http.MethodPost, "/v1/comments", a.createCommentHandler)
   router.HandlerFunc(http.MethodPost, "/v1/balance", a.checkBalanceHandler)
   router.HandlerFunc(http.MethodPost, "/v1/deposit", a.depositHandler)

   return a.recoverPanic(router)      
  
}
