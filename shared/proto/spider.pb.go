// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/x0rzkov/gcse/shared/proto/spider.proto

/*
Package gcsepb is a generated protocol buffer package.

It is generated from these files:
	github.com/x0rzkov/gcse/shared/proto/spider.proto
	github.com/x0rzkov/gcse/shared/proto/store.proto
	github.com/x0rzkov/gcse/shared/proto/stored.proto

It has these top-level messages:
	GoFileInfo
	RepoInfo
	FolderInfo
	CrawlingInfo
	HistoryEvent
	HistoryInfo
	Package
	PackageInfo
	PersonInfo
	Repository
	PackageCrawlHistoryReq
	PackageCrawlHistoryResp
*/
package gcsepb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GoFileInfo_Status int32

const (
	GoFileInfo_Unknown      GoFileInfo_Status = 0
	GoFileInfo_ParseSuccess GoFileInfo_Status = 1
	GoFileInfo_ParseFailed  GoFileInfo_Status = 2
	GoFileInfo_ShouldIgnore GoFileInfo_Status = 3
)

var GoFileInfo_Status_name = map[int32]string{
	0: "Unknown",
	1: "ParseSuccess",
	2: "ParseFailed",
	3: "ShouldIgnore",
}
var GoFileInfo_Status_value = map[string]int32{
	"Unknown":      0,
	"ParseSuccess": 1,
	"ParseFailed":  2,
	"ShouldIgnore": 3,
}

func (x GoFileInfo_Status) String() string {
	return proto.EnumName(GoFileInfo_Status_name, int32(x))
}
func (GoFileInfo_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type HistoryEvent_Action_Enum int32

const (
	HistoryEvent_Action_None    HistoryEvent_Action_Enum = 0
	HistoryEvent_Action_Success HistoryEvent_Action_Enum = 1
	HistoryEvent_Action_Failed  HistoryEvent_Action_Enum = 2
	HistoryEvent_Action_Invalid HistoryEvent_Action_Enum = 3
)

var HistoryEvent_Action_Enum_name = map[int32]string{
	0: "None",
	1: "Success",
	2: "Failed",
	3: "Invalid",
}
var HistoryEvent_Action_Enum_value = map[string]int32{
	"None":    0,
	"Success": 1,
	"Failed":  2,
	"Invalid": 3,
}

func (x HistoryEvent_Action_Enum) String() string {
	return proto.EnumName(HistoryEvent_Action_Enum_name, int32(x))
}
func (HistoryEvent_Action_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0, 0}
}

