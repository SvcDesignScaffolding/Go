package util

// 定义一个通用的错误类型
type Error struct {
    // 错误信息
    Message string
}

// 返回一个错误
func NewError(message string) error {
    return &Error{Message: message}
}

// 返回错误信息
func (err *Error) Error() string {
    return err.Message
}
