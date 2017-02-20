package handle

type Handle struct {

}

func NewHandle() *Handle {
	handle := new(Handle)
	return handle
}

func (h *Handle) Process(data interface{}) interface{} {
	if data == nil {
		return data
	}
	return data
}
