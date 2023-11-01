// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"fmt"
	"io"
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type AddPageBadRequest struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// GetField returns the value of Field.
func (s *AddPageBadRequest) GetField() string {
	return s.Field
}

// GetError returns the value of Error.
func (s *AddPageBadRequest) GetError() string {
	return s.Error
}

// SetField sets the value of Field.
func (s *AddPageBadRequest) SetField(val string) {
	s.Field = val
}

// SetError sets the value of Error.
func (s *AddPageBadRequest) SetError(val string) {
	s.Error = val
}

func (*AddPageBadRequest) addPageRes() {}

type AddPageReq struct {
	URL         string    `json:"url"`
	Description OptString `json:"description"`
	Formats     []Format  `json:"formats"`
}

// GetURL returns the value of URL.
func (s *AddPageReq) GetURL() string {
	return s.URL
}

// GetDescription returns the value of Description.
func (s *AddPageReq) GetDescription() OptString {
	return s.Description
}

// GetFormats returns the value of Formats.
func (s *AddPageReq) GetFormats() []Format {
	return s.Formats
}

// SetURL sets the value of URL.
func (s *AddPageReq) SetURL(val string) {
	s.URL = val
}

// SetDescription sets the value of Description.
func (s *AddPageReq) SetDescription(val OptString) {
	s.Description = val
}

// SetFormats sets the value of Formats.
func (s *AddPageReq) SetFormats(val []Format) {
	s.Formats = val
}

