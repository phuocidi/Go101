package services

type HelloService interface {
	SayHello() string
}

func NewHelloService() HelloService{
	return &helloService{}
}
type helloService struct {}

func (h *helloService) SayHello() string {
	return "hello, world!"
}
