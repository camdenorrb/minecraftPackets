package common

type Bound uint8

const (
	ClientBound Bound = iota
	ServerBound
)

type State uint8

const (
	HandshakingState State = iota
	StatusState
	LoginState
	PlayState
)
