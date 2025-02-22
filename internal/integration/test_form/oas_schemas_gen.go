// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/google/uuid"

	ht "github.com/ogen-go/ogen/http"
)

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMultipartFile returns new OptMultipartFile with value set to v.
func NewOptMultipartFile(v ht.MultipartFile) OptMultipartFile {
	return OptMultipartFile{
		Value: v,
		Set:   true,
	}
}

// OptMultipartFile is optional ht.MultipartFile.
type OptMultipartFile struct {
	Value ht.MultipartFile
	Set   bool
}

// IsSet returns true if OptMultipartFile was set.
func (o OptMultipartFile) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMultipartFile) Reset() {
	var v ht.MultipartFile
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMultipartFile) SetTo(v ht.MultipartFile) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMultipartFile) Get() (v ht.MultipartFile, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMultipartFile) Or(d ht.MultipartFile) ht.MultipartFile {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTestFormDeepObject returns new OptTestFormDeepObject with value set to v.
func NewOptTestFormDeepObject(v TestFormDeepObject) OptTestFormDeepObject {
	return OptTestFormDeepObject{
		Value: v,
		Set:   true,
	}
}

// OptTestFormDeepObject is optional TestFormDeepObject.
type OptTestFormDeepObject struct {
	Value TestFormDeepObject
	Set   bool
}

// IsSet returns true if OptTestFormDeepObject was set.
func (o OptTestFormDeepObject) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTestFormDeepObject) Reset() {
	var v TestFormDeepObject
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTestFormDeepObject) SetTo(v TestFormDeepObject) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTestFormDeepObject) Get() (v TestFormDeepObject, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTestFormDeepObject) Or(d TestFormDeepObject) TestFormDeepObject {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTestFormObject returns new OptTestFormObject with value set to v.
func NewOptTestFormObject(v TestFormObject) OptTestFormObject {
	return OptTestFormObject{
		Value: v,
		Set:   true,
	}
}

// OptTestFormObject is optional TestFormObject.
type OptTestFormObject struct {
	Value TestFormObject
	Set   bool
}

// IsSet returns true if OptTestFormObject was set.
func (o OptTestFormObject) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTestFormObject) Reset() {
	var v TestFormObject
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTestFormObject) SetTo(v TestFormObject) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTestFormObject) Get() (v TestFormObject, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTestFormObject) Or(d TestFormObject) TestFormObject {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUUID returns new OptUUID with value set to v.
func NewOptUUID(v uuid.UUID) OptUUID {
	return OptUUID{
		Value: v,
		Set:   true,
	}
}

// OptUUID is optional uuid.UUID.
type OptUUID struct {
	Value uuid.UUID
	Set   bool
}

// IsSet returns true if OptUUID was set.
func (o OptUUID) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUUID) Reset() {
	var v uuid.UUID
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUUID) SetTo(v uuid.UUID) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUUID) Get() (v uuid.UUID, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUUID) Or(d uuid.UUID) uuid.UUID {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/SharedRequest
type SharedRequest struct {
	Filename OptString `json:"filename"`
	File     OptString `json:"file"`
}

// GetFilename returns the value of Filename.
func (s *SharedRequest) GetFilename() OptString {
	return s.Filename
}

// GetFile returns the value of File.
func (s *SharedRequest) GetFile() OptString {
	return s.File
}

// SetFilename sets the value of Filename.
func (s *SharedRequest) SetFilename(val OptString) {
	s.Filename = val
}

// SetFile sets the value of File.
func (s *SharedRequest) SetFile(val OptString) {
	s.File = val
}

func (*SharedRequest) testShareFormSchemaReq() {}

// Ref: #/components/schemas/SharedRequest
type SharedRequestForm struct {
	Filename OptString        `json:"filename"`
	File     OptMultipartFile `json:"file"`
}

// GetFilename returns the value of Filename.
func (s *SharedRequestForm) GetFilename() OptString {
	return s.Filename
}

// GetFile returns the value of File.
func (s *SharedRequestForm) GetFile() OptMultipartFile {
	return s.File
}

// SetFilename sets the value of Filename.
func (s *SharedRequestForm) SetFilename(val OptString) {
	s.Filename = val
}

// SetFile sets the value of File.
func (s *SharedRequestForm) SetFile(val OptMultipartFile) {
	s.File = val
}

func (*SharedRequestForm) testShareFormSchemaReq() {}

// Ref: #/components/schemas/TestForm
type TestForm struct {
	ID          OptInt                `json:"id"`
	UUID        OptUUID               `json:"uuid"`
	Description string                `json:"description"`
	Array       []string              `json:"array"`
	Object      OptTestFormObject     `json:"object"`
	DeepObject  OptTestFormDeepObject `json:"deepObject"`
}

// GetID returns the value of ID.
func (s *TestForm) GetID() OptInt {
	return s.ID
}

// GetUUID returns the value of UUID.
func (s *TestForm) GetUUID() OptUUID {
	return s.UUID
}

// GetDescription returns the value of Description.
func (s *TestForm) GetDescription() string {
	return s.Description
}

// GetArray returns the value of Array.
func (s *TestForm) GetArray() []string {
	return s.Array
}

// GetObject returns the value of Object.
func (s *TestForm) GetObject() OptTestFormObject {
	return s.Object
}

// GetDeepObject returns the value of DeepObject.
func (s *TestForm) GetDeepObject() OptTestFormDeepObject {
	return s.DeepObject
}

// SetID sets the value of ID.
func (s *TestForm) SetID(val OptInt) {
	s.ID = val
}

// SetUUID sets the value of UUID.
func (s *TestForm) SetUUID(val OptUUID) {
	s.UUID = val
}

// SetDescription sets the value of Description.
func (s *TestForm) SetDescription(val string) {
	s.Description = val
}

// SetArray sets the value of Array.
func (s *TestForm) SetArray(val []string) {
	s.Array = val
}

// SetObject sets the value of Object.
func (s *TestForm) SetObject(val OptTestFormObject) {
	s.Object = val
}

// SetDeepObject sets the value of DeepObject.
func (s *TestForm) SetDeepObject(val OptTestFormDeepObject) {
	s.DeepObject = val
}

type TestFormDeepObject struct {
	Min OptInt `json:"min"`
	Max int    `json:"max"`
}

// GetMin returns the value of Min.
func (s *TestFormDeepObject) GetMin() OptInt {
	return s.Min
}

// GetMax returns the value of Max.
func (s *TestFormDeepObject) GetMax() int {
	return s.Max
}

// SetMin sets the value of Min.
func (s *TestFormDeepObject) SetMin(val OptInt) {
	s.Min = val
}

// SetMax sets the value of Max.
func (s *TestFormDeepObject) SetMax(val int) {
	s.Max = val
}

type TestFormObject struct {
	Min OptInt `json:"min"`
	Max int    `json:"max"`
}

// GetMin returns the value of Min.
func (s *TestFormObject) GetMin() OptInt {
	return s.Min
}

// GetMax returns the value of Max.
func (s *TestFormObject) GetMax() int {
	return s.Max
}

// SetMin sets the value of Min.
func (s *TestFormObject) SetMin(val OptInt) {
	s.Min = val
}

// SetMax sets the value of Max.
func (s *TestFormObject) SetMax(val int) {
	s.Max = val
}

// TestFormURLEncodedOK is response for TestFormURLEncoded operation.
type TestFormURLEncodedOK struct{}

// TestMultipartOK is response for TestMultipart operation.
type TestMultipartOK struct{}

type TestMultipartUploadOK struct {
	File         string    `json:"file"`
	OptionalFile OptString `json:"optional_file"`
	Files        []string  `json:"files"`
}

// GetFile returns the value of File.
func (s *TestMultipartUploadOK) GetFile() string {
	return s.File
}

// GetOptionalFile returns the value of OptionalFile.
func (s *TestMultipartUploadOK) GetOptionalFile() OptString {
	return s.OptionalFile
}

// GetFiles returns the value of Files.
func (s *TestMultipartUploadOK) GetFiles() []string {
	return s.Files
}

// SetFile sets the value of File.
func (s *TestMultipartUploadOK) SetFile(val string) {
	s.File = val
}

// SetOptionalFile sets the value of OptionalFile.
func (s *TestMultipartUploadOK) SetOptionalFile(val OptString) {
	s.OptionalFile = val
}

// SetFiles sets the value of Files.
func (s *TestMultipartUploadOK) SetFiles(val []string) {
	s.Files = val
}

type TestMultipartUploadReq struct {
	OrderId      OptInt    `json:"orderId"`
	UserId       OptInt    `json:"userId"`
	File         string    `json:"file"`
	OptionalFile OptString `json:"optional_file"`
	Files        []string  `json:"files"`
}

// GetOrderId returns the value of OrderId.
func (s *TestMultipartUploadReq) GetOrderId() OptInt {
	return s.OrderId
}

// GetUserId returns the value of UserId.
func (s *TestMultipartUploadReq) GetUserId() OptInt {
	return s.UserId
}

// GetFile returns the value of File.
func (s *TestMultipartUploadReq) GetFile() string {
	return s.File
}

// GetOptionalFile returns the value of OptionalFile.
func (s *TestMultipartUploadReq) GetOptionalFile() OptString {
	return s.OptionalFile
}

// GetFiles returns the value of Files.
func (s *TestMultipartUploadReq) GetFiles() []string {
	return s.Files
}

// SetOrderId sets the value of OrderId.
func (s *TestMultipartUploadReq) SetOrderId(val OptInt) {
	s.OrderId = val
}

// SetUserId sets the value of UserId.
func (s *TestMultipartUploadReq) SetUserId(val OptInt) {
	s.UserId = val
}

// SetFile sets the value of File.
func (s *TestMultipartUploadReq) SetFile(val string) {
	s.File = val
}

// SetOptionalFile sets the value of OptionalFile.
func (s *TestMultipartUploadReq) SetOptionalFile(val OptString) {
	s.OptionalFile = val
}

// SetFiles sets the value of Files.
func (s *TestMultipartUploadReq) SetFiles(val []string) {
	s.Files = val
}

type TestMultipartUploadReqForm struct {
	OrderId      OptInt             `json:"orderId"`
	UserId       OptInt             `json:"userId"`
	File         ht.MultipartFile   `json:"file"`
	OptionalFile OptMultipartFile   `json:"optional_file"`
	Files        []ht.MultipartFile `json:"files"`
}

// GetOrderId returns the value of OrderId.
func (s *TestMultipartUploadReqForm) GetOrderId() OptInt {
	return s.OrderId
}

// GetUserId returns the value of UserId.
func (s *TestMultipartUploadReqForm) GetUserId() OptInt {
	return s.UserId
}

// GetFile returns the value of File.
func (s *TestMultipartUploadReqForm) GetFile() ht.MultipartFile {
	return s.File
}

// GetOptionalFile returns the value of OptionalFile.
func (s *TestMultipartUploadReqForm) GetOptionalFile() OptMultipartFile {
	return s.OptionalFile
}

// GetFiles returns the value of Files.
func (s *TestMultipartUploadReqForm) GetFiles() []ht.MultipartFile {
	return s.Files
}

// SetOrderId sets the value of OrderId.
func (s *TestMultipartUploadReqForm) SetOrderId(val OptInt) {
	s.OrderId = val
}

// SetUserId sets the value of UserId.
func (s *TestMultipartUploadReqForm) SetUserId(val OptInt) {
	s.UserId = val
}

// SetFile sets the value of File.
func (s *TestMultipartUploadReqForm) SetFile(val ht.MultipartFile) {
	s.File = val
}

// SetOptionalFile sets the value of OptionalFile.
func (s *TestMultipartUploadReqForm) SetOptionalFile(val OptMultipartFile) {
	s.OptionalFile = val
}

// SetFiles sets the value of Files.
func (s *TestMultipartUploadReqForm) SetFiles(val []ht.MultipartFile) {
	s.Files = val
}

// TestShareFormSchemaOK is response for TestShareFormSchema operation.
type TestShareFormSchemaOK struct{}
