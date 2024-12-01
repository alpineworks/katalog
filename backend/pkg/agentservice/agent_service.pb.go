// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.1
// source: backend/pkg/agentservice/agent_service.proto

package agentservice

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

// request object
type PublishDeploymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployments []*Deployment `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments,omitempty"`
}

func (x *PublishDeploymentsRequest) Reset() {
	*x = PublishDeploymentsRequest{}
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishDeploymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishDeploymentsRequest) ProtoMessage() {}

func (x *PublishDeploymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishDeploymentsRequest.ProtoReflect.Descriptor instead.
func (*PublishDeploymentsRequest) Descriptor() ([]byte, []int) {
	return file_backend_pkg_agentservice_agent_service_proto_rawDescGZIP(), []int{0}
}

func (x *PublishDeploymentsRequest) GetDeployments() []*Deployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

// simple abstraction of a kubernetes deployment
type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace    string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Cluster      string            `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Replicas     int32             `protobuf:"varint,4,opt,name=replicas,proto3" json:"replicas,omitempty"`
	TrueReplicas int32             `protobuf:"varint,5,opt,name=true_replicas,json=trueReplicas,proto3" json:"true_replicas,omitempty"`
	Labels       map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Containers   []*Container      `protobuf:"bytes,7,rep,name=containers,proto3" json:"containers,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_backend_pkg_agentservice_agent_service_proto_rawDescGZIP(), []int{1}
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployment) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Deployment) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Deployment) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Deployment) GetTrueReplicas() int32 {
	if x != nil {
		return x.TrueReplicas
	}
	return 0
}

func (x *Deployment) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Deployment) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

// simple abstraction of a kubernetes container
type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Tag   string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_backend_pkg_agentservice_agent_service_proto_rawDescGZIP(), []int{2}
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type PublishDeploymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *PublishDeploymentsResponse) Reset() {
	*x = PublishDeploymentsResponse{}
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishDeploymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishDeploymentsResponse) ProtoMessage() {}

func (x *PublishDeploymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_pkg_agentservice_agent_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishDeploymentsResponse.ProtoReflect.Descriptor instead.
func (*PublishDeploymentsResponse) Descriptor() ([]byte, []int) {
	return file_backend_pkg_agentservice_agent_service_proto_rawDescGZIP(), []int{3}
}

func (x *PublishDeploymentsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PublishDeploymentsResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_backend_pkg_agentservice_agent_service_proto protoreflect.FileDescriptor

var file_backend_pkg_agentservice_agent_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x57, 0x0a, 0x19,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xcb, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x72, 0x75, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x72, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x12, 0x3c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x37, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x5b, 0x0a, 0x1a,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x79, 0x0a, 0x0c, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x12, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x27, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x69, 0x6e, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x6b,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_pkg_agentservice_agent_service_proto_rawDescOnce sync.Once
	file_backend_pkg_agentservice_agent_service_proto_rawDescData = file_backend_pkg_agentservice_agent_service_proto_rawDesc
)

func file_backend_pkg_agentservice_agent_service_proto_rawDescGZIP() []byte {
	file_backend_pkg_agentservice_agent_service_proto_rawDescOnce.Do(func() {
		file_backend_pkg_agentservice_agent_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_pkg_agentservice_agent_service_proto_rawDescData)
	})
	return file_backend_pkg_agentservice_agent_service_proto_rawDescData
}

var file_backend_pkg_agentservice_agent_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_backend_pkg_agentservice_agent_service_proto_goTypes = []any{
	(*PublishDeploymentsRequest)(nil),  // 0: agentservice.PublishDeploymentsRequest
	(*Deployment)(nil),                 // 1: agentservice.Deployment
	(*Container)(nil),                  // 2: agentservice.Container
	(*PublishDeploymentsResponse)(nil), // 3: agentservice.PublishDeploymentsResponse
	nil,                                // 4: agentservice.Deployment.LabelsEntry
}
var file_backend_pkg_agentservice_agent_service_proto_depIdxs = []int32{
	1, // 0: agentservice.PublishDeploymentsRequest.deployments:type_name -> agentservice.Deployment
	4, // 1: agentservice.Deployment.labels:type_name -> agentservice.Deployment.LabelsEntry
	2, // 2: agentservice.Deployment.containers:type_name -> agentservice.Container
	0, // 3: agentservice.AgentService.PublishDeployments:input_type -> agentservice.PublishDeploymentsRequest
	3, // 4: agentservice.AgentService.PublishDeployments:output_type -> agentservice.PublishDeploymentsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_backend_pkg_agentservice_agent_service_proto_init() }
func file_backend_pkg_agentservice_agent_service_proto_init() {
	if File_backend_pkg_agentservice_agent_service_proto != nil {
		return
	}
	file_backend_pkg_agentservice_agent_service_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_pkg_agentservice_agent_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_pkg_agentservice_agent_service_proto_goTypes,
		DependencyIndexes: file_backend_pkg_agentservice_agent_service_proto_depIdxs,
		MessageInfos:      file_backend_pkg_agentservice_agent_service_proto_msgTypes,
	}.Build()
	File_backend_pkg_agentservice_agent_service_proto = out.File
	file_backend_pkg_agentservice_agent_service_proto_rawDesc = nil
	file_backend_pkg_agentservice_agent_service_proto_goTypes = nil
	file_backend_pkg_agentservice_agent_service_proto_depIdxs = nil
}