type GoFileInfo struct {
	Status      GoFileInfo_Status `protobuf:"varint,1,opt,name=status,enum=gcse.GoFileInfo_Status" json:"status,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	IsTest      bool              `protobuf:"varint,4,opt,name=is_test,json=isTest" json:"is_test,omitempty"`
	Imports     []string          `protobuf:"bytes,5,rep,name=imports" json:"imports,omitempty"`
}

func (m *GoFileInfo) Reset()                    { *m = GoFileInfo{} }
func (m *GoFileInfo) String() string            { return proto.CompactTextString(m) }
func (*GoFileInfo) ProtoMessage()               {}
func (*GoFileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GoFileInfo) GetStatus() GoFileInfo_Status {
	if m != nil {
		return m.Status
	}
	return GoFileInfo_Unknown
}

func (m *GoFileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GoFileInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GoFileInfo) GetIsTest() bool {
	if m != nil {
		return m.IsTest
	}
	return false
}

func (m *GoFileInfo) GetImports() []string {
	if m != nil {
		return m.Imports
	}
	return nil
}

type RepoInfo struct {
	// The timestamp this repo-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
	Stars        int32                      `protobuf:"varint,2,opt,name=stars" json:"stars,omitempty"`
	Description  string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Where this project was forked from, full path
	Source string `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	// As far as we know, when this repo was updated
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *RepoInfo) Reset()                    { *m = RepoInfo{} }
func (m *RepoInfo) String() string            { return proto.CompactTextString(m) }
func (*RepoInfo) ProtoMessage()               {}
func (*RepoInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepoInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

func (m *RepoInfo) GetStars() int32 {
	if m != nil {
		return m.Stars
	}
	return 0
}

func (m *RepoInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RepoInfo) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RepoInfo) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Information for a non-repository folder.
type FolderInfo struct {
	// E.g. "sub"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// E.g. "spider/sub"
	Path    string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Sha     string `protobuf:"bytes,3,opt,name=sha" json:"sha,omitempty"`
	HtmlUrl string `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	// The timestamp this folder-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
}

func (m *FolderInfo) Reset()                    { *m = FolderInfo{} }
func (m *FolderInfo) String() string            { return proto.CompactTextString(m) }
func (*FolderInfo) ProtoMessage()               {}
func (*FolderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FolderInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FolderInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FolderInfo) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *FolderInfo) GetHtmlUrl() string {
	if m != nil {
		return m.HtmlUrl
	}
	return ""
}

func (m *FolderInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

type CrawlingInfo struct {
	// The timestamp the related entry was crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
	Etag         string                     `protobuf:"bytes,2,opt,name=etag" json:"etag,omitempty"`
}

func (m *CrawlingInfo) Reset()                    { *m = CrawlingInfo{} }
func (m *CrawlingInfo) String() string            { return proto.CompactTextString(m) }
func (*CrawlingInfo) ProtoMessage()               {}
func (*CrawlingInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CrawlingInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

func (m *CrawlingInfo) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

type HistoryEvent struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Action    HistoryEvent_Action_Enum   `protobuf:"varint,2,opt,name=action,enum=gcse.HistoryEvent_Action_Enum" json:"action,omitempty"`
}

func (m *HistoryEvent) Reset()                    { *m = HistoryEvent{} }
func (m *HistoryEvent) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent) ProtoMessage()               {}
func (*HistoryEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HistoryEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *HistoryEvent) GetAction() HistoryEvent_Action_Enum {
	if m != nil {
		return m.Action
	}
	return HistoryEvent_Action_None
}

type HistoryEvent_Action struct {
}

func (m *HistoryEvent_Action) Reset()                    { *m = HistoryEvent_Action{} }
func (m *HistoryEvent_Action) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent_Action) ProtoMessage()               {}
func (*HistoryEvent_Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type HistoryInfo struct {
	Events    []*HistoryEvent            `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	FoundTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=found_time,json=foundTime" json:"found_time,omitempty"`
	// Possible value:
	//   web                  added from web
	//   user:<user>          found from user crawling
	//   parent               found by crawling his parent
	//   imported:<pkg>       imported by a <pkg>
	//   testimported:<pkg>   test imported by a <pkg>
	//   package:<pkg>
	//   reference:<pkg>      referenced in the readme file of <pkg>
	//   godoc                found by godoc.org/api
	FoundWay      string                     `protobuf:"bytes,3,opt,name=found_way,json=foundWay" json:"found_way,omitempty"`
	LatestSuccess *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=latest_success,json=latestSuccess" json:"latest_success,omitempty"`
	LatestFailed  *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=latest_failed,json=latestFailed" json:"latest_failed,omitempty"`
}

func (m *HistoryInfo) Reset()                    { *m = HistoryInfo{} }
func (m *HistoryInfo) String() string            { return proto.CompactTextString(m) }
func (*HistoryInfo) ProtoMessage()               {}
func (*HistoryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HistoryInfo) GetEvents() []*HistoryEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *HistoryInfo) GetFoundTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.FoundTime
	}
	return nil
}

func (m *HistoryInfo) GetFoundWay() string {
	if m != nil {
		return m.FoundWay
	}
	return ""
}

func (m *HistoryInfo) GetLatestSuccess() *google_protobuf.Timestamp {
	if m != nil {
		return m.LatestSuccess
	}
	return nil
}

func (m *HistoryInfo) GetLatestFailed() *google_protobuf.Timestamp {
	if m != nil {
		return m.LatestFailed
	}
	return nil
}

