package comms

import (
	"github.com/stretchr/testify/mock"
)

// MockEventServer is a mock implementation of a datastore client for testing purposes
type MockEventServer struct {
	mock.Mock
}

func (m *MockEventServer) AddEventToSendQueue(data []byte) {
	m.Mock.Called(data)
}

func (m *MockEventServer) Close() {
	m.Mock.Called()
}
