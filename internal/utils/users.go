package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nithishravindra/sql-locks/internal/models"
	"github.com/nithishravindra/sql-locks/internal/mysql"
)

func GetAllUsers(pool *mysql.ConnPool) ([]models.User, error) {
	conn, _ := pool.Get()
	defer pool.Put(conn)

	SQLQUERY := "SELECT id, name FROM users;"

	// Execute the query
	rows, err := conn.Db.Query(SQLQUERY)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to store the users
	var users []models.User

	// Iterate over the results
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return users, nil
}

func PrintSeatDetails(pool *mysql.ConnPool) {
	conn, _ := pool.Get()
	defer pool.Put(conn)

	SQLQUERY := "SELECT COUNT(*) FROM seats WHERE user_id IS NOT NULL"
	var bookedSeats int

	err := conn.Db.QueryRow(SQLQUERY).Scan(&bookedSeats)
	if err != nil {
		panic(err)
	}
	log.Printf("Number of booked seats: %d\n", bookedSeats)

}

func ResetSeatDetails(pool *mysql.ConnPool) {
	conn, _ := pool.Get()
	defer pool.Put(conn)

	_, err := conn.Db.Exec("UPDATE seats SET user_id = NULL WHERE theatre_id = ?", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User IDs set to NULL for theatre_id %d\n", 1)
}

// PrintSeatingArrangement prints the seating arrangement with occupied seats marked
func PrintSeatingArrangement(pool *mysql.ConnPool) {
	conn, err := pool.Get()
	if err != nil {
		log.Fatalf("Error getting connection from pool: %v", err)
	}
	defer pool.Put(conn)

	rows, err := conn.Db.Query("SELECT name, user_id FROM seats WHERE theatre_id = ?", 1)
	if err != nil {
		log.Fatalf("Error querying seats: %v", err)
	}
	defer rows.Close()

	// Initialize a map to track occupied seats
	occupiedSeats := make(map[string]bool)
	for rows.Next() {
		var seatName string
		var userID sql.NullInt64
		err := rows.Scan(&seatName, &userID)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		if userID.Valid {
			occupiedSeats[seatName] = true
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}

	// Define the seating grid dimensions
	rowsInGrid := 10 // Adjust based on your actual grid layout
	colsInGrid := 20 // Adjust based on your actual grid layout

	// Print the seating arrangement
	fmt.Println("Seating Arrangement:")
	for i := 1; i <= rowsInGrid; i++ {
		for j := 0; j < colsInGrid; j++ {
			seatName := fmt.Sprintf("%d%c", i, 'A'+j)
			if occupiedSeats[seatName] {
				fmt.Print("x ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
