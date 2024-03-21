package restaurant

import (
	"time"

	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Day int

const (
	DayMonday = Day(iota + 1)
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)

type Restaurant struct {
	ID             uuid.UUID
	Name           string
	WeeklySchedule WeeklySchedule
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Time struct {
	Hour   int
	Minute int
}

type Schedule struct {
	OpensAt  Time
	ClosesAt Time
}

type WeeklySchedule map[Day]Schedule

type CreateRequest struct {
	Name           string
	WeeklySchedule WeeklySchedule
}

func (r CreateRequest) ToRestaurant() Restaurant {
	now := time.Now()
	return Restaurant{
		ID:             uuid.New(),
		Name:           r.Name,
		WeeklySchedule: r.WeeklySchedule,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func createRestaurantRequestToRestaurant(req *restaurantv1.CreateRestaurantRequest) *restaurantv1.Restaurant {
	now := timestamppb.Now()
	return &restaurantv1.Restaurant{
		Id:             uuid.NewString(),
		Name:           req.GetName(),
		WeeklySchedule: req.GetWeeklySchedule(),
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
