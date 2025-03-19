package internal

import (
	"bytes"
	"io"
	"net/http"
)

// SendRequest 发送HTTP请求
func SendRequest(url string, headers map[string]string, body []byte, method string, isHeader bool) (interface{}, error) {
	// 创建一个新的请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if isHeader {
		return resp.Header, nil
	} else {
		// 读取响应体
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return responseBody, nil
	}

}
