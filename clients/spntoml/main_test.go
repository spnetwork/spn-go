package spntoml

import "log"

// ExampleGetTOML gets the spn.toml file for coins.asia
func ExampleClient_GetSpnToml() {
	_, err := DefaultClient.GetSpnToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
