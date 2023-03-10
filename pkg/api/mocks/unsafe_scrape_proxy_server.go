// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeScrapeProxyServer is an autogenerated mock type for the UnsafeScrapeProxyServer type
type UnsafeScrapeProxyServer struct {
	mock.Mock
}

type UnsafeScrapeProxyServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeScrapeProxyServer) EXPECT() *UnsafeScrapeProxyServer_Expecter {
	return &UnsafeScrapeProxyServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedScrapeProxyServer provides a mock function with given fields:
func (_m *UnsafeScrapeProxyServer) mustEmbedUnimplementedScrapeProxyServer() {
	_m.Called()
}

// UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedScrapeProxyServer'
type UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedScrapeProxyServer is a helper method to define mock.On call
func (_e *UnsafeScrapeProxyServer_Expecter) mustEmbedUnimplementedScrapeProxyServer() *UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call {
	return &UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call{Call: _e.mock.On("mustEmbedUnimplementedScrapeProxyServer")}
}

func (_c *UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call) Run(run func()) *UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call) Return() *UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call) RunAndReturn(run func()) *UnsafeScrapeProxyServer_mustEmbedUnimplementedScrapeProxyServer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUnsafeScrapeProxyServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeScrapeProxyServer creates a new instance of UnsafeScrapeProxyServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeScrapeProxyServer(t mockConstructorTestingTNewUnsafeScrapeProxyServer) *UnsafeScrapeProxyServer {
	mock := &UnsafeScrapeProxyServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
