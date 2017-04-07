package daos

import (
	"time"

	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func testDBCall(db *sqlx.DB, f func(rs app.RequestContext)) {
	rs := mockRequestScope(db)

	// defer func() {
	// 	rs.Tx().Rollback()
	// }()

	f(rs)
}

type requestContext struct {
	// app.Logger
	echo.Context
}

func mockRequestScope(db *sqlx.DB) app.RequestContext {
	// tx, _ := db.Begin()
	return &requestContext{}
}

// func (rs *requestContext) UserID() string {
// 	return "tester"
// }
//
// func (rs *requestContext) SetUserID(id string) {
// }

func (rs *requestContext) RequestID() string {
	return "fakeRequestID"
}

// func (rs *requestContext) Tx() *sql.Tx {
// 	return rs.tx
// }
//
// func (rs *requestContext) SetTx(tx *sql.Tx) {
// 	rs.tx = tx
// }
//
// func (rs *requestContext) Rollback() bool {
// 	return false
// }
//
// func (rs *requestContext) SetRollback(v bool) {
// }

func (rs *requestContext) Start() time.Time {
	return time.Now()
}
