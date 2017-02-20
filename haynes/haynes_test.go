package haynes

import (
	"testing"
	"haynes/models"
	"time"
	"fmt"
)

func newTestOp(addr models.DeviceAddrType, code string, timestamp int64) *models.Op{
	return &models.Op{
		DeviceAddr:addr,
		Code: code,
		Timestamp: timestamp,
	}
}

func TestHaynes_Process(t *testing.T) {
	h := NewHaynes("./rules/rules.json")
	time := time.Now().Unix()
	event := h.Process(newTestOp(12345, "rb", time))
	fmt.Println(event)
	event = h.Process(newTestOp(12345, "qi", time))
	fmt.Println(event)
	event = h.Process(newTestOp(12345, "in", time + 2))
	fmt.Println(event)
	event = h.Process(newTestOp(54321, "rb", time))
	fmt.Println(event)
	event = h.Process(newTestOp(54321, "md", time+3))
	fmt.Println(event)
	event = h.Process(newTestOp(54321, "mr", time+4))
	fmt.Println(event)
	event = h.Process(newTestOp(54321, "ro", time))
	fmt.Println(event)
}
