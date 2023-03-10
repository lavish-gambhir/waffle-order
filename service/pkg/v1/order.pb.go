// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: v1/order.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Waffles     []*Waffle `protobuf:"bytes,2,rep,name=waffles,proto3" json:"waffles,omitempty"`
	Description string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32   `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetWaffles() []*Waffle {
	if x != nil {
		return x.Waffles
	}
	return nil
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Waffle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Waffle) Reset() {
	*x = Waffle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Waffle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Waffle) ProtoMessage() {}

func (x *Waffle) ProtoReflect() protoreflect.Message {
	mi := &file_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Waffle.ProtoReflect.Descriptor instead.
func (*Waffle) Descriptor() ([]byte, []int) {
	return file_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *Waffle) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Waffle) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_v1_order_proto protoreflect.FileDescriptor

var file_v1_order_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x07, 0x77, 0x61, 0x66, 0x66, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x66, 0x66, 0x6c, 0x65, 0x52, 0x07, 0x77, 0x61, 0x66, 0x66,
	0x6c, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x06, 0x57,
	0x61, 0x66, 0x66, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0xf9, 0x01,
	0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x33, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x41, 0x0a, 0x14, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x69, 0x74, 0x68, 0x57, 0x61, 0x66, 0x66, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x01, 0x12, 0x39,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x09,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x28, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_order_proto_rawDescOnce sync.Once
	file_v1_order_proto_rawDescData = file_v1_order_proto_rawDesc
)

func file_v1_order_proto_rawDescGZIP() []byte {
	file_v1_order_proto_rawDescOnce.Do(func() {
		file_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_order_proto_rawDescData)
	})
	return file_v1_order_proto_rawDescData
}

var file_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_order_proto_goTypes = []interface{}{
	(*Order)(nil),                  // 0: v1.Order
	(*Waffle)(nil),                 // 1: v1.Waffle
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_v1_order_proto_depIdxs = []int32{
	1, // 0: v1.Order.waffles:type_name -> v1.Waffle
	2, // 1: v1.OrderManagement.getOrder:input_type -> google.protobuf.StringValue
	0, // 2: v1.OrderManagement.addOrder:input_type -> v1.Order
	2, // 3: v1.OrderManagement.searchWithWaffleName:input_type -> google.protobuf.StringValue
	0, // 4: v1.OrderManagement.updateOrders:input_type -> v1.Order
	0, // 5: v1.OrderManagement.getOrder:output_type -> v1.Order
	2, // 6: v1.OrderManagement.addOrder:output_type -> google.protobuf.StringValue
	0, // 7: v1.OrderManagement.searchWithWaffleName:output_type -> v1.Order
	2, // 8: v1.OrderManagement.updateOrders:output_type -> google.protobuf.StringValue
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_order_proto_init() }
func file_v1_order_proto_init() {
	if File_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Waffle); i {
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
			RawDescriptor: file_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_order_proto_goTypes,
		DependencyIndexes: file_v1_order_proto_depIdxs,
		MessageInfos:      file_v1_order_proto_msgTypes,
	}.Build()
	File_v1_order_proto = out.File
	file_v1_order_proto_rawDesc = nil
	file_v1_order_proto_goTypes = nil
	file_v1_order_proto_depIdxs = nil
}