// Ref: #/components/schemas/error
type Error struct {
	Message   string    `json:"message"`
	Localized OptString `json:"localized"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// GetLocalized returns the value of Localized.
func (s *Error) GetLocalized() OptString {
	return s.Localized
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// SetLocalized sets the value of Localized.
func (s *Error) SetLocalized(val OptString) {
	s.Localized = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/format
type Format string

const (
	FormatAll        Format = "all"
	FormatPdf        Format = "pdf"
	FormatSingleFile Format = "single_file"
	FormatHeaders    Format = "headers"
)

// AllValues returns all Format values.
func (Format) AllValues() []Format {
	return []Format{
		FormatAll,
		FormatPdf,
		FormatSingleFile,
		FormatHeaders,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Format) MarshalText() ([]byte, error) {
	switch s {
	case FormatAll:
		return []byte(s), nil
	case FormatPdf:
		return []byte(s), nil
	case FormatSingleFile:
		return []byte(s), nil
	case FormatHeaders:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Format) UnmarshalText(data []byte) error {
	switch Format(data) {
	case FormatAll:
		*s = FormatAll
		return nil
	case FormatPdf:
		*s = FormatPdf
		return nil
	case FormatSingleFile:
		*s = FormatSingleFile
		return nil
	case FormatHeaders:
		*s = FormatHeaders
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// GetFileNotFound is response for GetFile operation.
type GetFileNotFound struct{}

func (*GetFileNotFound) getFileRes() {}

type GetFileOKApplicationPdf struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetFileOKApplicationPdf) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*GetFileOKApplicationPdf) getFileRes() {}

type GetFileOKTextHTML struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetFileOKTextHTML) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*GetFileOKTextHTML) getFileRes() {}

type GetFileOKTextPlain struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetFileOKTextPlain) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*GetFileOKTextPlain) getFileRes() {}

// GetPageNotFound is response for GetPage operation.
type GetPageNotFound struct{}

func (*GetPageNotFound) getPageRes() {}

// NewOptAddPageReq returns new OptAddPageReq with value set to v.
func NewOptAddPageReq(v AddPageReq) OptAddPageReq {
	return OptAddPageReq{
		Value: v,
		Set:   true,
	}
}

// OptAddPageReq is optional AddPageReq.
type OptAddPageReq struct {
	Value AddPageReq
	Set   bool
}

// IsSet returns true if OptAddPageReq was set.
func (o OptAddPageReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptAddPageReq) Reset() {
	var v AddPageReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptAddPageReq) SetTo(v AddPageReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptAddPageReq) Get() (v AddPageReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptAddPageReq) Or(d AddPageReq) AddPageReq {
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

// Ref: #/components/schemas/page
type Page struct {
	ID      uuid.UUID `json:"id"`
	URL     string    `json:"url"`
	Created time.Time `json:"created"`
	Formats []Format  `json:"formats"`
	Status  Status    `json:"status"`
	Meta    PageMeta  `json:"meta"`
}

// GetID returns the value of ID.
func (s *Page) GetID() uuid.UUID {
	return s.ID
}

// GetURL returns the value of URL.
func (s *Page) GetURL() string {
	return s.URL
}

// GetCreated returns the value of Created.
func (s *Page) GetCreated() time.Time {
	return s.Created
}

// GetFormats returns the value of Formats.
func (s *Page) GetFormats() []Format {
	return s.Formats
}

// GetStatus returns the value of Status.
func (s *Page) GetStatus() Status {
	return s.Status
}

// GetMeta returns the value of Meta.
func (s *Page) GetMeta() PageMeta {
	return s.Meta
}

// SetID sets the value of ID.
func (s *Page) SetID(val uuid.UUID) {
	s.ID = val
}

// SetURL sets the value of URL.
func (s *Page) SetURL(val string) {
	s.URL = val
}

// SetCreated sets the value of Created.
func (s *Page) SetCreated(val time.Time) {
	s.Created = val
}

// SetFormats sets the value of Formats.
func (s *Page) SetFormats(val []Format) {
	s.Formats = val
}

// SetStatus sets the value of Status.
func (s *Page) SetStatus(val Status) {
	s.Status = val
}

// SetMeta sets the value of Meta.
func (s *Page) SetMeta(val PageMeta) {
	s.Meta = val
}

func (*Page) addPageRes() {}

type PageMeta struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Error       OptString `json:"error"`
}

// GetTitle returns the value of Title.
func (s *PageMeta) GetTitle() string {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *PageMeta) GetDescription() string {
	return s.Description
}

// GetError returns the value of Error.
func (s *PageMeta) GetError() OptString {
	return s.Error
}

// SetTitle sets the value of Title.
func (s *PageMeta) SetTitle(val string) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *PageMeta) SetDescription(val string) {
	s.Description = val
}

// SetError sets the value of Error.
func (s *PageMeta) SetError(val OptString) {
	s.Error = val
}

// Merged schema.
// Ref: #/components/schemas/pageWithResults
type PageWithResults struct {
	ID      uuid.UUID           `json:"id"`
	URL     string              `json:"url"`
	Created time.Time           `json:"created"`
	Formats []Format            `json:"formats"`
	Status  Status              `json:"status"`
	Meta    PageWithResultsMeta `json:"meta"`
	Results []Result            `json:"results"`
}

// GetID returns the value of ID.
func (s *PageWithResults) GetID() uuid.UUID {
	return s.ID
}

// GetURL returns the value of URL.
func (s *PageWithResults) GetURL() string {
	return s.URL
}

// GetCreated returns the value of Created.
func (s *PageWithResults) GetCreated() time.Time {
	return s.Created
}

// GetFormats returns the value of Formats.
func (s *PageWithResults) GetFormats() []Format {
	return s.Formats
}

// GetStatus returns the value of Status.
func (s *PageWithResults) GetStatus() Status {
	return s.Status
}

// GetMeta returns the value of Meta.
func (s *PageWithResults) GetMeta() PageWithResultsMeta {
	return s.Meta
}

// GetResults returns the value of Results.
func (s *PageWithResults) GetResults() []Result {
	return s.Results
}

// SetID sets the value of ID.
func (s *PageWithResults) SetID(val uuid.UUID) {
	s.ID = val
}

// SetURL sets the value of URL.
func (s *PageWithResults) SetURL(val string) {
	s.URL = val
}

// SetCreated sets the value of Created.
func (s *PageWithResults) SetCreated(val time.Time) {
	s.Created = val
}

// SetFormats sets the value of Formats.
func (s *PageWithResults) SetFormats(val []Format) {
	s.Formats = val
}

// SetStatus sets the value of Status.
func (s *PageWithResults) SetStatus(val Status) {
	s.Status = val
}

// SetMeta sets the value of Meta.
func (s *PageWithResults) SetMeta(val PageWithResultsMeta) {
	s.Meta = val
}

// SetResults sets the value of Results.
func (s *PageWithResults) SetResults(val []Result) {
	s.Results = val
}

func (*PageWithResults) getPageRes() {}

type PageWithResultsMeta struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Error       OptString `json:"error"`
}

// GetTitle returns the value of Title.
func (s *PageWithResultsMeta) GetTitle() string {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *PageWithResultsMeta) GetDescription() string {
	return s.Description
}

// GetError returns the value of Error.
func (s *PageWithResultsMeta) GetError() OptString {
	return s.Error
}

// SetTitle sets the value of Title.
func (s *PageWithResultsMeta) SetTitle(val string) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *PageWithResultsMeta) SetDescription(val string) {
	s.Description = val
}

// SetError sets the value of Error.
func (s *PageWithResultsMeta) SetError(val OptString) {
	s.Error = val
}

type Pages []Page

// Ref: #/components/schemas/result
type Result struct {
	Format Format            `json:"format"`
	Error  OptString         `json:"error"`
	Files  []ResultFilesItem `json:"files"`
}

// GetFormat returns the value of Format.
func (s *Result) GetFormat() Format {
	return s.Format
}

// GetError returns the value of Error.
func (s *Result) GetError() OptString {
	return s.Error
}

// GetFiles returns the value of Files.
func (s *Result) GetFiles() []ResultFilesItem {
	return s.Files
}

// SetFormat sets the value of Format.
func (s *Result) SetFormat(val Format) {
	s.Format = val
}

// SetError sets the value of Error.
func (s *Result) SetError(val OptString) {
	s.Error = val
}

// SetFiles sets the value of Files.
func (s *Result) SetFiles(val []ResultFilesItem) {
	s.Files = val
}

type ResultFilesItem struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Mimetype string    `json:"mimetype"`
	Size     int64     `json:"size"`
}

// GetID returns the value of ID.
func (s *ResultFilesItem) GetID() uuid.UUID {
	return s.ID
}

// GetName returns the value of Name.
func (s *ResultFilesItem) GetName() string {
	return s.Name
}

// GetMimetype returns the value of Mimetype.
func (s *ResultFilesItem) GetMimetype() string {
	return s.Mimetype
}

// GetSize returns the value of Size.
func (s *ResultFilesItem) GetSize() int64 {
	return s.Size
}

// SetID sets the value of ID.
func (s *ResultFilesItem) SetID(val uuid.UUID) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *ResultFilesItem) SetName(val string) {
	s.Name = val
}

// SetMimetype sets the value of Mimetype.
func (s *ResultFilesItem) SetMimetype(val string) {
	s.Mimetype = val
}

// SetSize sets the value of Size.
func (s *ResultFilesItem) SetSize(val int64) {
	s.Size = val
}

// Ref: #/components/schemas/status
type Status string

const (
	StatusNew        Status = "new"
	StatusProcessing Status = "processing"
	StatusDone       Status = "done"
	StatusFailed     Status = "failed"
	StatusWithErrors Status = "with_errors"
)

// AllValues returns all Status values.
func (Status) AllValues() []Status {
	return []Status{
		StatusNew,
		StatusProcessing,
		StatusDone,
		StatusFailed,
		StatusWithErrors,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Status) MarshalText() ([]byte, error) {
	switch s {
	case StatusNew:
		return []byte(s), nil
	case StatusProcessing:
		return []byte(s), nil
	case StatusDone:
		return []byte(s), nil
	case StatusFailed:
		return []byte(s), nil
	case StatusWithErrors:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Status) UnmarshalText(data []byte) error {
	switch Status(data) {
	case StatusNew:
		*s = StatusNew
		return nil
	case StatusProcessing:
		*s = StatusProcessing
		return nil
	case StatusDone:
		*s = StatusDone
		return nil
	case StatusFailed:
		*s = StatusFailed
		return nil
	case StatusWithErrors:
		*s = StatusWithErrors
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
