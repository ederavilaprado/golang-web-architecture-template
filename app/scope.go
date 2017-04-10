package app

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	Logger
	// UserID returns the ID of the user for the current request
	UserID() string
	// SetUserID sets the ID of the currently authenticated user
	SetUserID(id string)
	// RequestID returns the ID of the current request
	RequestID() string
	// Tx returns the currently active database transaction that can be used for DB query purpose
	Tx() *sql.Tx
	// SetTx sets the database transaction
	SetTx(tx *sql.Tx)
	// Rollback returns a value indicating whether the current database transaction should be rolled back
	Rollback() bool
	// SetRollback sets a value indicating whether the current database transaction should be rolled back
	SetRollback(bool)
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
}

type requestScope struct {
	Logger              // the logger tagged with the current request information
	now       time.Time // the time when the request is being processed
	requestID string    // an ID identifying one or multiple correlated HTTP requests
	userID    string    // an ID identifying the current user
	rollback  bool      // whether to roll back the current transaction
	tx        *sql.Tx   // the currently active transaction
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) SetUserID(id string) {
	rs.Logger.SetField("UserID", id)
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Tx() *sql.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *sql.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

// newRequestScope creates a new RequestScope with the current request information.
func newRequestScope(now time.Time, logger *logrus.Logger, request *http.Request) RequestScope {
	l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
	}
}

type RequestContext interface {
	Start() time.Time
	RequestID() string
	echo.Context
}

type requestContext struct {
	crossRequestID *string
	requestID      string
	start          time.Time // the time when the request is being processed
	echo.Context
}

func (uc *requestContext) Start() time.Time {
	return uc.start
}

func (uc *requestContext) RequestID() string {
	return uc.requestID
}

func newRequestContext(now time.Time, logger *logrus.Logger, c echo.Context) RequestContext {
	// l := NewLogger(logger, logrus.Fields{})
	uc := &requestContext{
		requestID: uuid.NewV4().String(),
		start:     now,
		Context:   c,
	}
	// Get requestID from header
	if crossRequestID := c.Request().Header.Get("X-Request-Id"); crossRequestID != "" {
		uc.crossRequestID = &crossRequestID
	}

	// TODO: set te requestID to the logger
	// if requestID != "" {
	// 	l.SetField("RequestID", requestID)
	// }
	// TODO: set crossRequestID to the logger

	return uc
}

func GetRequestContext(c echo.Context) RequestContext {
	return c.(RequestContext)
}

func RequestContextMiddleware(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			now := time.Now()
			uc := newRequestContext(now, logger, c)
			return next(uc)
		}
	}
}
