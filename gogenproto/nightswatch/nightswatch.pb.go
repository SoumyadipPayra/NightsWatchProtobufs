// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: protos/nightswatch/nightswatch.proto

package nightswatch

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceDataType int32

const (
	DeviceDataType_UNKNOWN         DeviceDataType = 0
	DeviceDataType_INSTALLED_APPS  DeviceDataType = 1
	DeviceDataType_OS_VERSION      DeviceDataType = 2
	DeviceDataType_OSQUERY_VERSION DeviceDataType = 3
)

// Enum value maps for DeviceDataType.
var (
	DeviceDataType_name = map[int32]string{
		0: "UNKNOWN",
		1: "INSTALLED_APPS",
		2: "OS_VERSION",
		3: "OSQUERY_VERSION",
	}
	DeviceDataType_value = map[string]int32{
		"UNKNOWN":         0,
		"INSTALLED_APPS":  1,
		"OS_VERSION":      2,
		"OSQUERY_VERSION": 3,
	}
)

func (x DeviceDataType) Enum() *DeviceDataType {
	p := new(DeviceDataType)
	*p = x
	return p
}

func (x DeviceDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_nightswatch_nightswatch_proto_enumTypes[0].Descriptor()
}

func (DeviceDataType) Type() protoreflect.EnumType {
	return &file_protos_nightswatch_nightswatch_proto_enumTypes[0]
}

func (x DeviceDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceDataType.Descriptor instead.
func (DeviceDataType) EnumDescriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{0}
}

