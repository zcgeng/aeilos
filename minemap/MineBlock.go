package minemap

// MineBlock ...
type MineBlock struct {
	value  uint8 // 9 as a bumb, 11 as unknown(border condition)
	status uint8
	user   string
}

// status
const (
	hidden = iota
	show
	flag
)
