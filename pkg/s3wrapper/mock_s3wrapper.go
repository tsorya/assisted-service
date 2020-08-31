// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package s3wrapper is a generated GoMock package.
package s3wrapper

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	io "io"
	reflect "reflect"
	time "time"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// IsAwsS3 mocks base method
func (m *MockAPI) IsAwsS3() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAwsS3")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAwsS3 indicates an expected call of IsAwsS3
func (mr *MockAPIMockRecorder) IsAwsS3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAwsS3", reflect.TypeOf((*MockAPI)(nil).IsAwsS3))
}

// CreateBucket mocks base method
func (m *MockAPI) CreateBucket() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockAPIMockRecorder) CreateBucket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockAPI)(nil).CreateBucket))
}

// Upload mocks base method
func (m *MockAPI) Upload(ctx context.Context, data []byte, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, data, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockAPIMockRecorder) Upload(ctx, data, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockAPI)(nil).Upload), ctx, data, objectName)
}

// UploadStream mocks base method
func (m *MockAPI) UploadStream(ctx context.Context, reader io.Reader, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadStream", ctx, reader, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadStream indicates an expected call of UploadStream
func (mr *MockAPIMockRecorder) UploadStream(ctx, reader, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadStream", reflect.TypeOf((*MockAPI)(nil).UploadStream), ctx, reader, objectName)
}

// UploadFile mocks base method
func (m *MockAPI) UploadFile(ctx context.Context, filePath, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx, filePath, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile
func (mr *MockAPIMockRecorder) UploadFile(ctx, filePath, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockAPI)(nil).UploadFile), ctx, filePath, objectName)
}

// Download mocks base method
func (m *MockAPI) Download(ctx context.Context, objectName string) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", ctx, objectName)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Download indicates an expected call of Download
func (mr *MockAPIMockRecorder) Download(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockAPI)(nil).Download), ctx, objectName)
}

// DoesObjectExist mocks base method
func (m *MockAPI) DoesObjectExist(ctx context.Context, objectName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesObjectExist", ctx, objectName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesObjectExist indicates an expected call of DoesObjectExist
func (mr *MockAPIMockRecorder) DoesObjectExist(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesObjectExist", reflect.TypeOf((*MockAPI)(nil).DoesObjectExist), ctx, objectName)
}

// DeleteObject mocks base method
func (m *MockAPI) DeleteObject(ctx context.Context, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockAPIMockRecorder) DeleteObject(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockAPI)(nil).DeleteObject), ctx, objectName)
}

// GetObjectSizeBytes mocks base method
func (m *MockAPI) GetObjectSizeBytes(ctx context.Context, objectName string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectSizeBytes", ctx, objectName)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectSizeBytes indicates an expected call of GetObjectSizeBytes
func (mr *MockAPIMockRecorder) GetObjectSizeBytes(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectSizeBytes", reflect.TypeOf((*MockAPI)(nil).GetObjectSizeBytes), ctx, objectName)
}

// GeneratePresignedDownloadURL mocks base method
func (m *MockAPI) GeneratePresignedDownloadURL(ctx context.Context, objectName string, duration time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePresignedDownloadURL", ctx, objectName, duration)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePresignedDownloadURL indicates an expected call of GeneratePresignedDownloadURL
func (mr *MockAPIMockRecorder) GeneratePresignedDownloadURL(ctx, objectName, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePresignedDownloadURL", reflect.TypeOf((*MockAPI)(nil).GeneratePresignedDownloadURL), ctx, objectName, duration)
}

// UpdateObjectTimestamp mocks base method
func (m *MockAPI) UpdateObjectTimestamp(ctx context.Context, objectName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectTimestamp", ctx, objectName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateObjectTimestamp indicates an expected call of UpdateObjectTimestamp
func (mr *MockAPIMockRecorder) UpdateObjectTimestamp(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectTimestamp", reflect.TypeOf((*MockAPI)(nil).UpdateObjectTimestamp), ctx, objectName)
}

// ExpireObjects mocks base method
func (m *MockAPI) ExpireObjects(ctx context.Context, prefix string, deleteTime time.Duration, callback func(context.Context, logrus.FieldLogger, string)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExpireObjects", ctx, prefix, deleteTime, callback)
}

// ExpireObjects indicates an expected call of ExpireObjects
func (mr *MockAPIMockRecorder) ExpireObjects(ctx, prefix, deleteTime, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireObjects", reflect.TypeOf((*MockAPI)(nil).ExpireObjects), ctx, prefix, deleteTime, callback)
}

// ListObjectsByPrefix mocks base method
func (m *MockAPI) ListObjectsByPrefix(ctx context.Context, prefix string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsByPrefix", ctx, prefix)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectsByPrefix indicates an expected call of ListObjectsByPrefix
func (mr *MockAPIMockRecorder) ListObjectsByPrefix(ctx, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsByPrefix", reflect.TypeOf((*MockAPI)(nil).ListObjectsByPrefix), ctx, prefix)
}

// DownloadListOfFiles mocks base method
func (m *MockAPI) DownloadListOfFiles(ctx context.Context, files []string) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadListOfFiles", ctx, files)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadListOfFiles indicates an expected call of DownloadListOfFiles
func (mr *MockAPIMockRecorder) DownloadListOfFiles(ctx, files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadListOfFiles", reflect.TypeOf((*MockAPI)(nil).DownloadListOfFiles), ctx, files)
}
