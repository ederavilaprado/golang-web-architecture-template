package daos

import (
	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// ArtistDAO persists artist data in database
type ArtistDAO struct {
	db *dbx.DB
}

// NewArtistDAO creates a new ArtistDAO
func NewArtistDAO(db *dbx.DB) *ArtistDAO {
	return &ArtistDAO{db}
}

// Get reads the artist with the specified ID from the database.
func (dao *ArtistDAO) Get(rs app.RequestScope, id int) (*models.Artist, error) {
	var artist models.Artist
	err := dao.db.Select().Model(id, &artist)
	return &artist, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *ArtistDAO) Create(rs app.RequestScope, artist *models.Artist) error {
	artist.Id = 0
	return dao.db.Model(artist).Insert()
}

// Update saves the changes to an artist in the database.
func (dao *ArtistDAO) Update(rs app.RequestScope, id int, artist *models.Artist) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	artist.Id = id
	return dao.db.Model(artist).Exclude("Id").Update()
}

// Delete deletes an artist with the specified ID from the database.
func (dao *ArtistDAO) Delete(rs app.RequestScope, id int) error {
	artist, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return dao.db.Model(artist).Delete()
}

// Count returns the number of the artist records in the database.
func (dao *ArtistDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := dao.db.Select("COUNT(*)").From("artist").Row(&count)
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *ArtistDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
	artists := []models.Artist{}
	err := dao.db.Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&artists)
	return artists, err
}
