// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: api/proto/v1/films.proto

package v1

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

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Genre       string `protobuf:"bytes,4,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_films_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_films_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_films_proto_rawDescGZIP(), []int{0}
}

func (x *Film) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Film) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Film) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Film) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type FilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FilmRequest) Reset() {
	*x = FilmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_films_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmRequest) ProtoMessage() {}

func (x *FilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_films_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmRequest.ProtoReflect.Descriptor instead.
func (*FilmRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_films_proto_rawDescGZIP(), []int{1}
}

func (x *FilmRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FilmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FilmResponse) Reset() {
	*x = FilmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_films_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmResponse) ProtoMessage() {}

func (x *FilmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_films_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmResponse.ProtoReflect.Descriptor instead.
func (*FilmResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_films_proto_rawDescGZIP(), []int{2}
}

func (x *FilmResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_films_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_films_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_films_proto_rawDescGZIP(), []int{3}
}

func (x *SuccessResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_api_proto_v1_films_proto protoreflect.FileDescriptor

var file_api_proto_v1_films_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x69, 0x6c, 0x6d,
	0x73, 0x22, 0x64, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xcb, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x12,
	0x0b, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x1a, 0x13, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x12, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x26, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x0b, 0x2e, 0x66, 0x69,
	0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x1a, 0x0b, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x6d, 0x12, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c,
	0x61, 0x7a, 0x65, 0x65, 0x35, 0x2f, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_films_proto_rawDescOnce sync.Once
	file_api_proto_v1_films_proto_rawDescData = file_api_proto_v1_films_proto_rawDesc
)

func file_api_proto_v1_films_proto_rawDescGZIP() []byte {
	file_api_proto_v1_films_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_films_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_films_proto_rawDescData)
	})
	return file_api_proto_v1_films_proto_rawDescData
}

var file_api_proto_v1_films_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_v1_films_proto_goTypes = []interface{}{
	(*Film)(nil),            // 0: films.Film
	(*FilmRequest)(nil),     // 1: films.FilmRequest
	(*FilmResponse)(nil),    // 2: films.FilmResponse
	(*SuccessResponse)(nil), // 3: films.SuccessResponse
}
var file_api_proto_v1_films_proto_depIdxs = []int32{
	0, // 0: films.FilmService.CreateFilm:input_type -> films.Film
	1, // 1: films.FilmService.GetFilm:input_type -> films.FilmRequest
	0, // 2: films.FilmService.UpdateFilm:input_type -> films.Film
	1, // 3: films.FilmService.DeleteFilm:input_type -> films.FilmRequest
	2, // 4: films.FilmService.CreateFilm:output_type -> films.FilmResponse
	0, // 5: films.FilmService.GetFilm:output_type -> films.Film
	0, // 6: films.FilmService.UpdateFilm:output_type -> films.Film
	3, // 7: films.FilmService.DeleteFilm:output_type -> films.SuccessResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_v1_films_proto_init() }
func file_api_proto_v1_films_proto_init() {
	if File_api_proto_v1_films_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_films_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Film); i {
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
		file_api_proto_v1_films_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmRequest); i {
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
		file_api_proto_v1_films_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmResponse); i {
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
		file_api_proto_v1_films_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
			RawDescriptor: file_api_proto_v1_films_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_films_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_films_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_films_proto_msgTypes,
	}.Build()
	File_api_proto_v1_films_proto = out.File
	file_api_proto_v1_films_proto_rawDesc = nil
	file_api_proto_v1_films_proto_goTypes = nil
	file_api_proto_v1_films_proto_depIdxs = nil
}
