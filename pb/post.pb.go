// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.10.0
// source: pb/post.proto

package postpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_pb_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_pb_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_pb_post_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type CreatePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_pb_post_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type ReadPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *ReadPostRequest) Reset() {
	*x = ReadPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPostRequest) ProtoMessage() {}

func (x *ReadPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPostRequest.ProtoReflect.Descriptor instead.
func (*ReadPostRequest) Descriptor() ([]byte, []int) {
	return file_pb_post_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type ReadPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *ReadPostResponse) Reset() {
	*x = ReadPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPostResponse) ProtoMessage() {}

func (x *ReadPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPostResponse.ProtoReflect.Descriptor instead.
func (*ReadPostResponse) Descriptor() ([]byte, []int) {
	return file_pb_post_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type ListPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPostRequest) Reset() {
	*x = ListPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostRequest) ProtoMessage() {}

func (x *ListPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostRequest.ProtoReflect.Descriptor instead.
func (*ListPostRequest) Descriptor() ([]byte, []int) {
	return file_pb_post_proto_rawDescGZIP(), []int{5}
}

type ListPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *ListPostResponse) Reset() {
	*x = ListPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostResponse) ProtoMessage() {}

func (x *ListPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostResponse.ProtoReflect.Descriptor instead.
func (*ListPostResponse) Descriptor() ([]byte, []int) {
	return file_pb_post_proto_rawDescGZIP(), []int{6}
}

func (x *ListPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

var File_pb_post_proto protoreflect.FileDescriptor

var file_pb_post_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22,
	0x34, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x22, 0x32, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x32, 0xc6, 0x01, 0x0a,
	0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x62, 0x3b, 0x70, 0x6f, 0x73, 0x74,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_post_proto_rawDescOnce sync.Once
	file_pb_post_proto_rawDescData = file_pb_post_proto_rawDesc
)

func file_pb_post_proto_rawDescGZIP() []byte {
	file_pb_post_proto_rawDescOnce.Do(func() {
		file_pb_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_post_proto_rawDescData)
	})
	return file_pb_post_proto_rawDescData
}

var file_pb_post_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_post_proto_goTypes = []interface{}{
	(*Post)(nil),               // 0: post.Post
	(*CreatePostRequest)(nil),  // 1: post.CreatePostRequest
	(*CreatePostResponse)(nil), // 2: post.CreatePostResponse
	(*ReadPostRequest)(nil),    // 3: post.ReadPostRequest
	(*ReadPostResponse)(nil),   // 4: post.ReadPostResponse
	(*ListPostRequest)(nil),    // 5: post.ListPostRequest
	(*ListPostResponse)(nil),   // 6: post.ListPostResponse
}
var file_pb_post_proto_depIdxs = []int32{
	0, // 0: post.CreatePostRequest.post:type_name -> post.Post
	0, // 1: post.CreatePostResponse.post:type_name -> post.Post
	0, // 2: post.ReadPostResponse.post:type_name -> post.Post
	0, // 3: post.ListPostResponse.post:type_name -> post.Post
	1, // 4: post.PostService.CreatePost:input_type -> post.CreatePostRequest
	3, // 5: post.PostService.ReadPost:input_type -> post.ReadPostRequest
	5, // 6: post.PostService.ListPost:input_type -> post.ListPostRequest
	2, // 7: post.PostService.CreatePost:output_type -> post.CreatePostResponse
	4, // 8: post.PostService.ReadPost:output_type -> post.ReadPostResponse
	6, // 9: post.PostService.ListPost:output_type -> post.ListPostResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_post_proto_init() }
func file_pb_post_proto_init() {
	if File_pb_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_pb_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_pb_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostResponse); i {
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
		file_pb_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPostRequest); i {
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
		file_pb_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPostResponse); i {
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
		file_pb_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostRequest); i {
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
		file_pb_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostResponse); i {
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
			RawDescriptor: file_pb_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_post_proto_goTypes,
		DependencyIndexes: file_pb_post_proto_depIdxs,
		MessageInfos:      file_pb_post_proto_msgTypes,
	}.Build()
	File_pb_post_proto = out.File
	file_pb_post_proto_rawDesc = nil
	file_pb_post_proto_goTypes = nil
	file_pb_post_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*ReadPostResponse, error)
	ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (PostService_ListPostClient, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*ReadPostResponse, error) {
	out := new(ReadPostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/ReadPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (PostService_ListPostClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PostService_serviceDesc.Streams[0], "/post.PostService/ListPost", opts...)
	if err != nil {
		return nil, err
	}
	x := &postServiceListPostClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PostService_ListPostClient interface {
	Recv() (*ListPostResponse, error)
	grpc.ClientStream
}

type postServiceListPostClient struct {
	grpc.ClientStream
}

func (x *postServiceListPostClient) Recv() (*ListPostResponse, error) {
	m := new(ListPostResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	ReadPost(context.Context, *ReadPostRequest) (*ReadPostResponse, error)
	ListPost(*ListPostRequest, PostService_ListPostServer) error
}

// UnimplementedPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (*UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (*UnimplementedPostServiceServer) ReadPost(context.Context, *ReadPostRequest) (*ReadPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPost not implemented")
}
func (*UnimplementedPostServiceServer) ListPost(*ListPostRequest, PostService_ListPostServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ReadPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ReadPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/ReadPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ReadPost(ctx, req.(*ReadPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPost_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPostRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServiceServer).ListPost(m, &postServiceListPostServer{stream})
}

type PostService_ListPostServer interface {
	Send(*ListPostResponse) error
	grpc.ServerStream
}

type postServiceListPostServer struct {
	grpc.ServerStream
}

func (x *postServiceListPostServer) Send(m *ListPostResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "ReadPost",
			Handler:    _PostService_ReadPost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPost",
			Handler:       _PostService_ListPost_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/post.proto",
}
