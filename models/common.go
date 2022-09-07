package models

var ResultCode = map[int64]string{
    0: "操作成功",
    1: "操作失败",
    2: "参数检验失败",
    3: "未登录或 token 已经过期",
    4: "没有相关权限",
}

type CommonResult struct {
    Code int64 `json:"code"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

func SuccessResult(data interface{}) *CommonResult {
    return &CommonResult{
        Code: 0,
        Message: ResultCode[0],
        Data: data,
    }
}

func SuccessResultWithMessage(data interface{}, message string) *CommonResult {
    return &CommonResult{
        Code: 0,
        Message: message,
        Data: data,
    }
}

func FailedResult() *CommonResult {
    return &CommonResult{
        Code: 1,
        Message: ResultCode[1],
    }
}

func FailedResultWithMessage(message string) *CommonResult {
    return &CommonResult{
        Code: 1,
        Message: message,
    }
}

func ValidateFailedResult() *CommonResult {
    return &CommonResult{
        Code: 2,
        Message: ResultCode[2],
    }
}

func ValidateResultWithMessage(message string) *CommonResult {
    return &CommonResult{
        Code: 2,
        Message: message,
    }
}

func UnauthorizedResult() *CommonResult {
    return &CommonResult{
        Code: 3,
        Message: ResultCode[3],
    }
}

func UnauthorizedResultWithMessage(message string) *CommonResult {
    return &CommonResult{
        Code: 3,
        Message: message,
    }
}

func ForbiddenResult() *CommonResult {
    return &CommonResult{
        Code: 4,
        Message: ResultCode[4],
    }
}

func ForbiddenResultWithMessage(message string) *CommonResult {
    return &CommonResult{
        Code: 4,
        Message: message,
    }
}
