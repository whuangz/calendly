package core

type Range struct {
	StartSecond, EndSecond int
}

// 75600 --> 21:00
func (r Range) Start() string {
	return r.translate(r.StartSecond)
}

func (r Range) EndStart() string {
	return r.translate(r.EndSecond)
}

func (r Range) translate(s int) string {
	return "00:00 "
}
