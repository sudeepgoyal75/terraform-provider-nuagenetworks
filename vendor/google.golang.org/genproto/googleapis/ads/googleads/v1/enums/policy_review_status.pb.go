// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/policy_review_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The possible policy review statuses.
type PolicyReviewStatusEnum_PolicyReviewStatus int32

const (
	// No value has been specified.
	PolicyReviewStatusEnum_UNSPECIFIED PolicyReviewStatusEnum_PolicyReviewStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyReviewStatusEnum_UNKNOWN PolicyReviewStatusEnum_PolicyReviewStatus = 1
	// Currently under review.
	PolicyReviewStatusEnum_REVIEW_IN_PROGRESS PolicyReviewStatusEnum_PolicyReviewStatus = 2
	// Primary review complete. Other reviews may be continuing.
	PolicyReviewStatusEnum_REVIEWED PolicyReviewStatusEnum_PolicyReviewStatus = 3
	// The resource has been resubmitted for approval or its policy decision has
	// been appealed.
	PolicyReviewStatusEnum_UNDER_APPEAL PolicyReviewStatusEnum_PolicyReviewStatus = 4
)

var PolicyReviewStatusEnum_PolicyReviewStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REVIEW_IN_PROGRESS",
	3: "REVIEWED",
	4: "UNDER_APPEAL",
}

var PolicyReviewStatusEnum_PolicyReviewStatus_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"REVIEW_IN_PROGRESS": 2,
	"REVIEWED":           3,
	"UNDER_APPEAL":       4,
}

func (x PolicyReviewStatusEnum_PolicyReviewStatus) String() string {
	return proto.EnumName(PolicyReviewStatusEnum_PolicyReviewStatus_name, int32(x))
}

func (PolicyReviewStatusEnum_PolicyReviewStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_07bf305fecc00f04, []int{0, 0}
}

// Container for enum describing possible policy review statuses.
type PolicyReviewStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyReviewStatusEnum) Reset()         { *m = PolicyReviewStatusEnum{} }
func (m *PolicyReviewStatusEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyReviewStatusEnum) ProtoMessage()    {}
func (*PolicyReviewStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_07bf305fecc00f04, []int{0}
}

func (m *PolicyReviewStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyReviewStatusEnum.Unmarshal(m, b)
}
func (m *PolicyReviewStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyReviewStatusEnum.Marshal(b, m, deterministic)
}
func (m *PolicyReviewStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyReviewStatusEnum.Merge(m, src)
}
func (m *PolicyReviewStatusEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyReviewStatusEnum.Size(m)
}
func (m *PolicyReviewStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyReviewStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyReviewStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.PolicyReviewStatusEnum_PolicyReviewStatus", PolicyReviewStatusEnum_PolicyReviewStatus_name, PolicyReviewStatusEnum_PolicyReviewStatus_value)
	proto.RegisterType((*PolicyReviewStatusEnum)(nil), "google.ads.googleads.v1.enums.PolicyReviewStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/policy_review_status.proto", fileDescriptor_07bf305fecc00f04)
}

var fileDescriptor_07bf305fecc00f04 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xb6, 0x9d, 0xa8, 0x64, 0x03, 0x4b, 0x0e, 0x13, 0xc4, 0x1d, 0xb6, 0x07, 0x48, 0x29, 0x5e,
	0x24, 0x9e, 0x32, 0x17, 0x47, 0x51, 0xba, 0xd0, 0xb2, 0x0e, 0xa4, 0x50, 0xea, 0x5a, 0x42, 0x65,
	0x4b, 0xca, 0xd2, 0x4d, 0xbc, 0xfb, 0x24, 0x1e, 0x7d, 0x14, 0x1f, 0x45, 0x7c, 0x08, 0x69, 0xb2,
	0xed, 0x32, 0xf4, 0x12, 0x3e, 0xf2, 0xfd, 0xe1, 0xf7, 0x7d, 0xe0, 0x86, 0x4b, 0xc9, 0x17, 0x85,
	0x9b, 0xe5, 0xca, 0x35, 0xb0, 0x41, 0x1b, 0xcf, 0x2d, 0xc4, 0x7a, 0xa9, 0xdc, 0x4a, 0x2e, 0xca,
	0xf9, 0x5b, 0xba, 0x2a, 0x36, 0x65, 0xf1, 0x9a, 0xaa, 0x3a, 0xab, 0xd7, 0x0a, 0x55, 0x2b, 0x59,
	0x4b, 0xd8, 0x33, 0x72, 0x94, 0xe5, 0x0a, 0xed, 0x9d, 0x68, 0xe3, 0x21, 0xed, 0xbc, 0xbc, 0xda,
	0x05, 0x57, 0xa5, 0x9b, 0x09, 0x21, 0xeb, 0xac, 0x2e, 0xa5, 0xd8, 0x9a, 0x07, 0xef, 0x16, 0xe8,
	0x32, 0x9d, 0x1d, 0xea, 0xe8, 0x48, 0x27, 0x53, 0xb1, 0x5e, 0x0e, 0x5e, 0x00, 0x3c, 0x64, 0xe0,
	0x39, 0x68, 0x4f, 0x83, 0x88, 0xd1, 0x3b, 0xff, 0xde, 0xa7, 0x23, 0xe7, 0x08, 0xb6, 0xc1, 0xe9,
	0x34, 0x78, 0x08, 0x26, 0xb3, 0xc0, 0xb1, 0x60, 0x17, 0xc0, 0x90, 0xc6, 0x3e, 0x9d, 0xa5, 0x7e,
	0x90, 0xb2, 0x70, 0x32, 0x0e, 0x69, 0x14, 0x39, 0x36, 0xec, 0x80, 0x33, 0xf3, 0x4f, 0x47, 0x4e,
	0x0b, 0x3a, 0xa0, 0x33, 0x0d, 0x46, 0x34, 0x4c, 0x09, 0x63, 0x94, 0x3c, 0x3a, 0xc7, 0xc3, 0x1f,
	0x0b, 0xf4, 0xe7, 0x72, 0x89, 0xfe, 0xad, 0x32, 0xbc, 0x38, 0xbc, 0x87, 0x35, 0x2d, 0x98, 0xf5,
	0x34, 0xdc, 0x3a, 0xb9, 0x5c, 0x64, 0x82, 0x23, 0xb9, 0xe2, 0x2e, 0x2f, 0x84, 0xee, 0xb8, 0x9b,
	0xb3, 0x2a, 0xd5, 0x1f, 0xeb, 0xde, 0xea, 0xf7, 0xc3, 0x6e, 0x8d, 0x09, 0xf9, 0xb4, 0x7b, 0x63,
	0x13, 0x45, 0x72, 0x85, 0x0c, 0x6c, 0x50, 0xec, 0xa1, 0x66, 0x15, 0xf5, 0xb5, 0xe3, 0x13, 0x92,
	0xab, 0x64, 0xcf, 0x27, 0xb1, 0x97, 0x68, 0xfe, 0xdb, 0xee, 0x9b, 0x4f, 0x8c, 0x49, 0xae, 0x30,
	0xde, 0x2b, 0x30, 0x8e, 0x3d, 0x8c, 0xb5, 0xe6, 0xf9, 0x44, 0x1f, 0x76, 0xfd, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x85, 0xab, 0xcf, 0x9a, 0xf5, 0x01, 0x00, 0x00,
}