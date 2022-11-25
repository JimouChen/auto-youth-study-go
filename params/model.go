package params

type CidStruct struct {
	Data struct {
		Entity struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"entity"`
	} `json:"data"`
}

type TokenStruct struct {
	Data struct {
		Entity struct {
			Token string `json:"token"`
		} `json:"entity"`
	} `json:"data"`
}

type SaveResult struct {
	Errmsg string `json:"errmsg"`
	Errno  int    `json:"errno"`
	Msg    string `json:"msg"`
}
