package entity

type Scissors struct {
	Value string
}

func NewScissors() *Scissors {
	return &Scissors{Value: "scissors"}
}

func (s Scissors) GetValue() string {
	return s.Value
}
