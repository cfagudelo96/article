package restaurant

import (
	"github.com/cfagudelo96/article/core/restaurant"
	apiv1 "github.com/cfagudelo96/article/proto/api/v1"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProtoRestaurant(r restaurant.Restaurant) *restaurantv1.Restaurant {
	return &restaurantv1.Restaurant{
		Id:             r.ID.String(),
		Name:           r.Name,
		WeeklySchedule: toProtoWeeklySchedule(r.WeeklySchedule),
		CreatedAt:      timestamppb.New(r.CreatedAt),
		UpdatedAt:      timestamppb.New(r.UpdatedAt),
	}
}

func toProtoWeeklySchedule(ws restaurant.WeeklySchedule) *restaurantv1.WeeklySchedule {
	var pws restaurantv1.WeeklySchedule
	if s, ok := ws[restaurant.DayMonday]; ok {
		pws.Monday = toProtoSchedule(s)
	}
	if s, ok := ws[restaurant.DayTuesday]; ok {
		pws.Tuesday = toProtoSchedule(s)
	}
	if s, ok := ws[restaurant.DayWednesday]; ok {
		pws.Wednesday = toProtoSchedule(s)
	}
	if s, ok := ws[restaurant.DayThursday]; ok {
		pws.Thursday = toProtoSchedule(s)
	}
	if s, ok := ws[restaurant.DayFriday]; ok {
		pws.Friday = toProtoSchedule(s)
	}
	if s, ok := ws[restaurant.DaySaturday]; ok {
		pws.Saturday = toProtoSchedule(s)
	}
	if s, ok := ws[restaurant.DaySunday]; ok {
		pws.Sunday = toProtoSchedule(s)
	}
	return &pws
}

func toCreateRestaurantRequest(req *restaurantv1.CreateRestaurantRequest) restaurant.CreateRequest {
	return restaurant.CreateRequest{
		Name:           req.GetName(),
		WeeklySchedule: toWeeklySchedule(req.GetWeeklySchedule()),
	}
}

func toWeeklySchedule(ws *restaurantv1.WeeklySchedule) restaurant.WeeklySchedule {
	s := make(restaurant.WeeklySchedule)
	if ws.GetMonday() != nil {
		s[restaurant.DayMonday] = toSchedule(ws.GetMonday())
	}
	if ws.GetTuesday() != nil {
		s[restaurant.DayTuesday] = toSchedule(ws.GetTuesday())
	}
	if ws.GetWednesday() != nil {
		s[restaurant.DayWednesday] = toSchedule(ws.GetWednesday())
	}
	if ws.GetThursday() != nil {
		s[restaurant.DayThursday] = toSchedule(ws.GetThursday())
	}
	if ws.GetFriday() != nil {
		s[restaurant.DayFriday] = toSchedule(ws.GetFriday())
	}
	if ws.GetSaturday() != nil {
		s[restaurant.DaySaturday] = toSchedule(ws.GetSaturday())
	}
	if ws.GetSunday() != nil {
		s[restaurant.DaySunday] = toSchedule(ws.GetSunday())
	}
	return s
}

func toSchedule(s *restaurantv1.Schedule) restaurant.Schedule {
	return restaurant.Schedule{
		OpensAt:  toTime(s.GetOpensAt()),
		ClosesAt: toTime(s.GetClosesAt()),
	}
}

func toProtoSchedule(s restaurant.Schedule) *restaurantv1.Schedule {
	return &restaurantv1.Schedule{
		OpensAt:  toProtoTime(s.OpensAt),
		ClosesAt: toProtoTime(s.ClosesAt),
	}
}

func toTime(t *apiv1.Time) restaurant.Time {
	return restaurant.Time{
		Hour:   int(t.GetHour()),
		Minute: int(t.GetMinute()),
	}
}

func toProtoTime(t restaurant.Time) *apiv1.Time {
	return &apiv1.Time{
		Hour:   int32(t.Hour),
		Minute: int32(t.Minute),
	}
}
