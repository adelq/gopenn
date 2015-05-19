package penn

import "os"

var client *Client

func setup() {
	client = NewClient(os.Getenv("DIRECTORY_API_USERNAME"), os.Getenv("DIRECTORY_API_PASSWORD"))
}

func teardown() {
}
