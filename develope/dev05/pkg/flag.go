package pkg

import "strings"

type Core struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Phrase     string
	CountMatch int
}

func NewCore() *Core {
	return &Core{
		After:      0,
		Before:     0,
		Context:    0,
		Count:      false,
		IgnoreCase: false,
		Invert:     false,
		Fixed:      false,
		LineNum:    false,
		Phrase:     "",
		CountMatch: 0,
	}
}
func (c *Core) PhraseToLower() {
	c.Phrase = strings.ToLower(c.Phrase)
}
func (c *Core) SyncOutLength() {
	// ориентируемся по Before & After, если  Context больше - перезаписываем  After & Before
	if c.Context > c.After {
		c.After = c.Context
	}
	if c.Context > c.Before {
		c.Before = c.Context
	}
}
func (c *Core) AddMatch() {
	c.CountMatch++
}
