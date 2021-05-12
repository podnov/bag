package bscscan

type Int64ApiResult struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Result int64 `json:"result,string"`
}
