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
		Get(rs app.RequestContext, id int) (*models.Artist, error)
		Query(rs app.RequestContext, offset, limit int) ([]models.Artist, error)
		Count(rs app.RequestContext) (int, error)
		Create(rs app.RequestContext, model *models.Artist) (*models.Artist, error)
		Update(rs app.RequestContext, id int, model *models.Artist) (*models.Artist, error)
		Delete(rs app.RequestContext, id int) (*models.Artist, error)
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
	artist, err := r.service.Get(app.GetRequestContext(c), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, artist)
}

func (r *artistResource) query(c echo.Context) error {
	rs := app.GetRequestContext(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestContext(c), paginatedList.Offset(), paginatedList.Limit())
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
	artist, err := r.service.Create(app.GetRequestContext(c), payload)
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
	rs := app.GetRequestContext(c)
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
	response, err := r.service.Delete(app.GetRequestContext(c), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}
