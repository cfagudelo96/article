// Code generated by MockGen. DO NOT EDIT.
// Source: restaurant.go
//
// Generated by this command:
//
//	mockgen -source=restaurant.go -destination=mocks/mocks.go Storer
//

// Package mock_restaurant is a generated GoMock package.
package mock_restaurant

import (
	context "context"
	reflect "reflect"

	restaurant "github.com/cfagudelo96/article/core/restaurant"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockStorer is a mock of Storer interface.
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerMockRecorder
}

// MockStorerMockRecorder is the mock recorder for MockStorer.
type MockStorerMockRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance.
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorer) EXPECT() *MockStorerMockRecorder {
	return m.recorder
}

// CreateRestaurant mocks base method.
func (m *MockStorer) CreateRestaurant(ctx context.Context, r restaurant.Restaurant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRestaurant", ctx, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRestaurant indicates an expected call of CreateRestaurant.
func (mr *MockStorerMockRecorder) CreateRestaurant(ctx, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRestaurant", reflect.TypeOf((*MockStorer)(nil).CreateRestaurant), ctx, r)
}

// CreateRestaurantPointer mocks base method.
func (m *MockStorer) CreateRestaurantPointer(ctx context.Context, r *restaurantv1.Restaurant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRestaurantPointer", ctx, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRestaurantPointer indicates an expected call of CreateRestaurantPointer.
func (mr *MockStorerMockRecorder) CreateRestaurantPointer(ctx, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRestaurantPointer", reflect.TypeOf((*MockStorer)(nil).CreateRestaurantPointer), ctx, r)
}
