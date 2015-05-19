package penn

var client *Client

func setup() {
	client = NewClient("UPENN_OD_emxM_1000904", "t4ii5rdud602n63ln2h1ld29hr")
}

func teardown() {
}
