package nv

type Response struct {
	ErrorCode    []byte
	ErrorMessage string
	DataLen      int
	Data         []byte
}

