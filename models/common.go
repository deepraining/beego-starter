package models

var ResultCode = map[int32]string{
    200: "操作成功",
    500: "操作失败",
    400: "参数检验失败",
    401: "未登录或 token 已经过期",
    403: "没有相关权限",
}

type CommonResult struct {
    Code int32 `json:"code"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

func SuccessResult(data interface{}) *CommonResult {
    return &CommonResult{
        Code: 200,
        Message: ResultCode[200],
        Data: data,
    }
}

func SuccessResultWithMessage(data interface{}, message string) *CommonResult {
    return &CommonResult{
        Code: 200,
        Message: message,
        Data: data,
    }
}

func FailedResult() *CommonResult {
    return &CommonResult{
        Code: 500,
        Message: ResultCode[500],
    }
}

func FailedResultWithMessage(message string) *CommonResult {
    return &CommonResult{
        Code: 500,
        Message: message,
    }
}

func ValidateFailedResult() *CommonResult {
    return &CommonResult{
        Code: 400,
        Message: ResultCode[400],
    }
}

func ValidateResultWithMessage(message string) *CommonResult {
    return &CommonResult{
        Code: 400,
        Message: message,
    }
}

func UnauthorizedResult(data interface{}) *CommonResult {
    return &CommonResult{
        Code: 401,
        Message: ResultCode[401],
        Data: data,
    }
}

func ForbiddenResult(data interface{}) *CommonResult {
    return &CommonResult{
        Code: 403,
        Message: ResultCode[403],
        Data: data,
    }
}
