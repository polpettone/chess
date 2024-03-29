package model

type Pos struct {
	X int
	Y int
}

func PositionFromString(v string) *Pos {
	var x uint8
	var y uint8
	if len(v) != 2 {
		return nil
	}
	x = (string(v[0])[0]) - 65
	if x < 0 || x > 7 {
		return nil
	}
	y = (string(v[1])[0]) - 49
	if y < 0 || y > 7 {
		return nil
	}
	return NewPos(int(x), int(y))
}

func (p Pos) Print() string {
	return string(rune(p.X+65)) + string(rune(p.Y+49))
}

func (p Pos) String() string {
	return string(rune(p.X+65)) + string(rune(p.Y+49))
}

func NewPos(x, y int) *Pos {
	return &Pos{X: x, Y: y}
}
