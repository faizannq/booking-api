package repository

import (
	"coach-booking/config"
	"coach-booking/models"
)

func CreateAvailability(a models.Availability) error {
	query := `
	INSERT INTO availability (coach_id, day_of_week, start_time, end_time)
	VALUES ($1, $2, $3, $4)`
	_, err := config.DB.Exec(query, a.CoachID, a.Day, a.StartTime, a.EndTime)
	return err
}

// 🔥 Concurrency-safe booking
func CreateBooking(b models.Booking) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	query := `
	INSERT INTO bookings (user_id, coach_id, slot)
	VALUES ($1, $2, $3)`

	_, err = tx.Exec(query, b.UserID, b.CoachID, b.Slot)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetBookingsByUser(userID int) ([]models.Booking, error) {
	rows, err := config.DB.Query(`
		SELECT user_id, coach_id, slot
		FROM bookings
		WHERE user_id=$1 AND slot > NOW()
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking

	for rows.Next() {
		var b models.Booking
		rows.Scan(&b.UserID, &b.CoachID, &b.Slot)
		bookings = append(bookings, b)
	}

	return bookings, nil
}