package slack

import "golang.org/x/oauth2"

func init() {
	// We cannot set the client_secret in the auth header, which is the
	// default behavior of package oauth2.
	oauth2.RegisterBrokenAuthHeaderProvider("https://slack.com/api")
}

var Endpoint = oauth2.Endpoint{
	TokenURL: "https://slack.com/api/oauth.access",
}
