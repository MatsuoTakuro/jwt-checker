package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/MatsuoTakuro/jwt-checker/idtoken"
)

func main() {
	rawData := strings.Split(idtoken.IDToken, ".")
	rawh, rawp, _ := rawData[0], rawData[1], rawData[2]
	fmt.Println("raw header: ", rawh)
	fmt.Println("raw payload: ", rawp)

	h, err := base64.RawURLEncoding.DecodeString(rawh)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	p, err := base64.RawURLEncoding.DecodeString(rawp)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("header: ", string(h))
	fmt.Println("payload: ", string(p))
}
