package apis

import (
	"strconv"

	"github.com/ederavilaprado/golang-web-architecture-template/util"
	"github.com/labstack/echo"
)

const (
	DEFAULT_PAGE_SIZE int = 100
	MAX_PAGE_SIZE     int = 1000
)

func getPaginatedListFromRequest(c echo.Context, count int) *util.PaginatedList {
	page := parseInt(c.QueryParam("page"), 1)
	perPage := parseInt(c.QueryParam("per_page"), DEFAULT_PAGE_SIZE)
	if perPage <= 0 {
		perPage = DEFAULT_PAGE_SIZE
	}
	if perPage > MAX_PAGE_SIZE {
		perPage = MAX_PAGE_SIZE
	}
	return util.NewPaginatedList(page, perPage, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
