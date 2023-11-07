package models

// User represents a user in the application.
type User struct {
	Id       int    `json:"id"`                  // Unique identifier for the user.
	Name     string `json:"name"`                // User's name.
	Email    string `json:"email" gorm:"unique"` // User's email (must be unique in the database).
	Password []byte `json:"-"`                   // User's password (stored as a byte slice and omitted from JSON).
}
