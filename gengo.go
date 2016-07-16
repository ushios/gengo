package gengo

import "time"

// Gengo 元号
type Gengo struct {
	start        time.Time
	japaneseName string
}

const (
	DateTimeFormat = "2006-01-02 15:04:05 -0700"
)

// NewGengoFromString is create gengo object.
func NewGengoFromString(start, jn string) Gengo {
	s, err := time.Parse(DateTimeFormat, start)
	if err != nil {
		panic(err)
	}
	return Gengo{
		start:        s,
		japaneseName: jn,
	}
}

// Start is gengo start time.
func (g Gengo) Start() time.Time {
	return g.start
}

// String for stringer.
func (g Gengo) String() string {
	return g.japaneseName
}

// Now return gengo now.
func Now() Gengo {
	return At(time.Now())
}

// At is get gengo at time.
func At(t time.Time) Gengo {
	for _, g := range gengos {
		if t.UnixNano() >= g.start.UnixNano() {
			return g
		}
	}

	panic("Gengo out of range")
}

var (
	gengos []Gengo
	Meiji  Gengo
	Taisho Gengo
	Showa  Gengo
	Heisei Gengo
)

func init() {
	Meiji = NewGengoFromString("1868-01-25 00:00:00 +0900", "明治")
	Taisho = NewGengoFromString("1912-07-30 00:00:00 +0900", "大正")
	Showa = NewGengoFromString("1926-12-25 00:00:00 +0900", "昭和")
	Heisei = NewGengoFromString("1989-01-08 00:00:00 +0900", "平成")

	gengos = []Gengo{
		Heisei, Showa, Taisho, Meiji,
	}
}
