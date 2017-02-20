package srv

type Modules interface {
	Process(data interface{}) interface{}
}
