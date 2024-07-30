// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: miniblog/v1/miniblog.proto

package miniblog

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Nickname  string                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email     string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string                 `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	PostCount int64                  `protobuf:"varint,6,opt,name=postCount,proto3" json:"postCount,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miniblog_v1_miniblog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_miniblog_v1_miniblog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_miniblog_v1_miniblog_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfo) GetPostCount() int64 {
	if x != nil {
		return x.PostCount
	}
	return 0
}

func (x *UserInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// ListUserRequest 指定了 `ListUser` 接口的请求参数.
type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  *int64 `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset *int64 `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miniblog_v1_miniblog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_miniblog_v1_miniblog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_miniblog_v1_miniblog_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *ListUserRequest) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

// ListUserResponse 指定了 `GET /v1/users` 接口的响应参数.
type ListUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64       `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Users      []*UserInfo `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miniblog_v1_miniblog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_miniblog_v1_miniblog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_miniblog_v1_miniblog_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListUserResponse) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_miniblog_v1_miniblog_proto protoreflect.FileDescriptor

var file_miniblog_v1_miniblog_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x69, 0x6e, 0x69, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69,
	0x6e, 0x69, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x80, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x5e, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x56, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x43, 0x0a, 0x08,
	0x4d, 0x69, 0x6e, 0x69, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x37, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x75, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x79, 0x75, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x62, 0x6c,
	0x6f, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x6e,
	0x69, 0x62, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_miniblog_v1_miniblog_proto_rawDescOnce sync.Once
	file_miniblog_v1_miniblog_proto_rawDescData = file_miniblog_v1_miniblog_proto_rawDesc
)

func file_miniblog_v1_miniblog_proto_rawDescGZIP() []byte {
	file_miniblog_v1_miniblog_proto_rawDescOnce.Do(func() {
		file_miniblog_v1_miniblog_proto_rawDescData = protoimpl.X.CompressGZIP(file_miniblog_v1_miniblog_proto_rawDescData)
	})
	return file_miniblog_v1_miniblog_proto_rawDescData
}

var file_miniblog_v1_miniblog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_miniblog_v1_miniblog_proto_goTypes = []interface{}{
	(*UserInfo)(nil),              // 0: v1.UserInfo
	(*ListUserRequest)(nil),       // 1: v1.ListUserRequest
	(*ListUserResponse)(nil),      // 2: v1.ListUserResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_miniblog_v1_miniblog_proto_depIdxs = []int32{
	3, // 0: v1.UserInfo.createdAt:type_name -> google.protobuf.Timestamp
	3, // 1: v1.UserInfo.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: v1.ListUserResponse.users:type_name -> v1.UserInfo
	1, // 3: v1.MiniBlog.ListUser:input_type -> v1.ListUserRequest
	2, // 4: v1.MiniBlog.ListUser:output_type -> v1.ListUserResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_miniblog_v1_miniblog_proto_init() }
func file_miniblog_v1_miniblog_proto_init() {
	if File_miniblog_v1_miniblog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_miniblog_v1_miniblog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_miniblog_v1_miniblog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_miniblog_v1_miniblog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserResponse); i {
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
	file_miniblog_v1_miniblog_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_miniblog_v1_miniblog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_miniblog_v1_miniblog_proto_goTypes,
		DependencyIndexes: file_miniblog_v1_miniblog_proto_depIdxs,
		MessageInfos:      file_miniblog_v1_miniblog_proto_msgTypes,
	}.Build()
	File_miniblog_v1_miniblog_proto = out.File
	file_miniblog_v1_miniblog_proto_rawDesc = nil
	file_miniblog_v1_miniblog_proto_goTypes = nil
	file_miniblog_v1_miniblog_proto_depIdxs = nil
}
