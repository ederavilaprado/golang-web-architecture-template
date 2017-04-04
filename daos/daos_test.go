package daos

import (
	"database/sql"
	"time"

	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/jmoiron/sqlx"
)

func testDBCall(db *sqlx.DB, f func(rs app.RequestScope)) {
	rs := mockRequestScope(db)

	defer func() {
		rs.Tx().Rollback()
	}()

	f(rs)
}

type requestScope struct {
	app.Logger
	tx *sql.Tx
}

func mockRequestScope(db *sqlx.DB) app.RequestScope {
	tx, _ := db.Begin()
	return &requestScope{
		tx: tx,
	}
}

func (rs *requestScope) UserID() string {
	return "tester"
}

func (rs *requestScope) SetUserID(id string) {
}

func (rs *requestScope) RequestID() string {
	return "test"
}

func (rs *requestScope) Tx() *sql.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *sql.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return false
}

func (rs *requestScope) SetRollback(v bool) {
}

func (rs *requestScope) Now() time.Time {
	return time.Now()
}
