package interfaces

type HTTPServer interface {
	Notify() <-chan error
	Shutdown() error
}

type GRPCServer interface {
	Notify() <-chan error
	Shutdown() error
}
