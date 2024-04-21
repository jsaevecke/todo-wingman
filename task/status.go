package task

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
	INPROGRESS
)
