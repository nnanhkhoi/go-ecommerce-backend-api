package response

const (
	ErrCodeSuccess      = 2001
	ErrCodeParamInvalid = 2003
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
}
