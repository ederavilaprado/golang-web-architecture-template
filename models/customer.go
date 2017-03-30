package models

type Customer struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"full_name"`
	Email *string `json:"email,omitempty" db:"email"`
	// Email sql.NullString `json:"email,omitempty" db:"email"`
}

// type Customer struct {
// 	ID    int            `db:"id"`
// 	Name  string         `db:"full_name"`
// 	Email sql.NullString `db:"email"`
// }
