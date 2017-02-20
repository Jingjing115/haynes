package models

func SeqCompare(left, right *Op) bool {
	return left.Seq == right.Seq
}
