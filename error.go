// Copyright 2016 zxfonline@sina.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gerror

import (
	"fmt"
)

const (
	//操作成功的返回码常量
	OK int32 = 0

	//通用错误
	CUSTOM_ERROR int32 = 100000
	//服务器端一般错误
	SERVER_INTERNAL_ERROR int32 = 100001
	//读取客户端发送的数据时产生该异常
	SERVER_CDATA_ERROR int32 = 100002
	//处理客户端发送的消息时产生该异常
	SERVER_CMSG_ERROR int32 = 100003
	//服务器端的文件没有找到错误
	SERVER_FILE_NOT_FOUND int32 = 100004
	//服务器端的访问被拒绝
	SERVER_ACCESS_REFUSED int32 = 100005

	//客户端访问超时
	CLIENT_TIMEOUT int32 = 200000
	//客户端IO错误(如无法建立连接，数据发送失败等)
	CLIENT_IO_ERROR int32 = 200001
)

//判断错误类型是否是允许的错误范围 true=消息内部定义的错误，false=系统级别错误，需要关闭连接
func IsCustomError(e int32) bool {
	return e < CUSTOM_ERROR
}

var ErrMap = map[int32]string{
	OK: "成功",

	CUSTOM_ERROR:          "通用错误",
	SERVER_INTERNAL_ERROR: "服务器端一般错误",
	SERVER_CDATA_ERROR:    "读取客户端发送的数据异常",
	SERVER_CMSG_ERROR:     "处理客户端发送的数据异常",
	SERVER_FILE_NOT_FOUND: "服务器端的文件没有找到",
	SERVER_ACCESS_REFUSED: "服务器端的访问被拒绝",

	CLIENT_TIMEOUT:  "客户端访问超时",
	CLIENT_IO_ERROR: "客户端IO错误",
}

//系统常规错误
type SysError struct {
	Code    int32  `json:"ret"`
	Content string `json:"result,omitempty"`
	Cause   error  `json:"-"`
}

func (this *SysError) Error() string {
	if this.Cause == nil {
		return fmt.Sprintf("ret:%d,msg:%s", this.Code, this.Content)
	} else if this.Content != "" {
		return fmt.Sprintf("ret:%d,msg:%s,cause:%s", this.Code, this.Content, this.Cause.Error())
	} else {
		return fmt.Sprintf("ret:%d,cause:%s", this.Code, this.Cause.Error())
	}
}

func New(code int32, err error) *SysError {
	return &SysError{Code: code, Content: err.Error()}
}
func NewError(code int32, err string) *SysError {
	return &SysError{Code: code, Content: err}
}
