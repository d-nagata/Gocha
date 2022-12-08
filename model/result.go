package model

type (
	ResultRequest struct {
		Size int64 `json:"size"`
	}
	ResultResponse struct {
		Result []string `json:"result"`
	}
)
