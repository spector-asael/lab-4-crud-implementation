package handler

import (
	"net/http"
)

const appVersion = "1.0.0"

func (a *ApplicationDependencies)healthcheckHandler(w http.ResponseWriter,
                                               r *http.Request) {
   data := envelope {
                     "status": "available",
                     "system_info": map[string]string{
                             "environment": a.Config.Environment,
                             "version": appVersion,
                    },
   }
   err := a.writeJSON(w, http.StatusOK, data, nil)
   if err != nil {
    a.serverErrorResponse(w, r, err)
   }
}




