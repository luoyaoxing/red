package common

import "fmt"

type ErrCodeMsg struct {
	code uint32
	msg string
}

var (
	ERR_SUCCESS = ErrCodeMsg{0, "success"}
	ERR_PARAM = ErrCodeMsg{10000, "param err"}
	ERR_OPERAROR = ErrCodeMsg{10001, "operator db err"}
	ERR_HAS_ACCOUNT = ErrCodeMsg{10002, "account has exist"}
	ERR_HAS_NOT_ACCOUNT = ErrCodeMsg{10003, "account has not exist"}
	ERR_TRANSFER_FAILED = ErrCodeMsg{10004, "transfer err"}
	ERR_NOT_BALANCE = ErrCodeMsg{10005, "余额不足"}
	ERR_NO_RED_ENVELOPE = ErrCodeMsg{10006, "没有足够的红包了"}
	ERR_HAS_NOT_RED_ENVELOPE = ErrCodeMsg{10007, "查不到红包"}
	ERR_HAS_NOT_RED_ENVELOPE_ITEM = ErrCodeMsg{10007, "查不到红包记录"}
)

func (e ErrCodeMsg) GetErrCode() uint32  {
	return e.code
}

func (e ErrCodeMsg) GetErrMsg() string  {
	return e.msg
}

func (e ErrCodeMsg) Error() string {
	return fmt.Sprintf("errCode:%d ErrMsg:%s", e.code, e.msg)
}
