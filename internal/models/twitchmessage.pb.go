// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: twitchmessage.proto

package models

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

type TwitchMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Channel    string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Message    string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Badges     []string `protobuf:"bytes,4,rep,name=badges,proto3" json:"badges,omitempty"`
	Bits       int32    `protobuf:"varint,5,opt,name=bits,proto3" json:"bits,omitempty"`
	Mod        int32    `protobuf:"varint,6,opt,name=mod,proto3" json:"mod,omitempty"`
	Subscribed int32    `protobuf:"varint,7,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
	Color      string   `protobuf:"bytes,8,opt,name=color,proto3" json:"color,omitempty"`
	RoomID     string   `protobuf:"bytes,9,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *TwitchMessage) Reset() {
	*x = TwitchMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_twitchmessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwitchMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwitchMessage) ProtoMessage() {}

func (x *TwitchMessage) ProtoReflect() protoreflect.Message {
	mi := &file_twitchmessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwitchMessage.ProtoReflect.Descriptor instead.
func (*TwitchMessage) Descriptor() ([]byte, []int) {
	return file_twitchmessage_proto_rawDescGZIP(), []int{0}
}

func (x *TwitchMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TwitchMessage) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *TwitchMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TwitchMessage) GetBadges() []string {
	if x != nil {
		return x.Badges
	}
	return nil
}

func (x *TwitchMessage) GetBits() int32 {
	if x != nil {
		return x.Bits
	}
	return 0
}

func (x *TwitchMessage) GetMod() int32 {
	if x != nil {
		return x.Mod
	}
	return 0
}

func (x *TwitchMessage) GetSubscribed() int32 {
	if x != nil {
		return x.Subscribed
	}
	return 0
}

func (x *TwitchMessage) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TwitchMessage) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

var File_twitchmessage_proto protoreflect.FileDescriptor

var file_twitchmessage_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xeb, 0x01,
	0x0a, 0x0d, 0x54, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x69, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x69, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x6f, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x42, 0x42, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65,
	0x6c, 0x66, 0x69, 0x6f, 0x72, 0x65, 0x74, 0x74, 0x69, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x2d, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_twitchmessage_proto_rawDescOnce sync.Once
	file_twitchmessage_proto_rawDescData = file_twitchmessage_proto_rawDesc
)

func file_twitchmessage_proto_rawDescGZIP() []byte {
	file_twitchmessage_proto_rawDescOnce.Do(func() {
		file_twitchmessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_twitchmessage_proto_rawDescData)
	})
	return file_twitchmessage_proto_rawDescData
}

var file_twitchmessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_twitchmessage_proto_goTypes = []interface{}{
	(*TwitchMessage)(nil), // 0: models.TwitchMessage
}
var file_twitchmessage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_twitchmessage_proto_init() }
func file_twitchmessage_proto_init() {
	if File_twitchmessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_twitchmessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwitchMessage); i {
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
			RawDescriptor: file_twitchmessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_twitchmessage_proto_goTypes,
		DependencyIndexes: file_twitchmessage_proto_depIdxs,
		MessageInfos:      file_twitchmessage_proto_msgTypes,
	}.Build()
	File_twitchmessage_proto = out.File
	file_twitchmessage_proto_rawDesc = nil
	file_twitchmessage_proto_goTypes = nil
	file_twitchmessage_proto_depIdxs = nil
}
