package utils

import (
	"bytes"
	"io"
	"net/http"
)

func RespToBuf(resp *http.Response, buf *bytes.Buffer) {
	if resp.ContentLength == 0 {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Errorln("返回体关闭失败:", err)
		}
	}(resp.Body)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		Logger.Errorln("返回体复制失败:", err)
		return
	}
}
