package restaurant_test

import (
	"context"
	"os"
	"testing"

	"github.com/cfagudelo96/article/core/restaurant"
	mock_restaurant "github.com/cfagudelo96/article/core/restaurant/mocks"
	restauranthdlr "github.com/cfagudelo96/article/handlers/restaurant"
	apiv1 "github.com/cfagudelo96/article/proto/api/v1"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"
)

type testHandler struct {
	storerMock *mock_restaurant.MockStorer
	handler    *restauranthdlr.Handler
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateRestaurant(t *testing.T) {
	ctx := context.Background()
	th := buildTestHandler(t)
	h := th.handler
	tests := []struct {
		name    string
		req     *restaurantv1.CreateRestaurantRequest
		mocker  func()
		want    *restaurantv1.CreateRestaurantResponse
		wantErr bool
	}{
		{
			name: "rejects a request if there's no name specified",
			req: &restaurantv1.CreateRestaurantRequest{
				Name: "",
				WeeklySchedule: &restaurantv1.WeeklySchedule{
					Monday: &restaurantv1.Schedule{
						OpensAt:  &apiv1.Time{Hour: 9},
						ClosesAt: &apiv1.Time{Hour: 17},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "rejects a request if no day is specified in the weekly schedule",
			req: &restaurantv1.CreateRestaurantRequest{
				Name:           "McTasty",
				WeeklySchedule: &restaurantv1.WeeklySchedule{},
			},
			wantErr: true,
		},
		{
			name: "rejects a request if the schedule specified in the week is invalid",
			req: &restaurantv1.CreateRestaurantRequest{
				Name: "McTasty",
				WeeklySchedule: &restaurantv1.WeeklySchedule{
					Monday: &restaurantv1.Schedule{
						OpensAt:  &apiv1.Time{Hour: 9},
						ClosesAt: &apiv1.Time{Hour: 5},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "works correctly",
			req: &restaurantv1.CreateRestaurantRequest{
				Name: "McTasty",
				WeeklySchedule: &restaurantv1.WeeklySchedule{
					Monday: &restaurantv1.Schedule{
						OpensAt:  &apiv1.Time{Hour: 9},
						ClosesAt: &apiv1.Time{Hour: 17},
					},
				},
			},
			mocker: func() {
				th.storerMock.EXPECT().CreateRestaurant(ctx, gomock.Any()).Return(nil)
			},
			want: &restaurantv1.CreateRestaurantResponse{
				Restaurant: &restaurantv1.Restaurant{
					Name: "McTasty",
					WeeklySchedule: &restaurantv1.WeeklySchedule{
						Monday: &restaurantv1.Schedule{
							OpensAt:  &apiv1.Time{Hour: 9},
							ClosesAt: &apiv1.Time{Hour: 17},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mocker != nil {
				tt.mocker()
			}
			got, err := h.CreateRestaurant(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateRestaurant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got == nil {
				return
			}
			wr := tt.want.GetRestaurant()
			gr := got.GetRestaurant()
			if gr.GetCreatedAt() == nil || gr.GetUpdatedAt() == nil || gr.GetId() == "" {
				t.Error("Handler.CreateRestaurant() invalid restaurant")
				return
			}
			gr.CreatedAt = wr.GetCreatedAt()
			gr.UpdatedAt = wr.GetUpdatedAt()
			gr.Id = wr.GetId()
			if !proto.Equal(wr, gr) {
				t.Errorf("Handler.CreateRestaurant() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var blackhole *restaurantv1.CreateRestaurantResponse

func BenchmarkCreateRestaurant(b *testing.B) {
	ctx := context.Background()
	c := restaurant.NewCore(&MockStorer{})
	h, err := restauranthdlr.NewHandler(c)
	if err != nil {
		b.Fatal(err)
	}
	req := &restaurantv1.CreateRestaurantRequest{
		Name: "McTasty",
		WeeklySchedule: &restaurantv1.WeeklySchedule{
			Monday: &restaurantv1.Schedule{
				OpensAt:  &apiv1.Time{Hour: 9},
				ClosesAt: &apiv1.Time{Hour: 17},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		blackhole, err = h.CreateRestaurant(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCreateRestaurantPointer(b *testing.B) {
	ctx := context.Background()
	c := restaurant.NewCore(&MockStorer{})
	h, err := restauranthdlr.NewHandler(c)
	if err != nil {
		b.Fatal(err)
	}
	req := &restaurantv1.CreateRestaurantRequest{
		Name: "McTasty",
		WeeklySchedule: &restaurantv1.WeeklySchedule{
			Monday: &restaurantv1.Schedule{
				OpensAt:  &apiv1.Time{Hour: 9},
				ClosesAt: &apiv1.Time{Hour: 17},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		blackhole, err = h.CreateRestaurantPointer(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkCreateRestaurant-10    	  276921	      3712 ns/op	    1226 B/op	      22 allocs/op
// BenchmarkCreateRestaurantPointer-10    	  373357	      3185 ns/op	     503 B/op	      15 allocs/op

type MockStorer struct {
}

func (m *MockStorer) CreateRestaurant(_ context.Context, _ restaurant.Restaurant) error {
	return nil
}

func (m *MockStorer) CreateRestaurantPointer(_ context.Context, _ *restaurantv1.Restaurant) error {
	return nil
}

func buildTestHandler(t *testing.T) *testHandler {
	ctrl := gomock.NewController(t)
	storerMock := mock_restaurant.NewMockStorer(ctrl)
	c := restaurant.NewCore(storerMock)
	h, err := restauranthdlr.NewHandler(c)
	if err != nil {
		t.Fatalf("initializing handler: %s", err)
	}
	t.Cleanup(func() {
		ctrl.Finish()
	})
	return &testHandler{
		storerMock: storerMock,
		handler:    h,
	}
}
