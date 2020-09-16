package consts

type Status struct {
	Code int32
	Msg  string
}

var (
	OK                = &Status{Code: 0, Msg: "success"}
	ERR_ILLEGAL_PARAM = &Status{Code: 400, Msg: "illegal param"}
	ERR_REDIS         = &Status{Code: 401, Msg: "redis err"}
	ERR_UNDEFINED     = &Status{Code: 402, Msg: "unknown err"}
)

func (s Status) GetCode() int32 {
	return s.Code
}

func (s Status) GetError() (int32, string) {
	return s.Code, s.Msg
}

func (s Status) GetMsg() string {
	return s.Msg
}
