package errno

import "fmt"

type Errno struct {
	Code int    //错误码
	Msg  string //错误信息
}

func (e Errno) Error() string {
	return e.Msg
}

//Err 表示一个错误
type Err struct {
	Code int
	Msg  string
	Err  error
}

//新建自定义的错误
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Msg: errno.Msg, Err: err}
}

func (e *Err) Add(msg string) error {
	e.Msg = e.Msg + msg
	return e
}

func (e *Err) Addf(format string, args ...interface{}) error {
	e.Msg = e.Msg + " " + fmt.Sprintf(format, args...)
	return e
}

func (e *Err) Error() string {
	return fmt.Sprintf("code:%d,message:%s, error:%s\n", e.Code, e.Msg, e.Err)
}

//解析自定义的错误信息
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Msg
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Msg
	case *Errno:
		return typed.Code, typed.Msg
	default:
	}

	return InternalServerError.Code, InternalServerError.Msg
}

func IsErrUserNotFount(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotExist.Code
}
