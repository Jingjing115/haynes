package haynes

import (
	"haynes/models"
	"haynes/haynes/rules"
	"log"
	"fmt"
)

type Haynes struct {
	data    chan interface{}
	rules   map[models.DeviceAddrType]*rules.Rules
	opQueue map[models.DeviceAddrType][]*models.Op
	ruleFile string
}

func NewHaynes(ruleFile string) *Haynes {
	return &Haynes{
		rules:  make(map[models.DeviceAddrType]*rules.Rules, 1024),
		opQueue: make(map[models.DeviceAddrType][]*models.Op, 1024),
		ruleFile: ruleFile,
	}
}

func (h *Haynes) Process(data interface{}) interface{} {
	var err error
	if data == nil {
		return data
	}

	op, ok := data.(*models.Op);
	if !ok {
		return nil
	}

	if _, ok := h.opQueue[op.DeviceAddr]; !ok {
		h.opQueue[op.DeviceAddr] = make([]*models.Op, 0, 32)
	}

	if _, ok := h.rules[op.DeviceAddr]; !ok {
		if err, h.rules[op.DeviceAddr] = rules.NewRules(h.ruleFile); err != nil {
			log.Println(err)
		}
	}

	// 处理数据后返回, 构造识别的事件
	event := h.recognize(op)
	if event == nil {
		return nil
	}
	msg := fmt.Sprintf(`{op: %s, event: %q}`, op, event)

	return msg
}

func (h *Haynes) recognize(op *models.Op) interface{} {
	opQueue := h.opQueue[op.DeviceAddr]
	newQueue := make([]*models.Op, 0, len(opQueue))
	// 移除队列中超出时间的指令
	for _, _op := range opQueue {
		if (op.Timestamp - _op.Timestamp) < 5 {
			newQueue = append(newQueue, _op)
		}
	}
	newQueue = append(newQueue, op)

	queue, event := h.rules[op.DeviceAddr].Recognize(op, newQueue)

	// 修改指令队列
	h.opQueue[op.DeviceAddr] = queue

	return event
}
