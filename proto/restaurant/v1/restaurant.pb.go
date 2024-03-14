// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: restaurant/v1/restaurant.proto

package restaurantv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/cfagudelo96/article/proto/api/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRestaurantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	WeeklySchedule *WeeklySchedule `protobuf:"bytes,2,opt,name=weekly_schedule,json=weeklySchedule,proto3" json:"weekly_schedule,omitempty"`
}

func (x *CreateRestaurantRequest) Reset() {
	*x = CreateRestaurantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_v1_restaurant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRestaurantRequest) ProtoMessage() {}

func (x *CreateRestaurantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_v1_restaurant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRestaurantRequest.ProtoReflect.Descriptor instead.
func (*CreateRestaurantRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_v1_restaurant_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRestaurantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRestaurantRequest) GetWeeklySchedule() *WeeklySchedule {
	if x != nil {
		return x.WeeklySchedule
	}
	return nil
}

type CreateRestaurantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurant *Restaurant `protobuf:"bytes,1,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
}

func (x *CreateRestaurantResponse) Reset() {
	*x = CreateRestaurantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_v1_restaurant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRestaurantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRestaurantResponse) ProtoMessage() {}

func (x *CreateRestaurantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_v1_restaurant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRestaurantResponse.ProtoReflect.Descriptor instead.
func (*CreateRestaurantResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_v1_restaurant_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRestaurantResponse) GetRestaurant() *Restaurant {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	WeeklySchedule *WeeklySchedule        `protobuf:"bytes,3,opt,name=weekly_schedule,json=weeklySchedule,proto3" json:"weekly_schedule,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_v1_restaurant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_v1_restaurant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_restaurant_v1_restaurant_proto_rawDescGZIP(), []int{2}
}

func (x *Restaurant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetWeeklySchedule() *WeeklySchedule {
	if x != nil {
		return x.WeeklySchedule
	}
	return nil
}

func (x *Restaurant) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Restaurant) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type WeeklySchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monday    *Schedule `protobuf:"bytes,1,opt,name=monday,proto3" json:"monday,omitempty"`
	Tuesday   *Schedule `protobuf:"bytes,2,opt,name=tuesday,proto3" json:"tuesday,omitempty"`
	Wednesday *Schedule `protobuf:"bytes,3,opt,name=wednesday,proto3" json:"wednesday,omitempty"`
	Thursday  *Schedule `protobuf:"bytes,4,opt,name=thursday,proto3" json:"thursday,omitempty"`
	Friday    *Schedule `protobuf:"bytes,5,opt,name=friday,proto3" json:"friday,omitempty"`
	Saturday  *Schedule `protobuf:"bytes,6,opt,name=saturday,proto3" json:"saturday,omitempty"`
	Sunday    *Schedule `protobuf:"bytes,7,opt,name=sunday,proto3" json:"sunday,omitempty"`
}

func (x *WeeklySchedule) Reset() {
	*x = WeeklySchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_v1_restaurant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeeklySchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklySchedule) ProtoMessage() {}

func (x *WeeklySchedule) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_v1_restaurant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklySchedule.ProtoReflect.Descriptor instead.
func (*WeeklySchedule) Descriptor() ([]byte, []int) {
	return file_restaurant_v1_restaurant_proto_rawDescGZIP(), []int{3}
}

func (x *WeeklySchedule) GetMonday() *Schedule {
	if x != nil {
		return x.Monday
	}
	return nil
}

func (x *WeeklySchedule) GetTuesday() *Schedule {
	if x != nil {
		return x.Tuesday
	}
	return nil
}

func (x *WeeklySchedule) GetWednesday() *Schedule {
	if x != nil {
		return x.Wednesday
	}
	return nil
}

func (x *WeeklySchedule) GetThursday() *Schedule {
	if x != nil {
		return x.Thursday
	}
	return nil
}

func (x *WeeklySchedule) GetFriday() *Schedule {
	if x != nil {
		return x.Friday
	}
	return nil
}

func (x *WeeklySchedule) GetSaturday() *Schedule {
	if x != nil {
		return x.Saturday
	}
	return nil
}

