package request

type ServiceMock struct{}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (s *ServiceMock) Get(address string) []byte {
	return []byte("test request body " + address)
}