// Message definitions will be added here
type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// LoginRequest is the request for the Login method
type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type App struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *App) Reset() {
	*x = App{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{2}
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type InstalledApps struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Apps          []*App                 `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InstalledApps) Reset() {
	*x = InstalledApps{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstalledApps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstalledApps) ProtoMessage() {}

func (x *InstalledApps) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstalledApps.ProtoReflect.Descriptor instead.
func (*InstalledApps) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{3}
}

func (x *InstalledApps) GetApps() []*App {
	if x != nil {
		return x.Apps
	}
	return nil
}

type OSVersion struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OSVersion) Reset() {
	*x = OSVersion{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OSVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSVersion) ProtoMessage() {}

func (x *OSVersion) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSVersion.ProtoReflect.Descriptor instead.
func (*OSVersion) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{4}
}

func (x *OSVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type OSQueryVersion struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OSQueryVersion) Reset() {
	*x = OSQueryVersion{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OSQueryVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSQueryVersion) ProtoMessage() {}

func (x *OSQueryVersion) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSQueryVersion.ProtoReflect.Descriptor instead.
func (*OSQueryVersion) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{5}
}

func (x *OSQueryVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// DeviceDataRequest is the request for the SendDeviceData method
type DeviceDataRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	InstalledApps  *InstalledApps         `protobuf:"bytes,1,opt,name=installed_apps,json=installedApps,proto3" json:"installed_apps,omitempty"`
	OsVersion      *OSVersion             `protobuf:"bytes,2,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	OsqueryVersion *OSQueryVersion        `protobuf:"bytes,3,opt,name=osquery_version,json=osqueryVersion,proto3" json:"osquery_version,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeviceDataRequest) Reset() {
	*x = DeviceDataRequest{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDataRequest) ProtoMessage() {}

func (x *DeviceDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDataRequest.ProtoReflect.Descriptor instead.
func (*DeviceDataRequest) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{6}
}

func (x *DeviceDataRequest) GetInstalledApps() *InstalledApps {
	if x != nil {
		return x.InstalledApps
	}
	return nil
}

func (x *DeviceDataRequest) GetOsVersion() *OSVersion {
	if x != nil {
		return x.OsVersion
	}
	return nil
}

func (x *DeviceDataRequest) GetOsqueryVersion() *OSQueryVersion {
	if x != nil {
		return x.OsqueryVersion
	}
	return nil
}

// LatestDataRequest is the request for the LatestData method
type GetLatestDataRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	UserName         string                 `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	DataRequestTypes []DeviceDataType       `protobuf:"varint,2,rep,packed,name=data_request_types,json=dataRequestTypes,proto3,enum=nightswatch.DeviceDataType" json:"data_request_types,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetLatestDataRequest) Reset() {
	*x = GetLatestDataRequest{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLatestDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestDataRequest) ProtoMessage() {}

func (x *GetLatestDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestDataRequest.ProtoReflect.Descriptor instead.
func (*GetLatestDataRequest) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{7}
}

func (x *GetLatestDataRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetLatestDataRequest) GetDataRequestTypes() []DeviceDataType {
	if x != nil {
		return x.DataRequestTypes
	}
	return nil
}

type GetLatestDataResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	InstalledApps  *InstalledApps         `protobuf:"bytes,1,opt,name=installed_apps,json=installedApps,proto3" json:"installed_apps,omitempty"`
	OsVersion      *OSVersion             `protobuf:"bytes,2,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	OsqueryVersion *OSQueryVersion        `protobuf:"bytes,3,opt,name=osquery_version,json=osqueryVersion,proto3" json:"osquery_version,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetLatestDataResponse) Reset() {
	*x = GetLatestDataResponse{}
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLatestDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestDataResponse) ProtoMessage() {}

func (x *GetLatestDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_nightswatch_nightswatch_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestDataResponse.ProtoReflect.Descriptor instead.
func (*GetLatestDataResponse) Descriptor() ([]byte, []int) {
	return file_protos_nightswatch_nightswatch_proto_rawDescGZIP(), []int{8}
}

func (x *GetLatestDataResponse) GetInstalledApps() *InstalledApps {
	if x != nil {
		return x.InstalledApps
	}
	return nil
}

func (x *GetLatestDataResponse) GetOsVersion() *OSVersion {
	if x != nil {
		return x.OsVersion
	}
	return nil
}

func (x *GetLatestDataResponse) GetOsqueryVersion() *OSQueryVersion {
	if x != nil {
		return x.OsqueryVersion
	}
	return nil
}

var File_protos_nightswatch_nightswatch_proto protoreflect.FileDescriptor

const file_protos_nightswatch_nightswatch_proto_rawDesc = "" +
	"\n" +
	"$protos/nightswatch/nightswatch.proto\x12\vnightswatch\x1a\x1bgoogle/protobuf/empty.proto\"W\n" +
	"\x0fRegisterRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"T\n" +
	"\fLoginRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"3\n" +
	"\x03App\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion\"5\n" +
	"\rInstalledApps\x12$\n" +
	"\x04apps\x18\x01 \x03(\v2\x10.nightswatch.AppR\x04apps\"%\n" +
	"\tOSVersion\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\"*\n" +
	"\x0eOSQueryVersion\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\"\xd3\x01\n" +
	"\x11DeviceDataRequest\x12A\n" +
	"\x0einstalled_apps\x18\x01 \x01(\v2\x1a.nightswatch.InstalledAppsR\rinstalledApps\x125\n" +
	"\n" +
	"os_version\x18\x02 \x01(\v2\x16.nightswatch.OSVersionR\tosVersion\x12D\n" +
	"\x0fosquery_version\x18\x03 \x01(\v2\x1b.nightswatch.OSQueryVersionR\x0eosqueryVersion\"~\n" +
	"\x14GetLatestDataRequest\x12\x1b\n" +
	"\tuser_name\x18\x01 \x01(\tR\buserName\x12I\n" +
	"\x12data_request_types\x18\x02 \x03(\x0e2\x1b.nightswatch.DeviceDataTypeR\x10dataRequestTypes\"\xd7\x01\n" +
	"\x15GetLatestDataResponse\x12A\n" +
	"\x0einstalled_apps\x18\x01 \x01(\v2\x1a.nightswatch.InstalledAppsR\rinstalledApps\x125\n" +
	"\n" +
	"os_version\x18\x02 \x01(\v2\x16.nightswatch.OSVersionR\tosVersion\x12D\n" +
	"\x0fosquery_version\x18\x03 \x01(\v2\x1b.nightswatch.OSQueryVersionR\x0eosqueryVersion*V\n" +
	"\x0eDeviceDataType\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\x12\n" +
	"\x0eINSTALLED_APPS\x10\x01\x12\x0e\n" +
	"\n" +
	"OS_VERSION\x10\x02\x12\x13\n" +
	"\x0fOSQUERY_VERSION\x10\x032\xbc\x02\n" +
	"\x12NightsWatchService\x12B\n" +
	"\bRegister\x12\x1c.nightswatch.RegisterRequest\x1a\x16.google.protobuf.Empty\"\x00\x12<\n" +
	"\x05Login\x12\x19.nightswatch.LoginRequest\x1a\x16.google.protobuf.Empty\"\x00\x12J\n" +
	"\x0eSendDeviceData\x12\x1e.nightswatch.DeviceDataRequest\x1a\x16.google.protobuf.Empty\"\x00\x12X\n" +
	"\rGetLatestData\x12!.nightswatch.GetLatestDataRequest\x1a\".nightswatch.GetLatestDataResponse\"\x00BGZEgithub.com/SoumyadipPayra/NightsWatchProtobufs/gogenproto/nightswatchb\x06proto3"

var (
	file_protos_nightswatch_nightswatch_proto_rawDescOnce sync.Once
	file_protos_nightswatch_nightswatch_proto_rawDescData []byte
)

func file_protos_nightswatch_nightswatch_proto_rawDescGZIP() []byte {
	file_protos_nightswatch_nightswatch_proto_rawDescOnce.Do(func() {
		file_protos_nightswatch_nightswatch_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_nightswatch_nightswatch_proto_rawDesc), len(file_protos_nightswatch_nightswatch_proto_rawDesc)))
	})
	return file_protos_nightswatch_nightswatch_proto_rawDescData
}

