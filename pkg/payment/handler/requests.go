package handler

type PaginationRequest struct {
	StartId  uint `json:"startId" form:"startId"`
	PageSize uint `json:"pageSize" form:"pageSize"`
}

type PaginationResponse[T any] struct {
	StartId uint `json:"startId"`
	EndId   uint `json:"endId"`
	Count   uint `json:"count"`
	Data    []T  `json:"data"`
}
