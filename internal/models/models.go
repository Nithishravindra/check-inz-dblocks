package models

type User struct {
	ID   int
	Name string
}

type Seat struct {
	ID        int
	Name      string
	TheatreID int  // Make sure this field is in your database schema if needed
	UserID    *int // Use a pointer to handle NULL values
}
