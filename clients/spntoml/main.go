package spntoml

import "net/http"

// SpnTomlMaxSize is the maximum size of spn.toml file
const SpnTomlMaxSize = 5 * 1024

// WellKnownPath represents the url path at which the spn.toml file should
// exist to conform to the federation protocol.
const WellKnownPath = "/.well-known/spn.toml"

// DefaultClient is a default client using the default parameters
var DefaultClient = &Client{HTTP: http.DefaultClient}

// Client represents a client that is capable of resolving a Spn.toml file
// using the internet.
type Client struct {
	// HTTP is the http client used when resolving a Spn.toml file
	HTTP HTTP

	// UseHTTP forces the client to resolve against servers using plain HTTP.
	// Useful for debugging.
	UseHTTP bool
}

type ClientInterface interface {
	GetSpnToml(domain string) (*Response, error)
	GetSpnTomlByAddress(addy string) (*Response, error)
}

// HTTP represents the http client that a stellertoml resolver uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// Response represents the results of successfully resolving a spn.toml file
type Response struct {
	AuthServer       string `toml:"AUTH_SERVER"`
	FederationServer string `toml:"FEDERATION_SERVER"`
	EncryptionKey    string `toml:"ENCRYPTION_KEY"`
	SigningKey       string `toml:"SIGNING_KEY"`
}

// GetSpnToml returns spn.toml file for a given domain
func GetSpnToml(domain string) (*Response, error) {
	return DefaultClient.GetSpnToml(domain)
}

// GetSpnTomlByAddress returns spn.toml file of a domain fetched from a
// given address
func GetSpnTomlByAddress(addy string) (*Response, error) {
	return DefaultClient.GetSpnTomlByAddress(addy)
}
