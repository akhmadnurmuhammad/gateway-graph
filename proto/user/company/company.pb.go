// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/user/company/company.proto

package company

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Code           string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	RegisterNumber string `protobuf:"bytes,4,opt,name=RegisterNumber,proto3" json:"RegisterNumber,omitempty"`
	Phone          string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email          string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	Address        string `protobuf:"bytes,7,opt,name=Address,proto3" json:"Address,omitempty"`
	PicName        string `protobuf:"bytes,8,opt,name=PicName,proto3" json:"PicName,omitempty"`
	PicNumber      string `protobuf:"bytes,9,opt,name=PicNumber,proto3" json:"PicNumber,omitempty"`
	PicEmail       string `protobuf:"bytes,10,opt,name=PicEmail,proto3" json:"PicEmail,omitempty"`
	Country        string `protobuf:"bytes,11,opt,name=Country,proto3" json:"Country,omitempty"`
	City           string `protobuf:"bytes,12,opt,name=City,proto3" json:"City,omitempty"`
	Status         bool   `protobuf:"varint,13,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedBy      string `protobuf:"bytes,14,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	UpdatedBy      string `protobuf:"bytes,15,opt,name=UpdatedBy,proto3" json:"UpdatedBy,omitempty"`
	CreatedAt      string `protobuf:"bytes,16,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt      string `protobuf:"bytes,17,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_company_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_company_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_user_company_company_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Data) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Data) GetRegisterNumber() string {
	if x != nil {
		return x.RegisterNumber
	}
	return ""
}

func (x *Data) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Data) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Data) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Data) GetPicName() string {
	if x != nil {
		return x.PicName
	}
	return ""
}

func (x *Data) GetPicNumber() string {
	if x != nil {
		return x.PicNumber
	}
	return ""
}

func (x *Data) GetPicEmail() string {
	if x != nil {
		return x.PicEmail
	}
	return ""
}

func (x *Data) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Data) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Data) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Data) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Data) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Data) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Data) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Code           string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	RegisterNumber string `protobuf:"bytes,4,opt,name=RegisterNumber,proto3" json:"RegisterNumber,omitempty"`
	Phone          string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email          string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	Address        string `protobuf:"bytes,7,opt,name=Address,proto3" json:"Address,omitempty"`
	PicName        string `protobuf:"bytes,8,opt,name=PicName,proto3" json:"PicName,omitempty"`
	PicNumber      string `protobuf:"bytes,9,opt,name=PicNumber,proto3" json:"PicNumber,omitempty"`
	PicEmail       string `protobuf:"bytes,10,opt,name=PicEmail,proto3" json:"PicEmail,omitempty"`
	Country        string `protobuf:"bytes,11,opt,name=Country,proto3" json:"Country,omitempty"`
	City           string `protobuf:"bytes,12,opt,name=City,proto3" json:"City,omitempty"`
	Status         string `protobuf:"bytes,13,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_company_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_company_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_company_company_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRequest) GetRegisterNumber() string {
	if x != nil {
		return x.RegisterNumber
	}
	return ""
}

func (x *UpdateRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateRequest) GetPicName() string {
	if x != nil {
		return x.PicName
	}
	return ""
}

func (x *UpdateRequest) GetPicNumber() string {
	if x != nil {
		return x.PicNumber
	}
	return ""
}

func (x *UpdateRequest) GetPicEmail() string {
	if x != nil {
		return x.PicEmail
	}
	return ""
}

func (x *UpdateRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpdateRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=IsSuccess,proto3" json:"IsSuccess,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data      *Data  `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Status    int64  `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *MainResponse) Reset() {
	*x = MainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_company_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainResponse) ProtoMessage() {}

func (x *MainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_company_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainResponse.ProtoReflect.Descriptor instead.
func (*MainResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_company_company_proto_rawDescGZIP(), []int{2}
}

func (x *MainResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *MainResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MainResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MainResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type FindByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *FindByIDRequest) Reset() {
	*x = FindByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_company_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIDRequest) ProtoMessage() {}

func (x *FindByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_company_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIDRequest.ProtoReflect.Descriptor instead.
func (*FindByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_company_company_proto_rawDescGZIP(), []int{3}
}

func (x *FindByIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_company_company_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_company_company_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_company_company_proto_rawDescGZIP(), []int{4}
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool       `protobuf:"varint,1,opt,name=IsSuccess,proto3" json:"IsSuccess,omitempty"`
	Message   string     `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data      *Companies `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Status    int64      `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_company_company_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_company_company_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_company_company_proto_rawDescGZIP(), []int{5}
}

