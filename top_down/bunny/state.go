package bunny

// State represents the state of the bunny.
type State int

const (
	idle_front State = iota
	run_front State = iota
	idle_back State = iota
	run_back State = iota
	idle_left State = iota
	run_left State = iota
	idle_right State = iota
	run_right State = iota
)
