package main

import (
	"context"
	"fmt"
	"net/http"
	"crypto/tls"
	"os"

	"github.com/zemirco/keycloak"
	"golang.org/x/oauth2"
)

func CheckUser(serverip string, user string, password string, user1 string, password1 string) {
	// create a new config for the "admin-cli" client
	config := oauth2.Config{
		ClientID: "admin-cli",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://"+serverip+":8443/realms/master/protocol/openid-connect/token",
		},
		Scopes: []string{"openid"},
	}

	ctx := context.Background()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// get a token
	//token, err := config.PasswordCredentialsToken(ctx, "sentrywire", "1qaz@9759")
	token, err := config.PasswordCredentialsToken(ctx, user,password)
	if err != nil {
		fmt.Println(err);
		os.Exit(2);
	}

	// use the token on every http request
	httpClient := config.Client(ctx, token)

	// create a new keycloak client instance
	_, err2 := keycloak.NewKeycloak(httpClient, "https://"+serverip+":8443/")
	if err2 != nil {
		fmt.Println(err2);
		os.Exit(3);
	}
	fmt.Println(token);

	// pretend to be user_a
	userAConfig := oauth2.Config{
		ClientID:     "admin-cli",
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("https://"+serverip+":8443/realms/%s/protocol/openid-connect/token", "master"),//realm),
		},
	}

	userAToken, err := userAConfig.PasswordCredentialsToken(ctx, user1, password1)
	if err != nil {
		fmt.Println(err);
		os.Exit(4);
	}
	fmt.Println("***************")
	fmt.Println(userAToken);
	fmt.Println("***************")

	os.Exit(0);

}

func main() {
	if len(os.Args) < 6 {
		os.Exit(1);
	}
	CheckUser(os.Args[1],os.Args[2],os.Args[3],os.Args[4],os.Args[5])
}
