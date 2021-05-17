// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: crud/crud.proto

package crud

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_crud_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_crud_crud_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_crud_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_crud_crud_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_crud_crud_proto protoreflect.FileDescriptor

var file_crud_crud_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x72, 0x75, 0x64, 0x22, 0x62, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x32, 0xad, 0x01, 0x0a, 0x0b, 0x43,
	0x72, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x0a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x26, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0a, 0x2e, 0x63, 0x72,
	0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x26, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0d,
	0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x29, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e,
	0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x63, 0x72, 0x75, 0x64,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x63,
	0x72, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crud_crud_proto_rawDescOnce sync.Once
	file_crud_crud_proto_rawDescData = file_crud_crud_proto_rawDesc
)

func file_crud_crud_proto_rawDescGZIP() []byte {
	file_crud_crud_proto_rawDescOnce.Do(func() {
		file_crud_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_crud_crud_proto_rawDescData)
	})
	return file_crud_crud_proto_rawDescData
}

var file_crud_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crud_crud_proto_goTypes = []interface{}{
	(*User)(nil),    // 0: crud.User
	(*Message)(nil), // 1: crud.Message
}
var file_crud_crud_proto_depIdxs = []int32{
	0, // 0: crud.CrudService.GetUser:input_type -> crud.User
	0, // 1: crud.CrudService.GetUsers:input_type -> crud.User
	0, // 2: crud.CrudService.AddUser:input_type -> crud.User
	0, // 3: crud.CrudService.DeleteUser:input_type -> crud.User
	0, // 4: crud.CrudService.GetUser:output_type -> crud.User
	0, // 5: crud.CrudService.GetUsers:output_type -> crud.User
	1, // 6: crud.CrudService.AddUser:output_type -> crud.Message
	1, // 7: crud.CrudService.DeleteUser:output_type -> crud.Message
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crud_crud_proto_init() }
func file_crud_crud_proto_init() {
	if File_crud_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crud_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_crud_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_crud_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crud_crud_proto_goTypes,
		DependencyIndexes: file_crud_crud_proto_depIdxs,
		MessageInfos:      file_crud_crud_proto_msgTypes,
	}.Build()
	File_crud_crud_proto = out.File
	file_crud_crud_proto_rawDesc = nil
	file_crud_crud_proto_goTypes = nil
	file_crud_crud_proto_depIdxs = nil
}