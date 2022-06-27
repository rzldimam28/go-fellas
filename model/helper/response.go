package helper

import "github.com/rzldimam28/wlb-test/model/web/response"

func CreateWebResponse(code int, status string, data interface{}) response.WebResponse {
	return response.WebResponse{
		Code: code,
		Status: status,
		Data: data,
	}
}