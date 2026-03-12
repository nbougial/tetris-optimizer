package model

type Coordinate struct {
	Row int
	Col int
}

type Piece struct {
	Cells      []Coordinate
	LabelIndex int
}
