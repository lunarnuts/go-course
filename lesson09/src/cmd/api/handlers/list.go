package handlers

import (
	"net/http"

	"github.com/lunarnuts/go-course/tree/lesson09/src/cmd/api/lib"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	"github.com/lunarnuts/go-course/tree/lesson09/src/models"
)

func List(conn *db.DBConn, w http.ResponseWriter, r *http.Request) {
	recs, err := models.ListContacts(*conn)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, recs)
}
