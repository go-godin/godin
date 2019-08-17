package module

func Factory(moduleType Type) Module {
	switch moduleType {
	case TransportGrpcServer:
		return NewGrpcServerModule()
	default:
		return nil
	}
}
