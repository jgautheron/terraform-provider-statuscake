package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/jgautheron/terraform-provider-statuscake/v2/provider"
)

var testProviders = map[string]*schema.Provider{
	"statuscake": provider.Provider(),
}

func TestProvider(t *testing.T) {
	if err := testProviders["statuscake"].InternalValidate(); err != nil {
		t.Errorf("failed to validate provider: %+v", err)
	}
}
