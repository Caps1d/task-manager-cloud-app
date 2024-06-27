// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: notification/pb/notification.proto

package pb

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

// NotifyMessage
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body   string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Urgent bool   `protobuf:"varint,3,opt,name=urgent,proto3" json:"urgent,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_pb_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_pb_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notification_pb_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notification) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Notification) GetUrgent() bool {
	if x != nil {
		return x.Urgent
	}
	return false
}

// NotifyUser
type NotifyUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       int64         `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Notification *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *NotifyUserRequest) Reset() {
	*x = NotifyUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_pb_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyUserRequest) ProtoMessage() {}

func (x *NotifyUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_pb_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyUserRequest.ProtoReflect.Descriptor instead.
func (*NotifyUserRequest) Descriptor() ([]byte, []int) {
	return file_notification_pb_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyUserRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *NotifyUserRequest) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

type NotifyUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *NotifyUserResponse) Reset() {
	*x = NotifyUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_pb_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyUserResponse) ProtoMessage() {}

func (x *NotifyUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_pb_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyUserResponse.ProtoReflect.Descriptor instead.
func (*NotifyUserResponse) Descriptor() ([]byte, []int) {
	return file_notification_pb_notification_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyUserResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *NotifyUserResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// NotifyTeam
type NotifyTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamID int64         `protobuf:"varint,1,opt,name=teamID,proto3" json:"teamID,omitempty"`
	Body   *Notification `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *NotifyTeamRequest) Reset() {
	*x = NotifyTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_pb_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyTeamRequest) ProtoMessage() {}

func (x *NotifyTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_pb_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyTeamRequest.ProtoReflect.Descriptor instead.
func (*NotifyTeamRequest) Descriptor() ([]byte, []int) {
	return file_notification_pb_notification_proto_rawDescGZIP(), []int{3}
}

func (x *NotifyTeamRequest) GetTeamID() int64 {
	if x != nil {
		return x.TeamID
	}
	return 0
}

func (x *NotifyTeamRequest) GetBody() *Notification {
	if x != nil {
		return x.Body
	}
	return nil
}

type NotifyTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *NotifyTeamResponse) Reset() {
	*x = NotifyTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_pb_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyTeamResponse) ProtoMessage() {}

func (x *NotifyTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_pb_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyTeamResponse.ProtoReflect.Descriptor instead.
func (*NotifyTeamResponse) Descriptor() ([]byte, []int) {
	return file_notification_pb_notification_proto_rawDescGZIP(), []int{4}
}

func (x *NotifyTeamResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *NotifyTeamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_notification_pb_notification_proto protoreflect.FileDescriptor

var file_notification_pb_notification_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x72,
	0x67, 0x65, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x3e, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x42, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5b, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x44, 0x12, 0x2e, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x42, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xbb, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51,
	0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12,
	0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_notification_pb_notification_proto_rawDescOnce sync.Once
	file_notification_pb_notification_proto_rawDescData = file_notification_pb_notification_proto_rawDesc
)

func file_notification_pb_notification_proto_rawDescGZIP() []byte {
	file_notification_pb_notification_proto_rawDescOnce.Do(func() {
		file_notification_pb_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_pb_notification_proto_rawDescData)
	})
	return file_notification_pb_notification_proto_rawDescData
}

var file_notification_pb_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notification_pb_notification_proto_goTypes = []any{
	(*Notification)(nil),       // 0: notification.Notification
	(*NotifyUserRequest)(nil),  // 1: notification.NotifyUserRequest
	(*NotifyUserResponse)(nil), // 2: notification.NotifyUserResponse
	(*NotifyTeamRequest)(nil),  // 3: notification.NotifyTeamRequest
	(*NotifyTeamResponse)(nil), // 4: notification.NotifyTeamResponse
}
var file_notification_pb_notification_proto_depIdxs = []int32{
	0, // 0: notification.NotifyUserRequest.notification:type_name -> notification.Notification
	0, // 1: notification.NotifyTeamRequest.body:type_name -> notification.Notification
	1, // 2: notification.NotificationService.NotifyUser:input_type -> notification.NotifyUserRequest
	3, // 3: notification.NotificationService.NotifyTeam:input_type -> notification.NotifyTeamRequest
	2, // 4: notification.NotificationService.NotifyUser:output_type -> notification.NotifyUserResponse
	4, // 5: notification.NotificationService.NotifyTeam:output_type -> notification.NotifyTeamResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_pb_notification_proto_init() }
func file_notification_pb_notification_proto_init() {
	if File_notification_pb_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_pb_notification_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Notification); i {
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
		file_notification_pb_notification_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyUserRequest); i {
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
		file_notification_pb_notification_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyUserResponse); i {
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
		file_notification_pb_notification_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyTeamRequest); i {
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
		file_notification_pb_notification_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyTeamResponse); i {
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
			RawDescriptor: file_notification_pb_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_pb_notification_proto_goTypes,
		DependencyIndexes: file_notification_pb_notification_proto_depIdxs,
		MessageInfos:      file_notification_pb_notification_proto_msgTypes,
	}.Build()
	File_notification_pb_notification_proto = out.File
	file_notification_pb_notification_proto_rawDesc = nil
	file_notification_pb_notification_proto_goTypes = nil
	file_notification_pb_notification_proto_depIdxs = nil
}