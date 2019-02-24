package minemap

// MineBlock ...
type MineBlock struct {
	Value  uint8 // 9 as a bumb, 11 as unknown(border condition)
	Status uint8
	User   string
}

// status
const (
	hidden = iota
	show
	flag
)
