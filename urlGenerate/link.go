package urlGenerate

import (
	"fmt"
	"net/url"
	"random/randomNumber"
)

func Link() {
	base, err := url.Parse("http://www.example.com")
	if err != nil {
		return
	}

	// Path params
	base.Path += "referral"

	// Query params
	params := url.Values{}
	params.Add("q", randomNumber.RandomString(6))
	base.RawQuery = params.Encode()

	fmt.Printf("%q\n", base.String())
}
