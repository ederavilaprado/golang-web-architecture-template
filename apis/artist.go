package apis

import (
	"net/http"
	"strconv"

	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/models"
	"github.com/labstack/echo"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	artistService interface {
		Get(rs app.RequestScope, id int) (*models.Artist, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Artist) (*models.Artist, error)
		Update(rs app.RequestScope, id int, model *models.Artist) (*models.Artist, error)
		Delete(rs app.RequestScope, id int) (*models.Artist, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	artistResource struct {
		service artistService
	}
)

// ServeArtistResource sets up the routing of artist endpoints and the corresponding handlers.
func ServeArtistResource(rg *echo.Group, service artistService) {
	r := &artistResource{service}
	rg.GET("/artists/:id", r.get)
	rg.GET("/artists", r.query)
	rg.POST("/artists", r.create)
	rg.PUT("/artists/:id", r.update)
	rg.DELETE("/artists/:id", r.delete)
}

func (r *artistResource) get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	artist, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, artist)
}

func (r *artistResource) query(c echo.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.JSON(http.StatusOK, paginatedList)
}

func (r *artistResource) create(c echo.Context) error {
	payload := &models.Artist{}
	if err := c.Bind(payload); err != nil {
		return err
	}
	artist, err := r.service.Create(app.GetRequestScope(c), payload)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, artist)
}

func (r *artistResource) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	rs := app.GetRequestScope(c)
	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}
	if err = c.Bind(model); err != nil {
		return err
	}
	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}

func (r *artistResource) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}
