package service

import (
	"coach-booking/models"
	"coach-booking/repository"
	"time"
)

func SetAvailability(a models.Availability) error {
	return repository.CreateAvailability(a)
}

// Generate 30-min slots
func GenerateSlots(start, end time.Time) []time.Time {
	var slots []time.Time

	for start.Before(end) {
		slots = append(slots, start)
		start = start.Add(30 * time.Minute)
	}
	return slots
}

func BookSlot(b models.Booking) error {
	return repository.CreateBooking(b)
}

func GetUserBookings(userID int) ([]models.Booking, error) {
	return repository.GetBookingsByUser(userID)
}