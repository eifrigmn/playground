// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: tencent_rta_stat.proto

package tencent

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

type DB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables    map[string]*Table `protobuf:"bytes,1,rep,name=Tables,proto3" json:"Tables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //表格，Key为Table名，Value为具体的Table内容
	Timestamp int64             `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`                                                                                  //起始时间戳，Unix Timestamp的整型秒，两次发起的请求时间戳范围可能会有重叠，请使用新的覆盖老的
	Interval  int64             `protobuf:"varint,3,opt,name=Interval,proto3" json:"Interval,omitempty"`                                                                                    //每个Value的间隔，单位秒。例如分钟级的统计，该值为60。对于每个Value，其隐含代表的时间戳为 Timestamp + (n * Interval)。n=0 to len()-1；
}

func (x *DB) Reset() {
	*x = DB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tencent_rta_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DB) ProtoMessage() {}

func (x *DB) ProtoReflect() protoreflect.Message {
	mi := &file_tencent_rta_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DB.ProtoReflect.Descriptor instead.
func (*DB) Descriptor() ([]byte, []int) {
	return file_tencent_rta_stat_proto_rawDescGZIP(), []int{0}
}

func (x *DB) GetTables() map[string]*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

func (x *DB) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *DB) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Values `protobuf:"bytes,1,rep,name=Fields,proto3" json:"Fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //Key为Field名称，Value为Field的值集合
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tencent_rta_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_tencent_rta_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_tencent_rta_stat_proto_rawDescGZIP(), []int{1}
}

func (x *Table) GetFields() map[string]*Values {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []int64 `protobuf:"varint,1,rep,packed,name=Value,proto3" json:"Value,omitempty"` //测量点的数值，每Interval一个值。如实际该时间点无值，仍然会填入0
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tencent_rta_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_tencent_rta_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_tencent_rta_stat_proto_rawDescGZIP(), []int{2}
}

func (x *Values) GetValue() []int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type DBRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DBRsp) Reset() {
	*x = DBRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tencent_rta_stat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBRsp) ProtoMessage() {}

func (x *DBRsp) ProtoReflect() protoreflect.Message {
	mi := &file_tencent_rta_stat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBRsp.ProtoReflect.Descriptor instead.
func (*DBRsp) Descriptor() ([]byte, []int) {
	return file_tencent_rta_stat_proto_rawDescGZIP(), []int{3}
}

func (x *DBRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_tencent_rta_stat_proto protoreflect.FileDescriptor

var file_tencent_rta_stat_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x74, 0x61, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x22, 0xba, 0x01, 0x0a, 0x02, 0x44, 0x42, 0x12, 0x2f, 0x0a, 0x06, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x42, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x1a, 0x49, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x87,
	0x01, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x4a, 0x0a, 0x0b,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1e, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1b, 0x0a, 0x05, 0x44, 0x42, 0x52, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x2f, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tencent_rta_stat_proto_rawDescOnce sync.Once
	file_tencent_rta_stat_proto_rawDescData = file_tencent_rta_stat_proto_rawDesc
)

func file_tencent_rta_stat_proto_rawDescGZIP() []byte {
	file_tencent_rta_stat_proto_rawDescOnce.Do(func() {
		file_tencent_rta_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_tencent_rta_stat_proto_rawDescData)
	})
	return file_tencent_rta_stat_proto_rawDescData
}

var file_tencent_rta_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tencent_rta_stat_proto_goTypes = []interface{}{
	(*DB)(nil),     // 0: tencent.DB
	(*Table)(nil),  // 1: tencent.Table
	(*Values)(nil), // 2: tencent.Values
	(*DBRsp)(nil),  // 3: tencent.DBRsp
	nil,            // 4: tencent.DB.TablesEntry
	nil,            // 5: tencent.Table.FieldsEntry
}
var file_tencent_rta_stat_proto_depIdxs = []int32{
	4, // 0: tencent.DB.Tables:type_name -> tencent.DB.TablesEntry
	5, // 1: tencent.Table.Fields:type_name -> tencent.Table.FieldsEntry
	1, // 2: tencent.DB.TablesEntry.value:type_name -> tencent.Table
	2, // 3: tencent.Table.FieldsEntry.value:type_name -> tencent.Values
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tencent_rta_stat_proto_init() }
func file_tencent_rta_stat_proto_init() {
	if File_tencent_rta_stat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tencent_rta_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DB); i {
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
		file_tencent_rta_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_tencent_rta_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Values); i {
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
		file_tencent_rta_stat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBRsp); i {
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
			RawDescriptor: file_tencent_rta_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tencent_rta_stat_proto_goTypes,
		DependencyIndexes: file_tencent_rta_stat_proto_depIdxs,
		MessageInfos:      file_tencent_rta_stat_proto_msgTypes,
	}.Build()
	File_tencent_rta_stat_proto = out.File
	file_tencent_rta_stat_proto_rawDesc = nil
	file_tencent_rta_stat_proto_goTypes = nil
	file_tencent_rta_stat_proto_depIdxs = nil
}