// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/app/proto/shortener.proto

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

type Empty struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{0}
}

// requests
type ShortenRequest struct {
	state         protoimpl.MessageState
	Url           string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteUrlsRequest struct {
	state         protoimpl.MessageState
	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	UrlIds        []string `protobuf:"bytes,2,rep,name=url_ids,json=urlIds,proto3" json:"url_ids,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUrlsRequest) Reset() {
	*x = DeleteUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUrlsRequest) ProtoMessage() {}

func (x *DeleteUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUrlsRequest.ProtoReflect.Descriptor instead.
func (*DeleteUrlsRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUrlsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteUrlsRequest) GetUrlIds() []string {
	if x != nil {
		return x.UrlIds
	}
	return nil
}

type ExpandRequest struct {
	state         protoimpl.MessageState
	UrlId         string `protobuf:"bytes,1,opt,name=url_id,json=urlId,proto3" json:"url_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandRequest) Reset() {
	*x = ExpandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandRequest) ProtoMessage() {}

func (x *ExpandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandRequest.ProtoReflect.Descriptor instead.
func (*ExpandRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *ExpandRequest) GetUrlId() string {
	if x != nil {
		return x.UrlId
	}
	return ""
}

type ShortenBatchRequest struct {
	state         protoimpl.MessageState
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	Urls          []*ShortenBatchItemRequest `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenBatchRequest) Reset() {
	*x = ShortenBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchRequest) ProtoMessage() {}

func (x *ShortenBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchRequest.ProtoReflect.Descriptor instead.
func (*ShortenBatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *ShortenBatchRequest) GetUrls() []*ShortenBatchItemRequest {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *ShortenBatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ShortenBatchItemRequest struct {
	state         protoimpl.MessageState
	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenBatchItemRequest) Reset() {
	*x = ShortenBatchItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenBatchItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchItemRequest) ProtoMessage() {}

func (x *ShortenBatchItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchItemRequest.ProtoReflect.Descriptor instead.
func (*ShortenBatchItemRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *ShortenBatchItemRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ShortenBatchItemRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

//responses
type ShorteningResponse struct {
	state         protoimpl.MessageState
	ResultUrl     string `protobuf:"bytes,1,opt,name=result_url,json=resultUrl,proto3" json:"result_url,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UrlId         string `protobuf:"bytes,3,opt,name=url_id,json=urlId,proto3" json:"url_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShorteningResponse) Reset() {
	*x = ShorteningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShorteningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShorteningResponse) ProtoMessage() {}

func (x *ShorteningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShorteningResponse.ProtoReflect.Descriptor instead.
func (*ShorteningResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *ShorteningResponse) GetResultUrl() string {
	if x != nil {
		return x.ResultUrl
	}
	return ""
}

func (x *ShorteningResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShorteningResponse) GetUrlId() string {
	if x != nil {
		return x.UrlId
	}
	return ""
}

type ExpandResponse struct {
	state         protoimpl.MessageState
	FullUrl       string `protobuf:"bytes,1,opt,name=full_url,json=fullUrl,proto3" json:"full_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandResponse) Reset() {
	*x = ExpandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandResponse) ProtoMessage() {}

func (x *ExpandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandResponse.ProtoReflect.Descriptor instead.
func (*ExpandResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *ExpandResponse) GetFullUrl() string {
	if x != nil {
		return x.FullUrl
	}
	return ""
}

type ShortenBatchResponse struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	Urls          []*ShortenBatchItemResponse `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenBatchResponse) Reset() {
	*x = ShortenBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchResponse) ProtoMessage() {}

func (x *ShortenBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchResponse.ProtoReflect.Descriptor instead.
func (*ShortenBatchResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *ShortenBatchResponse) GetUrls() []*ShortenBatchItemResponse {
	if x != nil {
		return x.Urls
	}
	return nil
}

type ShortenBatchItemResponse struct {
	state         protoimpl.MessageState
	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ResultUrl     string `protobuf:"bytes,2,opt,name=result_url,json=resultUrl,proto3" json:"result_url,omitempty"`
	UrlId         string `protobuf:"bytes,3,opt,name=url_id,json=urlId,proto3" json:"url_id,omitempty"`
	UserId        string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenBatchItemResponse) Reset() {
	*x = ShortenBatchItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenBatchItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchItemResponse) ProtoMessage() {}

func (x *ShortenBatchItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchItemResponse.ProtoReflect.Descriptor instead.
func (*ShortenBatchItemResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *ShortenBatchItemResponse) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ShortenBatchItemResponse) GetResultUrl() string {
	if x != nil {
		return x.ResultUrl
	}
	return ""
}

func (x *ShortenBatchItemResponse) GetUrlId() string {
	if x != nil {
		return x.UrlId
	}
	return ""
}

func (x *ShortenBatchItemResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_internal_app_proto_shortener_proto protoreflect.FileDescriptor

var file_internal_app_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3b, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x72, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x72, 0x6c, 0x49, 0x64, 0x73, 0x22, 0x26, 0x0a, 0x0d,
	0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x75, 0x72, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75,
	0x72, 0x6c, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x13, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x75,
	0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x75,
	0x72, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x17,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72,
	0x6c, 0x22, 0x63, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x75, 0x72, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x75, 0x72, 0x6c, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6c, 0x6c,
	0x55, 0x72, 0x6c, 0x22, 0x4f, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x75,
	0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x75, 0x72, 0x6c, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x18, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x75, 0x72, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x72, 0x6c, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x9e, 0x02, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x12, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x61,
	0x6e, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_proto_shortener_proto_rawDescOnce sync.Once
	file_internal_app_proto_shortener_proto_rawDescData = file_internal_app_proto_shortener_proto_rawDesc
)

func file_internal_app_proto_shortener_proto_rawDescGZIP() []byte {
	file_internal_app_proto_shortener_proto_rawDescOnce.Do(func() {
		file_internal_app_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_proto_shortener_proto_rawDescData)
	})
	return file_internal_app_proto_shortener_proto_rawDescData
}

var file_internal_app_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_app_proto_shortener_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: shortener.Empty
	(*ShortenRequest)(nil),           // 1: shortener.ShortenRequest
	(*DeleteUrlsRequest)(nil),        // 2: shortener.DeleteUrlsRequest
	(*ExpandRequest)(nil),            // 3: shortener.ExpandRequest
	(*ShortenBatchRequest)(nil),      // 4: shortener.ShortenBatchRequest
	(*ShortenBatchItemRequest)(nil),  // 5: shortener.ShortenBatchItemRequest
	(*ShorteningResponse)(nil),       // 6: shortener.ShorteningResponse
	(*ExpandResponse)(nil),           // 7: shortener.ExpandResponse
	(*ShortenBatchResponse)(nil),     // 8: shortener.ShortenBatchResponse
	(*ShortenBatchItemResponse)(nil), // 9: shortener.ShortenBatchItemResponse
}
var file_internal_app_proto_shortener_proto_depIdxs = []int32{
	5, // 0: shortener.ShortenBatchRequest.urls:type_name -> shortener.ShortenBatchItemRequest
	9, // 1: shortener.ShortenBatchResponse.urls:type_name -> shortener.ShortenBatchItemResponse
	1, // 2: shortener.Shortener.Shorten:input_type -> shortener.ShortenRequest
	2, // 3: shortener.Shortener.DeleteUrls:input_type -> shortener.DeleteUrlsRequest
	3, // 4: shortener.Shortener.Expand:input_type -> shortener.ExpandRequest
	4, // 5: shortener.Shortener.ShortenBatch:input_type -> shortener.ShortenBatchRequest
	6, // 6: shortener.Shortener.Shorten:output_type -> shortener.ShorteningResponse
	0, // 7: shortener.Shortener.DeleteUrls:output_type -> shortener.Empty
	7, // 8: shortener.Shortener.Expand:output_type -> shortener.ExpandResponse
	8, // 9: shortener.Shortener.ShortenBatch:output_type -> shortener.ShortenBatchResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_app_proto_shortener_proto_init() }
func file_internal_app_proto_shortener_proto_init() {
	if File_internal_app_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUrlsRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenBatchRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenBatchItemRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShorteningResponse); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandResponse); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenBatchResponse); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenBatchItemResponse); i {
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
			RawDescriptor: file_internal_app_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_proto_shortener_proto_goTypes,
		DependencyIndexes: file_internal_app_proto_shortener_proto_depIdxs,
		MessageInfos:      file_internal_app_proto_shortener_proto_msgTypes,
	}.Build()
	File_internal_app_proto_shortener_proto = out.File
	file_internal_app_proto_shortener_proto_rawDesc = nil
	file_internal_app_proto_shortener_proto_goTypes = nil
	file_internal_app_proto_shortener_proto_depIdxs = nil
}
