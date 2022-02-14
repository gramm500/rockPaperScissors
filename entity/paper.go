package entity

type Paper struct {
	Value string
}

func NewPaper() *Paper {
	return &Paper{Value: "paper"}
}

func (p Paper) GetValue() string {
	return p.Value
}
