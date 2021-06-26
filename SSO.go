
package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	conf *oauth2.Config
	ctx  context.Context
)

func main() {

	ctx = context.Background()
	conf = &oauth2.Config{
		ClientID:     "demo-client",
		ClientSecret: "274cd3ed-dfdc-4b9f-bfb7-9593cc78cd2d",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://10.91.170.177:8443/auth/realms/demo",
			TokenURL: "https://10.91.170.177:8443/auth/realms/demo/protocol/openid-connect",
		},
		// my own callback URL
		RedirectURL: "http://127.0.0.1:9999/oauth/callback",
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	log.Println(color.CyanString("You will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)
	open.Run(url)
	time.Sleep(1 * time.Second)
	log.Printf("Authentication URL: %s\n", url)

	http.HandleFunc("/oauth/callback", callbackHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
