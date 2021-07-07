// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: pkg/sysapi/sysapi.proto

package sysapi

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

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{1}
}

type AuthSubjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderId string `protobuf:"bytes,1,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AuthSubjectsRequest) Reset() {
	*x = AuthSubjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthSubjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSubjectsRequest) ProtoMessage() {}

func (x *AuthSubjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthSubjectsRequest.ProtoReflect.Descriptor instead.
func (*AuthSubjectsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{2}
}

func (x *AuthSubjectsRequest) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *AuthSubjectsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AuthSubject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProviderId string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	SubjectId  string `protobuf:"bytes,3,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *AuthSubject) Reset() {
	*x = AuthSubject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthSubject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSubject) ProtoMessage() {}

func (x *AuthSubject) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthSubject.ProtoReflect.Descriptor instead.
func (*AuthSubject) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{3}
}

func (x *AuthSubject) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthSubject) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *AuthSubject) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

var File_pkg_sysapi_sysapi_proto protoreflect.FileDescriptor

var file_pkg_sysapi_sysapi_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x13, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x32, 0xa5, 0x01, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x41, 0x50, 0x49, 0x12, 0x4c, 0x0a,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2f,
	0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x73, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sysapi_sysapi_proto_rawDescOnce sync.Once
	file_pkg_sysapi_sysapi_proto_rawDescData = file_pkg_sysapi_sysapi_proto_rawDesc
)

func file_pkg_sysapi_sysapi_proto_rawDescGZIP() []byte {
	file_pkg_sysapi_sysapi_proto_rawDescOnce.Do(func() {
		file_pkg_sysapi_sysapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sysapi_sysapi_proto_rawDescData)
	})
	return file_pkg_sysapi_sysapi_proto_rawDescData
}

var file_pkg_sysapi_sysapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_sysapi_sysapi_proto_goTypes = []interface{}{
	(*DeleteUserRequest)(nil),   // 0: goalert.v1.DeleteUserRequest
	(*DeleteUserResponse)(nil),  // 1: goalert.v1.DeleteUserResponse
	(*AuthSubjectsRequest)(nil), // 2: goalert.v1.AuthSubjectsRequest
	(*AuthSubject)(nil),         // 3: goalert.v1.AuthSubject
}
var file_pkg_sysapi_sysapi_proto_depIdxs = []int32{
	2, // 0: goalert.v1.SysAPI.AuthSubjects:input_type -> goalert.v1.AuthSubjectsRequest
	0, // 1: goalert.v1.SysAPI.DeleteUser:input_type -> goalert.v1.DeleteUserRequest
	3, // 2: goalert.v1.SysAPI.AuthSubjects:output_type -> goalert.v1.AuthSubject
	1, // 3: goalert.v1.SysAPI.DeleteUser:output_type -> goalert.v1.DeleteUserResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_sysapi_sysapi_proto_init() }
func file_pkg_sysapi_sysapi_proto_init() {
	if File_pkg_sysapi_sysapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sysapi_sysapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthSubjectsRequest); i {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthSubject); i {
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
			RawDescriptor: file_pkg_sysapi_sysapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_sysapi_sysapi_proto_goTypes,
		DependencyIndexes: file_pkg_sysapi_sysapi_proto_depIdxs,
		MessageInfos:      file_pkg_sysapi_sysapi_proto_msgTypes,
	}.Build()
	File_pkg_sysapi_sysapi_proto = out.File
	file_pkg_sysapi_sysapi_proto_rawDesc = nil
	file_pkg_sysapi_sysapi_proto_goTypes = nil
	file_pkg_sysapi_sysapi_proto_depIdxs = nil
}
