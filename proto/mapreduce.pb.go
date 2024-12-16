// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.19.6
// source: proto/mapreduce.proto

package proto

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

type ArrayInt32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []int32 `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"` // Dati in input per la fase di map
}

func (x *ArrayInt32) Reset() {
	*x = ArrayInt32{}
	mi := &file_proto_mapreduce_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArrayInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayInt32) ProtoMessage() {}

func (x *ArrayInt32) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mapreduce_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayInt32.ProtoReflect.Descriptor instead.
func (*ArrayInt32) Descriptor() ([]byte, []int) {
	return file_proto_mapreduce_proto_rawDescGZIP(), []int{0}
}

func (x *ArrayInt32) GetData() []int32 {
	if x != nil {
		return x.Data
	}
	return nil
}

type Bool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B bool `protobuf:"varint,1,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Bool) Reset() {
	*x = Bool{}
	mi := &file_proto_mapreduce_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool) ProtoMessage() {}

func (x *Bool) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mapreduce_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bool.ProtoReflect.Descriptor instead.
func (*Bool) Descriptor() ([]byte, []int) {
	return file_proto_mapreduce_proto_rawDescGZIP(), []int{1}
}

func (x *Bool) GetB() bool {
	if x != nil {
		return x.B
	}
	return false
}

type PartialInputReduce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker string  `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
	Data   []int32 `protobuf:"varint,2,rep,packed,name=data,proto3" json:"data,omitempty"`
}

func (x *PartialInputReduce) Reset() {
	*x = PartialInputReduce{}
	mi := &file_proto_mapreduce_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartialInputReduce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialInputReduce) ProtoMessage() {}

func (x *PartialInputReduce) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mapreduce_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialInputReduce.ProtoReflect.Descriptor instead.
func (*PartialInputReduce) Descriptor() ([]byte, []int) {
	return file_proto_mapreduce_proto_rawDescGZIP(), []int{2}
}

func (x *PartialInputReduce) GetWorker() string {
	if x != nil {
		return x.Worker
	}
	return ""
}

func (x *PartialInputReduce) GetData() []int32 {
	if x != nil {
		return x.Data
	}
	return nil
}

type Tuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker string `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
	Min    int32  `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max    int32  `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Tuple) Reset() {
	*x = Tuple{}
	mi := &file_proto_mapreduce_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tuple) ProtoMessage() {}

func (x *Tuple) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mapreduce_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tuple.ProtoReflect.Descriptor instead.
func (*Tuple) Descriptor() ([]byte, []int) {
	return file_proto_mapreduce_proto_rawDescGZIP(), []int{3}
}

func (x *Tuple) GetWorker() string {
	if x != nil {
		return x.Worker
	}
	return ""
}

func (x *Tuple) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Tuple) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type Partition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition []*Tuple `protobuf:"bytes,1,rep,name=partition,proto3" json:"partition,omitempty"`
	Worker    string   `protobuf:"bytes,2,opt,name=Worker,proto3" json:"Worker,omitempty"`
}

func (x *Partition) Reset() {
	*x = Partition{}
	mi := &file_proto_mapreduce_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partition) ProtoMessage() {}

func (x *Partition) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mapreduce_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partition.ProtoReflect.Descriptor instead.
func (*Partition) Descriptor() ([]byte, []int) {
	return file_proto_mapreduce_proto_rawDescGZIP(), []int{4}
}

func (x *Partition) GetPartition() []*Tuple {
	if x != nil {
		return x.Partition
	}
	return nil
}

func (x *Partition) GetWorker() string {
	if x != nil {
		return x.Worker
	}
	return ""
}

type Worker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker string `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
}

func (x *Worker) Reset() {
	*x = Worker{}
	mi := &file_proto_mapreduce_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Worker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worker) ProtoMessage() {}

func (x *Worker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mapreduce_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worker.ProtoReflect.Descriptor instead.
func (*Worker) Descriptor() ([]byte, []int) {
	return file_proto_mapreduce_proto_rawDescGZIP(), []int{5}
}

func (x *Worker) GetWorker() string {
	if x != nil {
		return x.Worker
	}
	return ""
}

var File_proto_mapreduce_proto protoreflect.FileDescriptor

var file_proto_mapreduce_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20,
	0x0a, 0x0a, 0x41, 0x72, 0x72, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x14, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x01, 0x62, 0x22, 0x40, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x05, 0x54, 0x75, 0x70, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x4f, 0x0a,
	0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x22, 0x20,
	0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x32, 0xd3, 0x01, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x2d, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x12,
	0x3f, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x24, 0x0a, 0x06, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x32, 0x0f, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x53, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mapreduce_proto_rawDescOnce sync.Once
	file_proto_mapreduce_proto_rawDescData = file_proto_mapreduce_proto_rawDesc
)

func file_proto_mapreduce_proto_rawDescGZIP() []byte {
	file_proto_mapreduce_proto_rawDescOnce.Do(func() {
		file_proto_mapreduce_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mapreduce_proto_rawDescData)
	})
	return file_proto_mapreduce_proto_rawDescData
}

var file_proto_mapreduce_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_mapreduce_proto_goTypes = []any{
	(*ArrayInt32)(nil),         // 0: proto.ArrayInt32
	(*Bool)(nil),               // 1: proto.Bool
	(*PartialInputReduce)(nil), // 2: proto.PartialInputReduce
	(*Tuple)(nil),              // 3: proto.Tuple
	(*Partition)(nil),          // 4: proto.Partition
	(*Worker)(nil),             // 5: proto.Worker
}
var file_proto_mapreduce_proto_depIdxs = []int32{
	3, // 0: proto.Partition.partition:type_name -> proto.Tuple
	0, // 1: proto.WorkerService.MapExecute:input_type -> proto.ArrayInt32
	4, // 2: proto.WorkerService.StartShuffle:input_type -> proto.Partition
	2, // 3: proto.WorkerService.AddPartialInputReduce:input_type -> proto.PartialInputReduce
	5, // 4: proto.WorkerService.Reduce:input_type -> proto.Worker
	1, // 5: proto.WorkerService.MapExecute:output_type -> proto.Bool
	1, // 6: proto.WorkerService.StartShuffle:output_type -> proto.Bool
	1, // 7: proto.WorkerService.AddPartialInputReduce:output_type -> proto.Bool
	1, // 8: proto.WorkerService.Reduce:output_type -> proto.Bool
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_mapreduce_proto_init() }
func file_proto_mapreduce_proto_init() {
	if File_proto_mapreduce_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mapreduce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_mapreduce_proto_goTypes,
		DependencyIndexes: file_proto_mapreduce_proto_depIdxs,
		MessageInfos:      file_proto_mapreduce_proto_msgTypes,
	}.Build()
	File_proto_mapreduce_proto = out.File
	file_proto_mapreduce_proto_rawDesc = nil
	file_proto_mapreduce_proto_goTypes = nil
	file_proto_mapreduce_proto_depIdxs = nil
}
