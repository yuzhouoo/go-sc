package response

type AccountLoginData struct {
	ID      int64  `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	UUID    string `json:"uuid"`
}

type AccountLogin struct {
	ResponseCommon
	Data AccountLoginData `json:"data"`
}

type AccountListItemData struct {
	ID      int64  `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
}

type AccountList struct {
	ResponseCommon
	Data ResponseListDataCommon `json:"data"`
}

type AccountInfoItemData struct {
	ID      int64  `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
}

type AccountInfoData struct {
	ResponseCommon
	Data AccountInfoItemData `json:"data"`
}

type AccountClose struct {
	ResponseCommon
}
