package pubkey

import "os"

var (
	GoogleOAuth2N = os.Getenv("GOOGLE_OAUTH2_N")
	GoogleOAuth2E = os.Getenv("GOOGLE_OAUTH2_E")
)

type Oauth2Cert struct {
	N string
	E string
}

func NewOauth2Cert() *Oauth2Cert {
	return &Oauth2Cert{
		N: GoogleOAuth2N,
		E: GoogleOAuth2E,
	}
}
