package app

import (
	"fmt"
	"net/http"
	"time"
)

type RequestScope interface {
	// TODO: logger should came here

	// TODO: DB should came here, so we can use transactions
	// DB *sqlx.DB

	RequestID() string
	Now() time.Time
}

type requestScope struct {
	now       time.Time
	requestID string
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

func newRequestScope(now time.Time, request *http.Request) RequestScope {
	// l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	// if requestID != "" {
	// 	l.SetField("RequestID", requestID)
	// }

	// TODO: remove this here
	if requestID == "" {
		requestID = "ederrequestxxxxxx"
	}

	fmt.Printf("=> %+v\n", "newRequestScope")
	fmt.Printf("=> %+v\n", requestID)

	return &requestScope{
		now:       now,
		requestID: requestID,
	}

}
