package response

type ResponseCommon struct {
	MsgCode int    `json:"msgCode"`
	Desc    string `json:"desc"`
}

type ResponseListDataCommon struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	List     any `json:"list"`
}
