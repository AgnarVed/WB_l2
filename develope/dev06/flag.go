package main

type Core struct {
	Fields    int
	Delimiter string
	Separated bool
}

func NewCore() *Core {
	return &Core{
		Fields:    0,
		Delimiter: "",
		Separated: false,
	}
}
