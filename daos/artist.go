package daos

import (
	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/models"
	"github.com/jmoiron/sqlx"
)

// ArtistDAO persists artist data in database
type ArtistDAO struct {
	db *sqlx.DB
}

// NewArtistDAO creates a new ArtistDAO
func NewArtistDAO(db *sqlx.DB) *ArtistDAO {
	return &ArtistDAO{db}
}

// Get reads the artist with the specified ID from the database.
func (dao *ArtistDAO) Get(rs app.RequestScope, id int) (*models.Artist, error) {
	// TODO: try to improve this part of the code
	artist := &models.Artist{}
	err := dao.db.Get(artist, "SELECT * FROM artist WHERE id = $1", id)
	return artist, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *ArtistDAO) Create(rs app.RequestScope, artist *models.Artist) error {
	var id int
	err := dao.db.Get(&id, "INSERT INTO artist (name) VALUES ($1) RETURNING id", artist.Name)
	// r, err := dao.db.Exec("INSERT INTO artist (name) VALUES ($1) RETURNING id", artist.Name)
	// rows, _ := r.RowsAffected()
	// fmt.Printf("=> %+v\n", rows)
	// id, _ := r.LastInsertId()
	// fmt.Printf("=> %+v\n", err)
	// fmt.Printf("=> %+v\n", id)
	artist.Id = id
	return err
}

// Update saves the changes to an artist in the database.
func (dao *ArtistDAO) Update(rs app.RequestScope, id int, artist *models.Artist) error {
	// TODO: we can also improve this one... i dont need to check first to update later
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	_, err := dao.db.Exec("UPDATE artist SET name = $1 WHERE id = $2", artist.Name, artist.Id)
	return err
}

// Delete deletes an artist with the specified ID from the database.
func (dao *ArtistDAO) Delete(rs app.RequestScope, id int) error {
	// TODO: improve this one too
	_, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	_, err = dao.db.Exec("DELETE FROM artist WHERE id = $1", id)
	return err
}

// Count returns the number of the artist records in the database.
func (dao *ArtistDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := dao.db.Get(&count, "SELECT COUNT(*) FROM artist")
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *ArtistDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Artist, error) {
	artists := []models.Artist{}
	err := dao.db.Select(&artists, "SELECT * FROM artist ORDER BY id OFFSET $1 LIMIT $2", offset, limit)
	return artists, err
}
