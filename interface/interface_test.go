package interfaces

type Reader interface {
	DoReader() string
}

type Write interface {
	DoWrite() string
}

type WR interface {
	Reader
	Write
}

type WriteAndRead struct {
}

func (r *WriteAndRead) DoRead() string {
	return "this is doread"
}
