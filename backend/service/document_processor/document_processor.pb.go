// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: document_processor.proto

package document_processor

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

type ProcessDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputPath  string `protobuf:"bytes,1,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	OutputPath string `protobuf:"bytes,2,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	FileType   string `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
}

func (x *ProcessDocumentRequest) Reset() {
	*x = ProcessDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_processor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDocumentRequest) ProtoMessage() {}

func (x *ProcessDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_document_processor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDocumentRequest.ProtoReflect.Descriptor instead.
func (*ProcessDocumentRequest) Descriptor() ([]byte, []int) {
	return file_document_processor_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessDocumentRequest) GetInputPath() string {
	if x != nil {
		return x.InputPath
	}
	return ""
}

func (x *ProcessDocumentRequest) GetOutputPath() string {
	if x != nil {
		return x.OutputPath
	}
	return ""
}

func (x *ProcessDocumentRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

type ProcessDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segments string `protobuf:"bytes,1,opt,name=segments,proto3" json:"segments,omitempty"`
}

func (x *ProcessDocumentResponse) Reset() {
	*x = ProcessDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_processor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDocumentResponse) ProtoMessage() {}

func (x *ProcessDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_document_processor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDocumentResponse.ProtoReflect.Descriptor instead.
func (*ProcessDocumentResponse) Descriptor() ([]byte, []int) {
	return file_document_processor_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessDocumentResponse) GetSegments() string {
	if x != nil {
		return x.Segments
	}
	return ""
}

type ModifyDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputPath    string `protobuf:"bytes,1,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	OutputPath   string `protobuf:"bytes,2,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	Replacements string `protobuf:"bytes,3,opt,name=replacements,proto3" json:"replacements,omitempty"`
	FileType     string `protobuf:"bytes,4,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
}

func (x *ModifyDocumentRequest) Reset() {
	*x = ModifyDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_processor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyDocumentRequest) ProtoMessage() {}

func (x *ModifyDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_document_processor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyDocumentRequest.ProtoReflect.Descriptor instead.
func (*ModifyDocumentRequest) Descriptor() ([]byte, []int) {
	return file_document_processor_proto_rawDescGZIP(), []int{2}
}

func (x *ModifyDocumentRequest) GetInputPath() string {
	if x != nil {
		return x.InputPath
	}
	return ""
}

func (x *ModifyDocumentRequest) GetOutputPath() string {
	if x != nil {
		return x.OutputPath
	}
	return ""
}

func (x *ModifyDocumentRequest) GetReplacements() string {
	if x != nil {
		return x.Replacements
	}
	return ""
}

func (x *ModifyDocumentRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

type ModifyDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ModifyDocumentResponse) Reset() {
	*x = ModifyDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_processor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyDocumentResponse) ProtoMessage() {}

func (x *ModifyDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_document_processor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyDocumentResponse.ProtoReflect.Descriptor instead.
func (*ModifyDocumentResponse) Descriptor() ([]byte, []int) {
	return file_document_processor_proto_rawDescGZIP(), []int{3}
}

func (x *ModifyDocumentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_document_processor_proto protoreflect.FileDescriptor

var file_document_processor_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x22, 0x75,
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x35, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x98, 0x01, 0x0a,
	0x15, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xe8, 0x01, 0x0a, 0x11,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x12, 0x6a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a,
	0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x29, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x72, 0x77, 0x6f, 0x6f, 0x64, 0x6c, 0x69, 0x6e, 0x2f,
	0x6c, 0x65, 0x74, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_document_processor_proto_rawDescOnce sync.Once
	file_document_processor_proto_rawDescData = file_document_processor_proto_rawDesc
)

func file_document_processor_proto_rawDescGZIP() []byte {
	file_document_processor_proto_rawDescOnce.Do(func() {
		file_document_processor_proto_rawDescData = protoimpl.X.CompressGZIP(file_document_processor_proto_rawDescData)
	})
	return file_document_processor_proto_rawDescData
}

var file_document_processor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_document_processor_proto_goTypes = []interface{}{
	(*ProcessDocumentRequest)(nil),  // 0: document_processor.ProcessDocumentRequest
	(*ProcessDocumentResponse)(nil), // 1: document_processor.ProcessDocumentResponse
	(*ModifyDocumentRequest)(nil),   // 2: document_processor.ModifyDocumentRequest
	(*ModifyDocumentResponse)(nil),  // 3: document_processor.ModifyDocumentResponse
}
var file_document_processor_proto_depIdxs = []int32{
	0, // 0: document_processor.DocumentProcessor.ProcessDocument:input_type -> document_processor.ProcessDocumentRequest
	2, // 1: document_processor.DocumentProcessor.ModifyDocument:input_type -> document_processor.ModifyDocumentRequest
	1, // 2: document_processor.DocumentProcessor.ProcessDocument:output_type -> document_processor.ProcessDocumentResponse
	3, // 3: document_processor.DocumentProcessor.ModifyDocument:output_type -> document_processor.ModifyDocumentResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_document_processor_proto_init() }
func file_document_processor_proto_init() {
	if File_document_processor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_document_processor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessDocumentRequest); i {
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
		file_document_processor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessDocumentResponse); i {
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
		file_document_processor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyDocumentRequest); i {
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
		file_document_processor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyDocumentResponse); i {
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
			RawDescriptor: file_document_processor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_document_processor_proto_goTypes,
		DependencyIndexes: file_document_processor_proto_depIdxs,
		MessageInfos:      file_document_processor_proto_msgTypes,
	}.Build()
	File_document_processor_proto = out.File
	file_document_processor_proto_rawDesc = nil
	file_document_processor_proto_goTypes = nil
	file_document_processor_proto_depIdxs = nil
}
