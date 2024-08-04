package main

import (
	"log"
	"sync"

	"github.com/nithishravindra/sql-locks/internal/models"
	"github.com/nithishravindra/sql-locks/internal/mysql"
	"github.com/nithishravindra/sql-locks/internal/utils"
	a "github.com/nithishravindra/sql-locks/internal/withoutlock"
)

func main() {
	// Create a connection pool
	pool, err := mysql.GetConnPool(5)
	if err != nil {
		log.Fatalf("Error creating connection pool: %v", err)
	}

	// Fetch all users
	users, err := utils.GetAllUsers(pool)
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(users))

	log.Printf("Found %d users", len(users))

	utils.ResetSeatDetails(pool)

	// Print users and assign seats
	for _, user := range users {
		go func(user models.User) {
			defer wg.Done()
			seat, err := a.BookSeat(user, pool)
			if err != nil {
				log.Printf("Error booking seat for user %d: %v", user.ID, err)
				return
			}
			if seat != nil {
				log.Printf("User %d was assigned seat %s", user.ID, seat.Name)
			} else {
				log.Printf("No available seats for user %d", user.ID)
			}
		}(user)
	}
	wg.Wait()

	utils.PrintSeatDetails(pool)

	utils.PrintSeatingArrangement(pool)
}
