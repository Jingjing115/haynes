package srv

import (
	"haynes/handle"
	"haynes/haynes"
	"haynes/data_receive"
	"haynes/config"
	"fmt"
)

type Srv struct {
	modules []Modules
}

func NewSrv(configFile string) *Srv {
	cfg := config.NewConfig(configFile)
	dc := cfg.DataReceiveConf
	mc := cfg.MsgDistributeConf
	fmt.Println(dc, mc)
	s := new(Srv)
	s.modules = []Modules{
		data_receive.NewDataReceive(),
		haynes.NewHaynes(),
		handle.NewHandle(),
	}
	return s
}

func (s *Srv) Start() error {
	return nil
}

func (s *Srv) Process() {
	var data interface{}
	for {
		for _, module := range s.modules {
			data = module.Process(data)
			if data == nil {
				break
			}
		}
	}
}
