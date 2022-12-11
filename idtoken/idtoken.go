package idtoken

import "os"

var (
	IDToken = os.Getenv("ID_TOKEN")
)