var file_protos_nightswatch_nightswatch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_nightswatch_nightswatch_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_nightswatch_nightswatch_proto_goTypes = []any{
	(DeviceDataType)(0),           // 0: nightswatch.DeviceDataType
	(*RegisterRequest)(nil),       // 1: nightswatch.RegisterRequest
	(*LoginRequest)(nil),          // 2: nightswatch.LoginRequest
	(*App)(nil),                   // 3: nightswatch.App
	(*InstalledApps)(nil),         // 4: nightswatch.InstalledApps
	(*OSVersion)(nil),             // 5: nightswatch.OSVersion
	(*OSQueryVersion)(nil),        // 6: nightswatch.OSQueryVersion
	(*DeviceDataRequest)(nil),     // 7: nightswatch.DeviceDataRequest
	(*GetLatestDataRequest)(nil),  // 8: nightswatch.GetLatestDataRequest
	(*GetLatestDataResponse)(nil), // 9: nightswatch.GetLatestDataResponse
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_protos_nightswatch_nightswatch_proto_depIdxs = []int32{
	3,  // 0: nightswatch.InstalledApps.apps:type_name -> nightswatch.App
	4,  // 1: nightswatch.DeviceDataRequest.installed_apps:type_name -> nightswatch.InstalledApps
	5,  // 2: nightswatch.DeviceDataRequest.os_version:type_name -> nightswatch.OSVersion
	6,  // 3: nightswatch.DeviceDataRequest.osquery_version:type_name -> nightswatch.OSQueryVersion
	0,  // 4: nightswatch.GetLatestDataRequest.data_request_types:type_name -> nightswatch.DeviceDataType
	4,  // 5: nightswatch.GetLatestDataResponse.installed_apps:type_name -> nightswatch.InstalledApps
	5,  // 6: nightswatch.GetLatestDataResponse.os_version:type_name -> nightswatch.OSVersion
	6,  // 7: nightswatch.GetLatestDataResponse.osquery_version:type_name -> nightswatch.OSQueryVersion
	1,  // 8: nightswatch.NightsWatchService.Register:input_type -> nightswatch.RegisterRequest
	2,  // 9: nightswatch.NightsWatchService.Login:input_type -> nightswatch.LoginRequest
	7,  // 10: nightswatch.NightsWatchService.SendDeviceData:input_type -> nightswatch.DeviceDataRequest
	8,  // 11: nightswatch.NightsWatchService.GetLatestData:input_type -> nightswatch.GetLatestDataRequest
	10, // 12: nightswatch.NightsWatchService.Register:output_type -> google.protobuf.Empty
	10, // 13: nightswatch.NightsWatchService.Login:output_type -> google.protobuf.Empty
	10, // 14: nightswatch.NightsWatchService.SendDeviceData:output_type -> google.protobuf.Empty
	9,  // 15: nightswatch.NightsWatchService.GetLatestData:output_type -> nightswatch.GetLatestDataResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_protos_nightswatch_nightswatch_proto_init() }
func file_protos_nightswatch_nightswatch_proto_init() {
	if File_protos_nightswatch_nightswatch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_nightswatch_nightswatch_proto_rawDesc), len(file_protos_nightswatch_nightswatch_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_nightswatch_nightswatch_proto_goTypes,
		DependencyIndexes: file_protos_nightswatch_nightswatch_proto_depIdxs,
		EnumInfos:         file_protos_nightswatch_nightswatch_proto_enumTypes,
		MessageInfos:      file_protos_nightswatch_nightswatch_proto_msgTypes,
	}.Build()
	File_protos_nightswatch_nightswatch_proto = out.File
	file_protos_nightswatch_nightswatch_proto_goTypes = nil
	file_protos_nightswatch_nightswatch_proto_depIdxs = nil
}
