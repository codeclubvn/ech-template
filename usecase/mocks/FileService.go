// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"
	model "trial_backend/domain/repo/model"

	mock "github.com/stretchr/testify/mock"

	request "trial_backend/presenter/request"
)

// FileService is an autogenerated mock type for the FileService type
type FileService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, req
func (_m *FileService) Delete(ctx context.Context, req request.DeleteFileRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.DeleteFileRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Download provides a mock function with given fields: ctx, id
func (_m *FileService) Download(ctx context.Context, id string) (*model.File, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.File
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.File, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.File); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.File)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: ctx, id
func (_m *FileService) GetOne(ctx context.Context, id string) (*model.File, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.File
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.File, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.File); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.File)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveFile provides a mock function with given fields: ctx, req
func (_m *FileService) SaveFile(ctx context.Context, req request.UploadFileRequest) (*model.File, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.File
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.UploadFileRequest) (*model.File, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.UploadFileRequest) *model.File); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.File)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.UploadFileRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, req
func (_m *FileService) Update(ctx context.Context, req request.UpdateFileRequest) (*model.File, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.File
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.UpdateFileRequest) (*model.File, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.UpdateFileRequest) *model.File); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.File)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.UpdateFileRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFileService creates a new instance of FileService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileService {
	mock := &FileService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}