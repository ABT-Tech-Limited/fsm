package fsm

type TransitionType string

var TransitionTypes = struct {
	External TransitionType
	Internal TransitionType
}{
	External: "external",
	Internal: "internal",
}