type Package struct {
	// package "name"
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	// Relative path to the repository, "" for root repository, "/sub" for a sub package.
	// Full path: site + "/" + user + "/" + repo + path
	Path        string `protobuf:"bytes,2,opt,name=Path" json:"Path,omitempty"`
	Synopsis    string `protobuf:"bytes,9,opt,name=Synopsis" json:"Synopsis,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description" json:"Description,omitempty"`
	// No directory info
	ReadmeFn string `protobuf:"bytes,4,opt,name=ReadmeFn" json:"ReadmeFn,omitempty"`
	// Raw content, cound be md, txt, etc.
	ReadmeData  string   `protobuf:"bytes,5,opt,name=ReadmeData" json:"ReadmeData,omitempty"`
	Imports     []string `protobuf:"bytes,6,rep,name=Imports" json:"Imports,omitempty"`
	TestImports []string `protobuf:"bytes,7,rep,name=TestImports" json:"TestImports,omitempty"`
	// URL to the package source code.
	Url string `protobuf:"bytes,8,opt,name=url" json:"url,omitempty"`
}

func (m *Package) Reset()                    { *m = Package{} }
func (m *Package) String() string            { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()               {}
func (*Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Package) GetSynopsis() string {
	if m != nil {
		return m.Synopsis
	}
	return ""
}

func (m *Package) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Package) GetReadmeFn() string {
	if m != nil {
		return m.ReadmeFn
	}
	return ""
}

func (m *Package) GetReadmeData() string {
	if m != nil {
		return m.ReadmeData
	}
	return ""
}

func (m *Package) GetImports() []string {
	if m != nil {
		return m.Imports
	}
	return nil
}

func (m *Package) GetTestImports() []string {
	if m != nil {
		return m.TestImports
	}
	return nil
}

func (m *Package) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*GoFileInfo)(nil), "gcse.GoFileInfo")
	proto.RegisterType((*RepoInfo)(nil), "gcse.RepoInfo")
	proto.RegisterType((*FolderInfo)(nil), "gcse.FolderInfo")
	proto.RegisterType((*CrawlingInfo)(nil), "gcse.CrawlingInfo")
	proto.RegisterType((*HistoryEvent)(nil), "gcse.HistoryEvent")
	proto.RegisterType((*HistoryEvent_Action)(nil), "gcse.HistoryEvent.Action")
	proto.RegisterType((*HistoryInfo)(nil), "gcse.HistoryInfo")
	proto.RegisterType((*Package)(nil), "gcse.Package")
	proto.RegisterEnum("gcse.GoFileInfo_Status", GoFileInfo_Status_name, GoFileInfo_Status_value)
	proto.RegisterEnum("gcse.HistoryEvent_Action_Enum", HistoryEvent_Action_Enum_name, HistoryEvent_Action_Enum_value)
}

