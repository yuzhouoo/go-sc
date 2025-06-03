package request

type AccountList struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Name     string `json:"name"`
}

type AccountRegister struct {
	Name    string `json:"name"`
	Account string `json:"account"`
}

type AccountLogin struct {
	Account string `json:"account"`
}

type AccountEditInfo struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Account string `json:"account"`
}

type AccountClose struct {
	UUID    string `json:"uuid"`
	Account string `json:"account"`
}
