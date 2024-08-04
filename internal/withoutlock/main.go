package h

import (
	"database/sql"
	"fmt"

	"github.com/nithishravindra/sql-locks/internal/models"
	"github.com/nithishravindra/sql-locks/internal/mysql"
)

func BookSeat(user models.User, pool *mysql.ConnPool) (*models.Seat, error) {
	conn, err := pool.Get()
	if err != nil {
		return nil, fmt.Errorf("error getting connection from pool: %v", err)
	}
	defer pool.Put(conn)

	txn, err := conn.Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("error beginning transaction: %v", err)
	}
	defer txn.Rollback()

	// Query for available seat
	row := txn.QueryRow(`SELECT id, name, theatre_id, user_id FROM seats WHERE theatre_id = 1 AND user_id IS NULL ORDER BY id LIMIT 1`)

	var seat models.Seat
	err = row.Scan(&seat.ID, &seat.Name, &seat.TheatreID, &seat.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No available seat
		}
		return nil, fmt.Errorf("error querying seat: %v", err)
	}

	// Update seat with user ID
	_, err = txn.Exec("UPDATE seats SET user_id = ? WHERE id = ?", user.ID, seat.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating seat: %v", err)
	}

	err = txn.Commit()
	if err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}
	return &seat, nil
}
