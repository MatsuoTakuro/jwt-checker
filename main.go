package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/MatsuoTakuro/jwt-checker/idtoken"
	"github.com/MatsuoTakuro/jwt-checker/pubkey"
)

func main() {
	rawData := strings.Split(idtoken.IDToken, ".")
	rawh, rawp, rawsig := rawData[0], rawData[1], rawData[2]

	h, err := base64.RawURLEncoding.DecodeString(rawh)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	err = os.WriteFile("/app/idtoken/header.json", h, 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	p, err := base64.RawURLEncoding.DecodeString(rawp)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	err = os.WriteFile("/app/idtoken/payload.json", p, 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	rawpk := pubkey.NewOauth2Cert()
	n, _ := base64.RawURLEncoding.DecodeString(rawpk.N)
	e, _ := base64.RawURLEncoding.DecodeString(rawpk.E)
	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(n),
		E: int(new(big.Int).SetBytes(e).Int64()),
	}

	msg := sha256.Sum256([]byte(rawh + "." + rawp))

	sig, err := base64.RawURLEncoding.DecodeString(rawsig)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, msg[:], sig); err != nil {
		fmt.Println("invalid id_token")
	} else {
		fmt.Println("valid id_token")
		fmt.Println("header: ", string(h))
		fmt.Println("payload: ", string(p))
	}
}
