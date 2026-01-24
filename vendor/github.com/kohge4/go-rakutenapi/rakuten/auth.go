package rakuten

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// 参考 https://webservice.rakuten.co.jp/document/oauth

const (
	AuthURL  = "https://app.rakuten.co.jp/services/authorize"
	TokenURL = "https://app.rakuten.co.jp/services/token"
)

type Authenticator struct {
	config  *oauth2.Config
	context context.Context
}

func NewAuthenticator(clientID string, clientSecret string, redirectURL string, scopes ...string) Authenticator {
	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}
	fmt.Printf(" \n TTTTTTT %v\n", scopes)
	tr := &http.Transport{
		TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
	}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: tr})

	return Authenticator{
		config:  cfg,
		context: ctx,
	}
}

func (a Authenticator) AuthURL(state string) string {
	return a.config.AuthCodeURL(state)
}

// Token pulls an authorization code from an HTTP request and attempts to exchange
// it for an access token.  The standard use case is to call Token from the handler
// that handles requests to your application's redirect URL.
func (a Authenticator) Token(state string, r *http.Request) (*oauth2.Token, error) {
	values := r.URL.Query()
	if e := values.Get("error"); e != "" {
		return nil, errors.New("rakuten: auth failed - " + e)
	}
	code := values.Get("code")
	if code == "" {
		return nil, errors.New("rakuten: didn't get access code")
	}
	actualState := values.Get("state")
	if actualState != state {
		return nil, errors.New("rakuten: redirect state parameter doesn't match")
	}
	return a.config.Exchange(a.context, code)
}

// NewClient creates a Client that will use the specified access token for its API requests.
func (a Authenticator) NewClient(token *oauth2.Token) Client {
	client := a.config.Client(a.context, token)
	//  client の違いについて rakuten.go の client を呼び出している
	return Client{
		client: client,
		//baseURL: baseAddress,
	}
}

// Exchange is like Token, except it allows you to manually specify the access
// code instead of pulling it out of an HTTP request.
func (a Authenticator) Exchange(code string) (*oauth2.Token, error) {
	return a.config.Exchange(a.context, code)
}

// Token gets the client's current token.
func (c *Client) Token() (*oauth2.Token, error) {
	// ここのclient は oauth2.0 のclient ==> つまり http.client
	transport, ok := c.client.Transport.(*oauth2.Transport)
	if !ok {
		return nil, errors.New("rakuten web service: oauth2 transport type not correct")
	}
	// TokenSource 型 は Token() をもつ interface{} (Tokenを返す何かのことをさす)
	t, err := transport.Source.Token()
	if err != nil {
		return nil, err
	}
	return t, nil
}
