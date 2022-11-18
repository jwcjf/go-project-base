package response

import (
	pbErrors "github.com/jwcjf/go-project-base/errors"
)

// Response ...
type Response struct {
	pbErrors.Error
	// 数据集
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

// Page ...
type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type page struct {
	Page
	List interface{} `json:"list"`
}
