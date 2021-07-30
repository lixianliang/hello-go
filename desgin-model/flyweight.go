package main

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "俥",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
}

type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

func NewVChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

type ChessBoard struct {
	chessPiecs map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPiecs: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPiecs[id] = &ChessPiece{
			Unit: NewVChessPieceUnit(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.chessPiecs[id].X = x
	c.chessPiecs[id].Y = y
}

func main() {
	board1 := NewChessBoard()
	board1.Move(1, 1, 2)
	board2 := NewChessBoard()
	board2.Move(2, 2, 3)
}