func (x *FetchResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *FetchResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FetchResponse) GetData() *Companies {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FetchResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Companies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Data `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *Companies) Reset() {
	*x = Companies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_company_company_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Companies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Companies) ProtoMessage() {}

func (x *Companies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_company_company_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Companies.ProtoReflect.Descriptor instead.
func (*Companies) Descriptor() ([]byte, []int) {
	return file_proto_user_company_company_proto_rawDescGZIP(), []int{6}
}

func (x *Companies) GetItems() []*Data {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_user_company_company_proto protoreflect.FileDescriptor

var file_proto_user_company_company_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x22, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbe, 0x03, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x69, 0x63, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x69, 0x63,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x69, 0x63, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x69, 0x63, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xcf, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x69, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x69, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x69, 0x63, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x69, 0x63, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x0d,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xd1, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x12, 0x5f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x4d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x1a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x2b, 0x5a, 0x29, 0x78,
	0x74, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_company_company_proto_rawDescOnce sync.Once
	file_proto_user_company_company_proto_rawDescData = file_proto_user_company_company_proto_rawDesc
)

func file_proto_user_company_company_proto_rawDescGZIP() []byte {
	file_proto_user_company_company_proto_rawDescOnce.Do(func() {
		file_proto_user_company_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_company_company_proto_rawDescData)
	})
	return file_proto_user_company_company_proto_rawDescData
}

var file_proto_user_company_company_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_user_company_company_proto_goTypes = []interface{}{
	(*Data)(nil),            // 0: company.Data
	(*UpdateRequest)(nil),   // 1: company.UpdateRequest
	(*MainResponse)(nil),    // 2: company.MainResponse
	(*FindByIDRequest)(nil), // 3: company.FindByIDRequest
	(*DeleteResponse)(nil),  // 4: company.DeleteResponse
	(*FetchResponse)(nil),   // 5: company.FetchResponse
	(*Companies)(nil),       // 6: company.Companies
}
var file_proto_user_company_company_proto_depIdxs = []int32{
	0, // 0: company.MainResponse.Data:type_name -> company.Data
	6, // 1: company.FetchResponse.Data:type_name -> company.Companies
	0, // 2: company.Companies.Items:type_name -> company.Data
	1, // 3: company.Company.Update:input_type -> company.UpdateRequest
	3, // 4: company.Company.Delete:input_type -> company.FindByIDRequest
	2, // 5: company.Company.Update:output_type -> company.MainResponse
	4, // 6: company.Company.Delete:output_type -> company.DeleteResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_user_company_company_proto_init() }
func file_proto_user_company_company_proto_init() {
	if File_proto_user_company_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_company_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_proto_user_company_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_user_company_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainResponse); i {
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
		file_proto_user_company_company_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIDRequest); i {
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
		file_proto_user_company_company_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_proto_user_company_company_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_proto_user_company_company_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Companies); i {
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
			RawDescriptor: file_proto_user_company_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_company_company_proto_goTypes,
		DependencyIndexes: file_proto_user_company_company_proto_depIdxs,
		MessageInfos:      file_proto_user_company_company_proto_msgTypes,
	}.Build()
	File_proto_user_company_company_proto = out.File
	file_proto_user_company_company_proto_rawDesc = nil
	file_proto_user_company_company_proto_goTypes = nil
	file_proto_user_company_company_proto_depIdxs = nil
}