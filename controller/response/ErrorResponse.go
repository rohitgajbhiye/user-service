package response

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	ErrorCode    int    `json:"errCode"`
	ErrorMessage string `json:"errMessage"`
	ErrorType    string `json:"errType"`
}