func init() {
	proto.RegisterFile("github.com/x0rzkov/gcse/shared/proto/spider.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6a, 0xdb, 0x48,
	0x14, 0x8e, 0xfc, 0x23, 0xcb, 0x47, 0x4e, 0x56, 0x0c, 0xcb, 0x46, 0xeb, 0x85, 0x60, 0x74, 0x65,
	0xf6, 0x42, 0x82, 0x2c, 0x1b, 0x76, 0x59, 0x96, 0x36, 0x6d, 0xe2, 0xd6, 0xbd, 0x08, 0x41, 0x4e,
	0x28, 0xf4, 0xc6, 0x8c, 0xa5, 0xb1, 0x2c, 0x22, 0xcd, 0x08, 0xcd, 0x28, 0xc1, 0x0f, 0xd2, 0x17,
	0xe8, 0xa3, 0xf4, 0x31, 0xfa, 0x18, 0xbd, 0xeb, 0x5d, 0x99, 0x19, 0x29, 0x16, 0x04, 0x5a, 0x17,
	0x7a, 0x77, 0x7e, 0xbe, 0x99, 0x33, 0xe7, 0x3b, 0xe7, 0x1b, 0xf8, 0x3b, 0x49, 0xc5, 0xa6, 0x5a,
	0xf9, 0x11, 0xcb, 0x83, 0x18, 0xdf, 0xa7, 0x71, 0x4c, 0x68, 0x12, 0xd1, 0x20, 0x89, 0x38, 0x09,
	0xf8, 0x06, 0x97, 0x24, 0x0e, 0x8a, 0x92, 0x09, 0x16, 0xf0, 0x22, 0x8d, 0x49, 0xe9, 0x2b, 0x07,
	0xf5, 0x64, 0x7e, 0xfc, 0x5f, 0xeb, 0x70, 0xc2, 0x32, 0x4c, 0x13, 0x8d, 0x5d, 0x55, 0xeb, 0xa0,
	0x10, 0xdb, 0x82, 0xf0, 0x40, 0xa4, 0x39, 0xe1, 0x02, 0xe7, 0xc5, 0xce, 0xd2, 0x57, 0x78, 0x9f,
	0x0d, 0x80, 0x57, 0x6c, 0x96, 0x66, 0x64, 0x4e, 0xd7, 0x0c, 0x05, 0x60, 0x72, 0x81, 0x45, 0xc5,
	0x5d, 0x63, 0x62, 0x4c, 0x8f, 0x4e, 0x8f, 0x7d, 0x59, 0xc2, 0xdf, 0x21, 0xfc, 0x85, 0x4a, 0x87,
	0x35, 0x0c, 0x21, 0xe8, 0x51, 0x9c, 0x13, 0xb7, 0x33, 0x31, 0xa6, 0xc3, 0x50, 0xd9, 0x68, 0x02,
	0x76, 0x4c, 0x78, 0x54, 0xa6, 0x85, 0x48, 0x19, 0x75, 0xbb, 0x2a, 0xd5, 0x0e, 0xa1, 0x63, 0x18,
	0xa4, 0x7c, 0x29, 0x08, 0x17, 0x6e, 0x6f, 0x62, 0x4c, 0xad, 0xd0, 0x4c, 0xf9, 0x0d, 0xe1, 0x02,
	0xb9, 0x30, 0x48, 0xf3, 0x82, 0x95, 0x82, 0xbb, 0xfd, 0x49, 0x77, 0x3a, 0x0c, 0x1b, 0xd7, 0x7b,
	0x03, 0xa6, 0x2e, 0x8d, 0x6c, 0x18, 0xdc, 0xd2, 0x3b, 0xca, 0x1e, 0xa8, 0x73, 0x80, 0x1c, 0x18,
	0x5d, 0xe3, 0x92, 0x93, 0x45, 0x15, 0x45, 0x84, 0x73, 0xc7, 0x40, 0xbf, 0x80, 0xad, 0x22, 0x33,
	0x9c, 0x66, 0x24, 0x76, 0x3a, 0x12, 0xb2, 0xd8, 0xb0, 0x2a, 0x8b, 0xe7, 0x09, 0x65, 0x25, 0x71,
	0xba, 0xde, 0x27, 0x03, 0xac, 0x90, 0x14, 0x4c, 0xb5, 0xfc, 0x0c, 0x0e, 0xa3, 0x12, 0x3f, 0x64,
	0x29, 0x4d, 0x96, 0x92, 0x1d, 0xd5, 0xb9, 0x7d, 0x3a, 0xf6, 0x13, 0xc6, 0x92, 0x8c, 0xf8, 0x0d,
	0x97, 0xfe, 0x4d, 0x43, 0x5d, 0x38, 0x6a, 0x0e, 0xc8, 0x10, 0xfa, 0x15, 0xfa, 0x5c, 0xe0, 0x92,
	0x2b, 0x0e, 0xfa, 0xa1, 0x76, 0xf6, 0x20, 0xe1, 0x37, 0x30, 0x39, 0xab, 0xca, 0x88, 0xb8, 0x7d,
	0x95, 0xac, 0x3d, 0xf4, 0x3f, 0x8c, 0x32, 0xcc, 0xc5, 0xb2, 0x2a, 0x62, 0x2c, 0x48, 0xac, 0x18,
	0xfa, 0xf6, 0x7b, 0x6c, 0x89, 0xbf, 0xd5, 0x70, 0xef, 0x83, 0x01, 0x30, 0x63, 0x59, 0x4c, 0x4a,
	0xd5, 0x5e, 0x33, 0x20, 0xa3, 0x35, 0x20, 0x04, 0xbd, 0x02, 0x8b, 0x4d, 0x33, 0x34, 0x69, 0x23,
	0x07, 0xba, 0x7c, 0x83, 0xeb, 0x77, 0x4a, 0x13, 0xfd, 0x0e, 0xd6, 0x46, 0xe4, 0xd9, 0xb2, 0x2a,
	0x33, 0xf5, 0x86, 0x61, 0x38, 0x90, 0xfe, 0x6d, 0x99, 0x3d, 0xe5, 0xac, 0xff, 0x63, 0x9c, 0x79,
	0x11, 0x8c, 0x5e, 0xd6, 0xfe, 0xcf, 0x19, 0x02, 0x82, 0x1e, 0x11, 0x38, 0x69, 0x5a, 0x92, 0xb6,
	0xf7, 0xd1, 0x80, 0xd1, 0xeb, 0x94, 0x0b, 0x56, 0x6e, 0x2f, 0xef, 0x09, 0x15, 0xe8, 0x1f, 0x18,
	0x3e, 0xee, 0xff, 0x1e, 0x15, 0x76, 0x60, 0x74, 0x06, 0x26, 0x8e, 0xd4, 0x20, 0x3b, 0x4a, 0x17,
	0x27, 0x5a, 0x17, 0xed, 0xdb, 0xfd, 0x73, 0x05, 0xf0, 0x2f, 0x69, 0x95, 0x87, 0x35, 0x7a, 0xfc,
	0x1c, 0x4c, 0x1d, 0xf6, 0xce, 0xa0, 0x27, 0x33, 0xc8, 0x82, 0xde, 0x15, 0xa3, 0xc4, 0x39, 0x90,
	0x7b, 0xbc, 0xdb, 0x5a, 0x00, 0xf3, 0x71, 0x61, 0x6d, 0x18, 0xcc, 0xe9, 0x3d, 0xce, 0xd2, 0xd8,
	0xe9, 0x7a, 0xef, 0x3b, 0x60, 0xd7, 0x65, 0x14, 0x53, 0x7f, 0x82, 0x49, 0x64, 0x39, 0xa9, 0xd0,
	0xee, 0xd4, 0x3e, 0x45, 0x4f, 0x5f, 0x12, 0xd6, 0x08, 0xf4, 0x2f, 0xc0, 0x9a, 0x55, 0x34, 0xd6,
	0x94, 0x76, 0xbe, 0xdf, 0xb0, 0x42, 0x2b, 0x3e, 0xff, 0x00, 0xed, 0x2c, 0x1f, 0xf0, 0xb6, 0x5e,
	0x0a, 0x4b, 0x05, 0xde, 0xe2, 0x2d, 0x3a, 0x87, 0xa3, 0x0c, 0x4b, 0xf5, 0x2e, 0xb9, 0x6e, 0x60,
	0x8f, 0x1d, 0x3d, 0xd4, 0x27, 0xea, 0x8e, 0xe5, 0xc0, 0xeb, 0x2b, 0xd6, 0xaa, 0xed, 0x7d, 0x36,
	0x48, 0x1f, 0xd0, 0x34, 0x79, 0x5f, 0x0c, 0x18, 0x5c, 0xe3, 0xe8, 0x0e, 0x27, 0x6a, 0xf8, 0x57,
	0xad, 0x1d, 0xbf, 0xaa, 0x77, 0xfc, 0xba, 0xb5, 0xe3, 0xd2, 0x46, 0x63, 0xb0, 0x16, 0x5b, 0xca,
	0x0a, 0x9e, 0x72, 0x77, 0xa8, 0x7b, 0x6a, 0x7c, 0xa9, 0xd7, 0x8b, 0xa7, 0x7a, 0x6d, 0x85, 0xe4,
	0xe9, 0x90, 0xe0, 0x38, 0x27, 0x33, 0x5a, 0xeb, 0xe1, 0xd1, 0x47, 0x27, 0x00, 0xda, 0xbe, 0xc0,
	0x02, 0xd7, 0x7a, 0x6e, 0x45, 0xe4, 0xbf, 0x36, 0xaf, 0xff, 0x35, 0x53, 0xff, 0x6b, 0xb5, 0x2b,
	0xeb, 0xca, 0x9f, 0xaf, 0xc9, 0x0e, 0x54, 0xb6, 0x1d, 0x92, 0xca, 0x94, 0x12, 0xb4, 0xb4, 0x32,
	0xab, 0x32, 0x7b, 0x61, 0xbd, 0x33, 0xe5, 0xd0, 0x8b, 0xd5, 0xca, 0x54, 0x3c, 0xfd, 0xf5, 0x35,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x45, 0xb0, 0x53, 0x41, 0x06, 0x00, 0x00,
}
