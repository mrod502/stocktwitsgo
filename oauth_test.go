package stocktwits

import "testing"

func TestOauth(t *testing.T) {
	requestCode(consumerKey, defaultRedirectURI, scope)

}
