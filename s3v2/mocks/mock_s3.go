// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/givetree/parquet-go-source/s3v2 (interfaces: S3API)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_s3.go -package=mocks github.com/givetree/parquet-go-source/s3v2 S3API
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "go.uber.org/mock/gomock"
)

// MockS3API is a mock of S3API interface.
type MockS3API struct {
	ctrl     *gomock.Controller
	recorder *MockS3APIMockRecorder
}

// MockS3APIMockRecorder is the mock recorder for MockS3API.
type MockS3APIMockRecorder struct {
	mock *MockS3API
}

// NewMockS3API creates a new mock instance.
func NewMockS3API(ctrl *gomock.Controller) *MockS3API {
	mock := &MockS3API{ctrl: ctrl}
	mock.recorder = &MockS3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3API) EXPECT() *MockS3APIMockRecorder {
	return m.recorder
}

// AbortMultipartUpload mocks base method.
func (m *MockS3API) AbortMultipartUpload(arg0 context.Context, arg1 *s3.AbortMultipartUploadInput, arg2 ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortMultipartUpload", varargs...)
	ret0, _ := ret[0].(*s3.AbortMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortMultipartUpload indicates an expected call of AbortMultipartUpload.
func (mr *MockS3APIMockRecorder) AbortMultipartUpload(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortMultipartUpload", reflect.TypeOf((*MockS3API)(nil).AbortMultipartUpload), varargs...)
}

// CompleteMultipartUpload mocks base method.
func (m *MockS3API) CompleteMultipartUpload(arg0 context.Context, arg1 *s3.CompleteMultipartUploadInput, arg2 ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteMultipartUpload", varargs...)
	ret0, _ := ret[0].(*s3.CompleteMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteMultipartUpload indicates an expected call of CompleteMultipartUpload.
func (mr *MockS3APIMockRecorder) CompleteMultipartUpload(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteMultipartUpload", reflect.TypeOf((*MockS3API)(nil).CompleteMultipartUpload), varargs...)
}

// CreateMultipartUpload mocks base method.
func (m *MockS3API) CreateMultipartUpload(arg0 context.Context, arg1 *s3.CreateMultipartUploadInput, arg2 ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMultipartUpload", varargs...)
	ret0, _ := ret[0].(*s3.CreateMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMultipartUpload indicates an expected call of CreateMultipartUpload.
func (mr *MockS3APIMockRecorder) CreateMultipartUpload(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultipartUpload", reflect.TypeOf((*MockS3API)(nil).CreateMultipartUpload), varargs...)
}

// GetObject mocks base method.
func (m *MockS3API) GetObject(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3APIMockRecorder) GetObject(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3API)(nil).GetObject), varargs...)
}

// HeadObject mocks base method.
func (m *MockS3API) HeadObject(arg0 context.Context, arg1 *s3.HeadObjectInput, arg2 ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HeadObject", varargs...)
	ret0, _ := ret[0].(*s3.HeadObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadObject indicates an expected call of HeadObject.
func (mr *MockS3APIMockRecorder) HeadObject(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadObject", reflect.TypeOf((*MockS3API)(nil).HeadObject), varargs...)
}

// PutObject mocks base method.
func (m *MockS3API) PutObject(arg0 context.Context, arg1 *s3.PutObjectInput, arg2 ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockS3APIMockRecorder) PutObject(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3API)(nil).PutObject), varargs...)
}

// UploadPart mocks base method.
func (m *MockS3API) UploadPart(arg0 context.Context, arg1 *s3.UploadPartInput, arg2 ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadPart", varargs...)
	ret0, _ := ret[0].(*s3.UploadPartOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPart indicates an expected call of UploadPart.
func (mr *MockS3APIMockRecorder) UploadPart(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPart", reflect.TypeOf((*MockS3API)(nil).UploadPart), varargs...)
}
