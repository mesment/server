package errno

/*错误码说明
    1   			00				02
服务级别错误       服务模块代码        具体错误代码

* 服务级别错误:1 为系统级错误;2 为普通错误，通常是由用 户非法操作引起的
* 服务模块为两位数:一个大型系统的服务模块通常不超过两位数，如果超过，说明这个系统该拆分了
* 错误码为两位数:防止一个模块定制过多的错误码，后期不好维护
* code = 0 说明是正确返回，code > 0 说明是错误返回 错误通常包括系统级错误码和服务级错误码
* 建议代码中按服务模块将错误分类
* 错误码均为 >= 0 的数
* 在 apiserver 中 HTTP Code 固定为 http.StatusOK，错误 码通过code 来表示。
*/

var (
	OK                  = &Errno{Code: 0, Msg: "OK"}
	InternalServerError = &Errno{Code: 10001, Msg: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Msg: "bind request to struct failed"}

	//用户错误
	ErrUserNameIsEmpty   = &Errno{Code: 20101, Msg: "用户名不能为空"}
	ErrPasswordIsEmpty   = &Errno{Code: 20102, Msg: "密码不能为空"}
	ErrUserNotExist      = &Errno{Code: 20103, Msg: "用户不存在"}
	ErrUserAlreadyExist  = &Errno{Code: 20104, Msg: "用户已存在"}
	ErrEncrypt           = &Errno{Code: 20101, Msg: "加密用户密码失败"}
	ErrTokenInvalid      = &Errno{Code: 20103, Msg: "token已失效"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Msg: "密码错误"}

	//服务错误
	ErrValidation = &Errno{Code: 20001, Msg: "验证失败"}
	ErrDatabase   = &Errno{Code: 20002, Msg: "数据库错误."}
	ErrToken      = &Errno{Code: 20003, Msg: "token校验失败"}
)
