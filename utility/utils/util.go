package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// CodeErr 返回带错误码的错误码
func CodeErr(code int, message string, detail ...interface{}) error {
	var d interface{}
	if len(detail) > 0 {
		d = detail[0]
	}
	gCode := gcode.New(code, message, d)
	return gerror.NewCode(gCode, message)
}

func MD5(str string) string {
	b := []byte(str)
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
func EncodePassword(password, salt string) string {
	saltSign := fmt.Sprintf("%s%s", MD5(password), salt)
	return MD5(saltSign)
}
