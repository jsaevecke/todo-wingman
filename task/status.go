package task

type Status int

const (
	UNKNOWN Status = iota
	TODO
	DONE
	INPROGRESS
)
