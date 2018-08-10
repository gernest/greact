package state

// State stores values to be used as state.
type State map[string]interface{}

// Merge returns a new state with key values from both a and b
func Merge(a, b State) State {
	m := make(State)
	for k, v := range a {
		m[k] = v
	}
	for k, v := range b {
		m[k] = v
	}
	return m
}
