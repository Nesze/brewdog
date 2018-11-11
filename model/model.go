package model

// Beer ...
type Beer struct {
	Name    string
	Style   string
	Brewery string
	ABV     string
	Label   string
}

type Bar struct {
	ID    string
	Name  string
	OnTap []Beer
}

type Bars []Bar

func (b Bars) Len() int {
	return len(b)
}

func (b Bars) Less(i int, j int) bool {
	return b[i].Name <= b[j].Name
}

func (b Bars) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}
