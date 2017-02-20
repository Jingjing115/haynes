package data_receive

type DataReceive struct {
	data chan interface{}
}

func NewDataReceive() *DataReceive {
	dr := new(DataReceive)
	dr.data = make(chan interface{}, 1)
	return dr
}

func (d *DataReceive) Process(data interface{}) interface{} {
	return d.data
}
