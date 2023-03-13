// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.0
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

type UsersWithoutAuthProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderId string `protobuf:"bytes,1,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
}

func (x *UsersWithoutAuthProviderRequest) Reset() {
	*x = UsersWithoutAuthProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersWithoutAuthProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersWithoutAuthProviderRequest) ProtoMessage() {}

func (x *UsersWithoutAuthProviderRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UsersWithoutAuthProviderRequest.ProtoReflect.Descriptor instead.
func (*UsersWithoutAuthProviderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{0}
}

func (x *UsersWithoutAuthProviderRequest) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

type SetAuthSubjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject *AuthSubject `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *SetAuthSubjectRequest) Reset() {
	*x = SetAuthSubjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAuthSubjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthSubjectRequest) ProtoMessage() {}

func (x *SetAuthSubjectRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetAuthSubjectRequest.ProtoReflect.Descriptor instead.
func (*SetAuthSubjectRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{1}
}

func (x *SetAuthSubjectRequest) GetSubject() *AuthSubject {
	if x != nil {
		return x.Subject
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SetAuthSubjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetAuthSubjectResponse) Reset() {
	*x = SetAuthSubjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAuthSubjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthSubjectResponse) ProtoMessage() {}

func (x *SetAuthSubjectResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SetAuthSubjectResponse.ProtoReflect.Descriptor instead.
func (*SetAuthSubjectResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{3}
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[4]
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
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{4}
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
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[5]
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
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{5}
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
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthSubjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSubjectsRequest) ProtoMessage() {}

func (x *AuthSubjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[6]
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
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{6}
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
		mi := &file_pkg_sysapi_sysapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthSubject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSubject) ProtoMessage() {}

func (x *AuthSubject) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sysapi_sysapi_proto_msgTypes[7]
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
	return file_pkg_sysapi_sysapi_proto_rawDescGZIP(), []int{7}
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
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x42, 0x0a, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x15, 0x53, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x44, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x18, 0x0a, 0x16, 0x53,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
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
	0x49, 0x64, 0x32, 0xe3, 0x02, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x41, 0x50, 0x49, 0x12, 0x4c, 0x0a,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x18, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x59, 0x0a,
	0x0e, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x21, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2f, 0x67, 0x6f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x73, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pkg_sysapi_sysapi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_sysapi_sysapi_proto_goTypes = []interface{}{
	(*UsersWithoutAuthProviderRequest)(nil), // 0: goalert.v1.UsersWithoutAuthProviderRequest
	(*SetAuthSubjectRequest)(nil),           // 1: goalert.v1.SetAuthSubjectRequest
	(*UserInfo)(nil),                        // 2: goalert.v1.UserInfo
	(*SetAuthSubjectResponse)(nil),          // 3: goalert.v1.SetAuthSubjectResponse
	(*DeleteUserRequest)(nil),               // 4: goalert.v1.DeleteUserRequest
	(*DeleteUserResponse)(nil),              // 5: goalert.v1.DeleteUserResponse
	(*AuthSubjectsRequest)(nil),             // 6: goalert.v1.AuthSubjectsRequest
	(*AuthSubject)(nil),                     // 7: goalert.v1.AuthSubject
}
var file_pkg_sysapi_sysapi_proto_depIdxs = []int32{
	7, // 0: goalert.v1.SetAuthSubjectRequest.subject:type_name -> goalert.v1.AuthSubject
	6, // 1: goalert.v1.SysAPI.AuthSubjects:input_type -> goalert.v1.AuthSubjectsRequest
	4, // 2: goalert.v1.SysAPI.DeleteUser:input_type -> goalert.v1.DeleteUserRequest
	0, // 3: goalert.v1.SysAPI.UsersWithoutAuthProvider:input_type -> goalert.v1.UsersWithoutAuthProviderRequest
	1, // 4: goalert.v1.SysAPI.SetAuthSubject:input_type -> goalert.v1.SetAuthSubjectRequest
	7, // 5: goalert.v1.SysAPI.AuthSubjects:output_type -> goalert.v1.AuthSubject
	5, // 6: goalert.v1.SysAPI.DeleteUser:output_type -> goalert.v1.DeleteUserResponse
	2, // 7: goalert.v1.SysAPI.UsersWithoutAuthProvider:output_type -> goalert.v1.UserInfo
	3, // 8: goalert.v1.SysAPI.SetAuthSubject:output_type -> goalert.v1.SetAuthSubjectResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_sysapi_sysapi_proto_init() }
func file_pkg_sysapi_sysapi_proto_init() {
	if File_pkg_sysapi_sysapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sysapi_sysapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersWithoutAuthProviderRequest); i {
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
			switch v := v.(*SetAuthSubjectRequest); i {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAuthSubjectResponse); i {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_sysapi_sysapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   8,
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
