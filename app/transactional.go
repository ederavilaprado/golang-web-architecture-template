package app

// Transactional returns a handler that encloses the nested handlers with a DB transaction.
// If a nested handler returns an error or a panic happens, it will rollback the transaction.
// Otherwise it will commit the transaction after the nested handlers finish execution.
// By calling app.Context.SetRollback(true), you may also explicitly request to rollback the transaction.
// func Transactional(db *sqlx.DB) routing.Handler {
// 	return func(c echo.Context) error {
// 		// TODO: Improve this code to not use transaction for every single request of the api
// 		tx, err := db.Begin()
// 		if err != nil {
// 			return err
// 		}
//
// 		rs := GetRequestScope(c)
// 		rs.SetTx(tx)
//
// 		err = fault.PanicHandler(rs.Errorf)(c)
//
// 		var e error
// 		if err != nil || rs.Rollback() {
// 			// rollback if a handler returns an error or rollback is explicitly requested
// 			e = tx.Rollback()
// 		} else {
// 			e = tx.Commit()
// 		}
//
// 		if e != nil {
// 			if err == nil {
// 				// the error will be logged by an error handler
// 				return e
// 			}
// 			// log the tx error only
// 			rs.Error(e)
// 		}
//
// 		return err
// 	}
// }
