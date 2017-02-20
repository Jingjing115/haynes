package rules

import (
	"haynes/models"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Rule struct {
	Name string	`json:"name"`
	Left string	`json:"left"`
	Right string	`json:"right"`
	Event string	`json:"event"`
	Block func(left, right *models.Op) bool	`json:"-"`

	// {|x, y| x.seq == y.seq}
	//func NewRules(x, y SeqInterface) bool {
	//		return x.seq == y.seq
	//}
	//func NewRules(x, y BulbInterface) bool {
	//	return x.hue == y.hue
	//}
}

type Rules struct {
	rules []Rule
}

func NewRules(rulesFile string) (error, *Rules) {
	var rules []Rule
	raw, err := ioutil.ReadFile(rulesFile)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	json.Unmarshal(raw, &rules)
	return nil, &Rules{
		rules:rules,
	}
}

func (r *Rules) Recognize(op *models.Op, opQueue []*models.Op) ([]*models.Op, string) {
	newQueue := make([]*models.Op, 0, len(opQueue))
	var found bool
	var delRight, delLeft *models.Op
	var event = ""
	for _, rule := range r.rules{
		if op.Code == rule.Right {
			if rule.Left == "" {
				// TODO 删除匹配到的指令
				delRight = op
				event = rule.Event
				found = true
			}
			for _, _op := range opQueue {
				if rule.Left == _op.Code {
					// TODO 删除匹配到的指令
					delRight = op
					delLeft = _op
					event = rule.Event
					found = true
				}
			}
		}
		if found {
			break
		}
	}

	for _, op := range opQueue {
		if !(op.Equal(delLeft) || op.Equal(delRight)) {
			newQueue = append(newQueue, op)
		}
	}

	return newQueue, event
}
