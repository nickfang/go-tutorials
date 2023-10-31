package repository

import (
	"time"

	"github.com/nickfang/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	SearchAvailabilityByDates(roomID int, start, end time.Time) (bool, error)
}
