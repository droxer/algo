package recursive

type Trampoline interface {
	Run() int
}

type Done struct {
	value int
}

func (d Done) Run() int {
	return d.value
}

type Call struct {
	next func() Trampoline
}

func (c Call) Run() int {
	t := c
	for {
		switch step := t.next().(type) {
		case Done:
			return step.value
		case Call:
			t = step
		}
	}
}
