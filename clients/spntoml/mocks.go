package spntoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable spntoml client.
type MockClient struct {
	mock.Mock
}

// GetSpnToml is a mocking a method
func (m *MockClient) GetSpnToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetSpnTomlByAddress is a mocking a method
func (m *MockClient) GetSpnTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
