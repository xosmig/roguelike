package status

// Status of the game
type Status int

// Available statuses. All except `Continue` mean that the game is over.
const (
	Continue Status = iota
	Victory
	Defeat
)
