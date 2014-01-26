package common

//	服务基本信息
type ServiceInfo struct {
	id   int
	name string
}

//	服务接口
type Service interface {
	Name() string
	Start(name string)
	Loop()
	Stop()

	//	同步的方式调用服务功能
	Call(name string, params ...interface{}) bool
}
