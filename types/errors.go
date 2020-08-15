package types

const (
	ErrorCodeMarshal = iota + 101
	ErrorCodeUnmarshal
)

const (
	ErrorMsgMarshal   = "error occurred while marshalling the data"
	ErrorMsgUnmarshal = "error occurred while unmarshalling the data"
)
