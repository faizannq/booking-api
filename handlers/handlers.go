package handlers

import (
	"coach-booking/models"
	"coach-booking/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetAvailability(c *gin.Context) {
	var a models.Availability

	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := service.SetAvailability(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "availability set"})
}

func GetSlots(c *gin.Context) {
	date := c.Query("date")
	start := date + "T09:00:00Z"
	end := date + "T17:00:00Z"

	startTime, _ := time.Parse(time.RFC3339, start)
	endTime, _ := time.Parse(time.RFC3339, end)

	slots := service.GenerateSlots(startTime, endTime)

	c.JSON(http.StatusOK, slots)
}

func BookSlot(c *gin.Context) {
	var b models.Booking

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := service.BookSlot(b)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "slot already booked"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking confirmed"})
}

func GetUserBookings(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	bookings, err := service.GetUserBookings(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, bookings)
}