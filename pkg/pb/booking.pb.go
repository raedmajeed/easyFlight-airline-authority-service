// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: booking.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PNR           string   `protobuf:"bytes,1,opt,name=PNR,proto3" json:"PNR,omitempty"`
	SeatArray     []string `protobuf:"bytes,2,rep,name=seat_array,json=seatArray,proto3" json:"seat_array,omitempty"`
	FlightChartId int32    `protobuf:"varint,3,opt,name=flight_chart_id,json=flightChartId,proto3" json:"flight_chart_id,omitempty"`
	Email         string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Economy       bool     `protobuf:"varint,5,opt,name=economy,proto3" json:"economy,omitempty"`
}

func (x *SeatRequest) Reset() {
	*x = SeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatRequest) ProtoMessage() {}

func (x *SeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatRequest.ProtoReflect.Descriptor instead.
func (*SeatRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *SeatRequest) GetPNR() string {
	if x != nil {
		return x.PNR
	}
	return ""
}

func (x *SeatRequest) GetSeatArray() []string {
	if x != nil {
		return x.SeatArray
	}
	return nil
}

func (x *SeatRequest) GetFlightChartId() int32 {
	if x != nil {
		return x.FlightChartId
	}
	return 0
}

func (x *SeatRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SeatRequest) GetEconomy() bool {
	if x != nil {
		return x.Economy
	}
	return false
}

type SeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PNR     string   `protobuf:"bytes,1,opt,name=PNR,proto3" json:"PNR,omitempty"`
	SeatNos []string `protobuf:"bytes,2,rep,name=seat_nos,json=seatNos,proto3" json:"seat_nos,omitempty"`
}

func (x *SeatResponse) Reset() {
	*x = SeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatResponse) ProtoMessage() {}

func (x *SeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatResponse.ProtoReflect.Descriptor instead.
func (*SeatResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *SeatResponse) GetPNR() string {
	if x != nil {
		return x.PNR
	}
	return ""
}

func (x *SeatResponse) GetSeatNos() []string {
	if x != nil {
		return x.SeatNos
	}
	return nil
}

type ConfirmedSeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Economy               bool    `protobuf:"varint,1,opt,name=economy,proto3" json:"economy,omitempty"`
	Travellers            int32   `protobuf:"varint,2,opt,name=travellers,proto3" json:"travellers,omitempty"`
	FlightChartIdDirect   []int32 `protobuf:"varint,3,rep,packed,name=flight_chart_id_direct,json=flightChartIdDirect,proto3" json:"flight_chart_id_direct,omitempty"`
	FlightChartIdIndirect []int32 `protobuf:"varint,4,rep,packed,name=flight_chart_id_indirect,json=flightChartIdIndirect,proto3" json:"flight_chart_id_indirect,omitempty"`
}

func (x *ConfirmedSeatRequest) Reset() {
	*x = ConfirmedSeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmedSeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmedSeatRequest) ProtoMessage() {}

func (x *ConfirmedSeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmedSeatRequest.ProtoReflect.Descriptor instead.
func (*ConfirmedSeatRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmedSeatRequest) GetEconomy() bool {
	if x != nil {
		return x.Economy
	}
	return false
}

func (x *ConfirmedSeatRequest) GetTravellers() int32 {
	if x != nil {
		return x.Travellers
	}
	return 0
}

func (x *ConfirmedSeatRequest) GetFlightChartIdDirect() []int32 {
	if x != nil {
		return x.FlightChartIdDirect
	}
	return nil
}

func (x *ConfirmedSeatRequest) GetFlightChartIdIndirect() []int32 {
	if x != nil {
		return x.FlightChartIdIndirect
	}
	return nil
}

type ConfirmedSeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfirmedSeatResponse) Reset() {
	*x = ConfirmedSeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmedSeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmedSeatResponse) ProtoMessage() {}

func (x *ConfirmedSeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmedSeatResponse.ProtoReflect.Descriptor instead.
func (*ConfirmedSeatResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x4e, 0x52, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x4e, 0x52, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x74,
	0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x61, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x22,
	0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x50, 0x4e, 0x52, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x4e,
	0x52, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x73, 0x22, 0xbe, 0x01, 0x0a,
	0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12,
	0x33, 0x0a, 0x16, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x13, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x49, 0x64, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x15, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x22, 0x17, 0x0a,
	0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9d, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x61, 0x74, 0x12, 0x12, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_booking_proto_goTypes = []interface{}{
	(*SeatRequest)(nil),           // 0: admin.SeatRequest
	(*SeatResponse)(nil),          // 1: admin.SeatResponse
	(*ConfirmedSeatRequest)(nil),  // 2: admin.ConfirmedSeatRequest
	(*ConfirmedSeatResponse)(nil), // 3: admin.ConfirmedSeatResponse
}
var file_booking_proto_depIdxs = []int32{
	0, // 0: admin.AdminService.RegisterSelectSeat:input_type -> admin.SeatRequest
	2, // 1: admin.AdminService.AddConfirmedSeats:input_type -> admin.ConfirmedSeatRequest
	1, // 2: admin.AdminService.RegisterSelectSeat:output_type -> admin.SeatResponse
	3, // 3: admin.AdminService.AddConfirmedSeats:output_type -> admin.ConfirmedSeatResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatRequest); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatResponse); i {
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
		file_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmedSeatRequest); i {
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
		file_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmedSeatResponse); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
