package models

import "encoding/json"

type DeviceAddrType uint32

type Op struct {
	Code       string
	DeviceAddr DeviceAddrType
	Timestamp  int64
	Seq        int
}

func (op *Op) String() string {
	data, err := json.Marshal(op)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (op *Op) Equal(ano *Op) bool {
	if op == nil || ano == nil {
		return false
	} else {
		return op.Code == ano.Code &&
			op.DeviceAddr == ano.DeviceAddr &&
			op.Timestamp == ano.Timestamp
	}
}
