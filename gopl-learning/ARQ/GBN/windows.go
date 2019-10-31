package windows

type Timer int
type Window struct {
	WinSize int
	Start   int
	End     int
	Ptr     int
	Status  []bool
	Timer   []Timer
}

func (p *Timer) init() {
	*p = 0
}
func (p *Window) Init(WinSize int) {
	p.WinSize = WinSize
	p.Start = 0
	p.End = WinSize - 1
	p.Ptr = 0
	p.Status = make([]bool, p.WinSize)
	p.Timer = make([]Timer, p.WinSize)
}

func (p *Window) Find(index int) (int, bool) {
	if index > p.End {
		return 0, false
	}
	return (p.Ptr + index - p.Start) % p.WinSize, true
}

func (p *Window) MoveOne() {
	p.Start++
	p.End++
	p.Status[p.Ptr] = false
	p.Timer[p.Ptr].init()
	p.Ptr = (p.Ptr + 1) % p.WinSize
}

func (p *Window) Move(n int) {
	for i := 0; i < n; i++ {
		p.MoveOne()
	}
}
