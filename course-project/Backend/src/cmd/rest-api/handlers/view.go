package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/cmd/rest-api/lib"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/db"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/models"
)

func Select(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := db.AcquireConn(p)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	id, err := lib.IDFromVars(r)
	if err != nil { // bad request
		lib.ReturnClientError(w, err.Error())
		return
	}
	rec, err := records.Select(conn, id)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, rec)
}