func (x *WeeklySchedule) GetSunday() *Schedule {
	if x != nil {
		return x.Sunday
	}
	return nil
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpensAt  *v1.Time `protobuf:"bytes,2,opt,name=opens_at,json=opensAt,proto3" json:"opens_at,omitempty"`
	ClosesAt *v1.Time `protobuf:"bytes,3,opt,name=closes_at,json=closesAt,proto3" json:"closes_at,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_v1_restaurant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_v1_restaurant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_restaurant_v1_restaurant_proto_rawDescGZIP(), []int{4}
}

func (x *Schedule) GetOpensAt() *v1.Time {
	if x != nil {
		return x.OpensAt
	}
	return nil
}

func (x *Schedule) GetClosesAt() *v1.Time {
	if x != nil {
		return x.ClosesAt
	}
	return nil
}

type NotFoundError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotFoundError) Reset() {
	*x = NotFoundError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_v1_restaurant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotFoundError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotFoundError) ProtoMessage() {}

func (x *NotFoundError) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_v1_restaurant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotFoundError.ProtoReflect.Descriptor instead.
func (*NotFoundError) Descriptor() ([]byte, []int) {
	return file_restaurant_v1_restaurant_proto_rawDescGZIP(), []int{5}
}

var File_restaurant_v1_restaurant_proto protoreflect.FileDescriptor

var file_restaurant_v1_restaurant_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89,
	0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x77, 0x65,
	0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0e, 0x77, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x55, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x22, 0xfd, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x77, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x0e, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xeb, 0x04, 0x0a, 0x0e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d,
	0x6f, 0x6e, 0x64, 0x61, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x74, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x07, 0x74, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79, 0x12, 0x35, 0x0a, 0x09, 0x77, 0x65, 0x64, 0x6e,
	0x65, 0x73, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x77, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x12,
	0x33, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x74, 0x68, 0x75, 0x72,
	0x73, 0x64, 0x61, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x72, 0x69, 0x64, 0x61, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x66,
	0x72, 0x69, 0x64, 0x61, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x08, 0x73, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x75,
	0x6e, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x3a, 0xf1, 0x01, 0xba, 0x48,
	0xed, 0x01, 0x1a, 0xea, 0x01, 0x0a, 0x24, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x61, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x5f,
	0x6f, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x2f, 0x41, 0x74, 0x20,
	0x6c, 0x65, 0x61, 0x73, 0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x64, 0x61, 0x79, 0x20, 0x6d, 0x75,
	0x73, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x69, 0x74, 0x73, 0x20, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x1a, 0x90, 0x01, 0x68,
	0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x29, 0x20,
	0x7c, 0x7c, 0x20, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x74, 0x75, 0x65, 0x73,
	0x64, 0x61, 0x79, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x77, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x68,
	0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79,
	0x29, 0x20, 0x7c, 0x7c, 0x20, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x66, 0x72,
	0x69, 0x64, 0x61, 0x79, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x73, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x68,
	0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x29, 0x22,
	0xbb, 0x02, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x08,
	0x6f, 0x70, 0x65, 0x6e, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x6f, 0x70,
	0x65, 0x6e, 0x73, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x41, 0x74,
	0x3a, 0xda, 0x01, 0xba, 0x48, 0xd6, 0x01, 0x1a, 0xd3, 0x01, 0x0a, 0x18, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x5f, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x54, 0x68, 0x65, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x20, 0x6d, 0x75, 0x73,
	0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x6f, 0x70, 0x65, 0x6e, 0x73,
	0x5f, 0x61, 0x74, 0x1a, 0x87, 0x01, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x73, 0x5f, 0x61, 0x74, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x20, 0x3e, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x5f, 0x61, 0x74, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x20, 0x7c,
	0x7c, 0x20, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x20, 0x3d, 0x3d, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x73, 0x5f, 0x61, 0x74, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x20, 0x26, 0x26, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x2e, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x20, 0x3e, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x73, 0x5f, 0x61, 0x74, 0x2e, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x29, 0x22, 0x0f, 0x0a,
	0x0d, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xf3,
	0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xdd, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x78, 0x92, 0x41, 0x4b, 0x0a,
	0x0b, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x1a, 0x27, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24,
	0x3a, 0x01, 0x2a, 0x62, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22,
	0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x73, 0x42, 0xa7, 0x02, 0x92, 0x41, 0x6a, 0x12, 0x12, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x2d, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x22, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x17, 0x1a, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x66, 0x61, 0x67, 0x75, 0x64, 0x65, 0x6c, 0x6f,
	0x39, 0x36, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x58,
	0x58, 0xaa, 0x02, 0x0d, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0d, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x19, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restaurant_v1_restaurant_proto_rawDescOnce sync.Once
	file_restaurant_v1_restaurant_proto_rawDescData = file_restaurant_v1_restaurant_proto_rawDesc
)

func file_restaurant_v1_restaurant_proto_rawDescGZIP() []byte {
	file_restaurant_v1_restaurant_proto_rawDescOnce.Do(func() {
		file_restaurant_v1_restaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_v1_restaurant_proto_rawDescData)
	})
	return file_restaurant_v1_restaurant_proto_rawDescData
}

var file_restaurant_v1_restaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_restaurant_v1_restaurant_proto_goTypes = []interface{}{
	(*CreateRestaurantRequest)(nil),  // 0: restaurant.v1.CreateRestaurantRequest
	(*CreateRestaurantResponse)(nil), // 1: restaurant.v1.CreateRestaurantResponse
	(*Restaurant)(nil),               // 2: restaurant.v1.Restaurant
	(*WeeklySchedule)(nil),           // 3: restaurant.v1.WeeklySchedule
	(*Schedule)(nil),                 // 4: restaurant.v1.Schedule
	(*NotFoundError)(nil),            // 5: restaurant.v1.NotFoundError
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*v1.Time)(nil),                  // 7: api.v1.Time
}
var file_restaurant_v1_restaurant_proto_depIdxs = []int32{
	3,  // 0: restaurant.v1.CreateRestaurantRequest.weekly_schedule:type_name -> restaurant.v1.WeeklySchedule
	2,  // 1: restaurant.v1.CreateRestaurantResponse.restaurant:type_name -> restaurant.v1.Restaurant
	3,  // 2: restaurant.v1.Restaurant.weekly_schedule:type_name -> restaurant.v1.WeeklySchedule
	6,  // 3: restaurant.v1.Restaurant.created_at:type_name -> google.protobuf.Timestamp
	6,  // 4: restaurant.v1.Restaurant.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 5: restaurant.v1.WeeklySchedule.monday:type_name -> restaurant.v1.Schedule
	4,  // 6: restaurant.v1.WeeklySchedule.tuesday:type_name -> restaurant.v1.Schedule
	4,  // 7: restaurant.v1.WeeklySchedule.wednesday:type_name -> restaurant.v1.Schedule
	4,  // 8: restaurant.v1.WeeklySchedule.thursday:type_name -> restaurant.v1.Schedule
	4,  // 9: restaurant.v1.WeeklySchedule.friday:type_name -> restaurant.v1.Schedule
	4,  // 10: restaurant.v1.WeeklySchedule.saturday:type_name -> restaurant.v1.Schedule
	4,  // 11: restaurant.v1.WeeklySchedule.sunday:type_name -> restaurant.v1.Schedule
	7,  // 12: restaurant.v1.Schedule.opens_at:type_name -> api.v1.Time
	7,  // 13: restaurant.v1.Schedule.closes_at:type_name -> api.v1.Time
	0,  // 14: restaurant.v1.RestaurantService.CreateRestaurant:input_type -> restaurant.v1.CreateRestaurantRequest
	1,  // 15: restaurant.v1.RestaurantService.CreateRestaurant:output_type -> restaurant.v1.CreateRestaurantResponse
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_restaurant_v1_restaurant_proto_init() }
func file_restaurant_v1_restaurant_proto_init() {
	if File_restaurant_v1_restaurant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restaurant_v1_restaurant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRestaurantRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_restaurant_v1_restaurant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRestaurantResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_restaurant_v1_restaurant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_restaurant_v1_restaurant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeeklySchedule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_restaurant_v1_restaurant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_restaurant_v1_restaurant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotFoundError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_restaurant_v1_restaurant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restaurant_v1_restaurant_proto_goTypes,
		DependencyIndexes: file_restaurant_v1_restaurant_proto_depIdxs,
		MessageInfos:      file_restaurant_v1_restaurant_proto_msgTypes,
	}.Build()
	File_restaurant_v1_restaurant_proto = out.File
	file_restaurant_v1_restaurant_proto_rawDesc = nil
	file_restaurant_v1_restaurant_proto_goTypes = nil
	file_restaurant_v1_restaurant_proto_depIdxs = nil
}
