package entity

type Rock struct {
	Value string
}

func NewRock() *Rock {
	return &Rock{Value: "rock"}
}

func (r Rock) GetValue() string {
	return r.Value
}
