package minemap

const ROW_LENGTH = 40 // should be consistant with frontend
const ROW_HEIGHT = 25

const SCORE_WRONG_FLAG = -10
const SCORE_RIGHT_FLAG = 5
const SCORE_WRONG_FLIP = -30
const SCORE_RIGHT_FLIP = 1

func block2area(x, y int) (int, int) {
	ax := x / 100
	ay := y / 100
	if x < 0 {
		ax = (x+1)/100 - 1
	}
	if y < 0 {
		ay = (y+1)/100 - 1
	}
	return ax, ay
}

func area2block(x, y int) (int, int) {
	return x * 100, y * 100
}
