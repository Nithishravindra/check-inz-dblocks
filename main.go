package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/nithishravindra/sql-locks/internal/exclusivelock"
	"github.com/nithishravindra/sql-locks/internal/models"
	"github.com/nithishravindra/sql-locks/internal/mysql"
	"github.com/nithishravindra/sql-locks/internal/skiplock"
	"github.com/nithishravindra/sql-locks/internal/utils"
	"github.com/nithishravindra/sql-locks/internal/withoutlock"
)

func withoutlockMethod(users []models.User, pool *mysql.ConnPool) {
	log.Println("Transaction without locks")
	var wg sync.WaitGroup
	wg.Add(len(users))
	log.Printf("Found %d users", len(users))
	utils.ResetSeatDetails(pool)
	for _, user := range users {
		go func(user models.User) {
			defer wg.Done()
			seat, err := withoutlock.BookSeat(user, pool)
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

func withExclusiveLock(users []models.User, pool *mysql.ConnPool) {
	log.Println("Transaction with exclusive locks")
	var wg sync.WaitGroup
	wg.Add(len(users))
	log.Printf("Found %d users", len(users))
	utils.ResetSeatDetails(pool)
	for _, user := range users {
		go func(user models.User) {
			defer wg.Done()
			seat, err := exclusivelock.BookSeat(user, pool)
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

func withSkipLock(users []models.User, pool *mysql.ConnPool) {
	log.Println("Transaction with skip locks")
	var wg sync.WaitGroup
	wg.Add(len(users))
	log.Printf("Found %d users", len(users))
	utils.ResetSeatDetails(pool)
	for _, user := range users {
		go func(user models.User) {
			defer wg.Done()
			seat, err := skiplock.BookSeat(user, pool)
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

func main() {
	runtime.GOMAXPROCS(6)

	pool, err := mysql.GetConnPool(50)
	if err != nil {
		log.Fatalf("Error creating connection pool: %v", err)
	}

	fmt.Printf("pool: %v\n", pool)
	// Fetch all users
	users, err := utils.GetAllUsers(pool)
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}

	withoutlockMethod(users, pool)
	withExclusiveLock(users, pool)
	withSkipLock(users, pool)
}
