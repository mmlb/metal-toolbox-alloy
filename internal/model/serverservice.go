package model

import (
	"net/url"
	"os"

	"github.com/pkg/errors"
)

const (
	// EnvServerserviceSkipOAuth when set to true will skip server service OAuth
	EnvVarServerserviceSkipOAuth = "SERVERSERVICE_SKIP_OAUTH"

	// server vendor, model attributes are stored in this namespace
	ServerVendorAttributeNS = "sh.hollow.alloy.server_vendor_attributes"

	// additional server metadata are stored in this namespace
	ServerMetadataAttributeNS = "sh.hollow.alloy.server_metadata_attributes"

	// server service server serial attribute key
	ServerSerialAttributeKey = "serial"

	// server service server model attribute key
	ServerModelAttributeKey = "model"

	// server service server vendor attribute key
	ServerVendorAttributeKey = "vendor"
)

// LoadServerServiceEnvVars sets any env SERVERSERVICE_* configuration parameters
func (c *Config) LoadServerServiceEnvVars() {
	if facility := os.Getenv("SERVERSERVICE_FACILITY_CODE"); facility != "" {
		c.ServerService.FacilityCode = facility
	}

	// env var serverService endpoint
	if endpoint := os.Getenv("SERVERSERVICE_ENDPOINT"); endpoint != "" {
		c.ServerService.Endpoint = endpoint
	}

	// OIDC provider endpoint
	if oidcProviderEndpoint := os.Getenv("SERVERSERVICE_OIDC_PROVIDER_ENDPOINT"); oidcProviderEndpoint != "" {
		c.ServerService.OidcProviderEndpoint = oidcProviderEndpoint
	}

	// Audience endpoint
	if audienceEndpoint := os.Getenv("SERVERSERVICE_AUDIENCE_ENDPOINT"); audienceEndpoint != "" {
		c.ServerService.AudienceEndpoint = audienceEndpoint
	}

	// env var OAuth client secret
	if clientSecret := os.Getenv("SERVERSERVICE_CLIENT_SECRET"); clientSecret != "" {
		c.ServerService.ClientSecret = clientSecret
	}

	// env var OAuth client ID
	if clientID := os.Getenv("SERVERSERVICE_CLIENT_ID"); clientID != "" {
		c.ServerService.ClientID = clientID
	}
}

// ValidateServerServiceParams checks required serverservice configuration parameters are present
// and returns the serverservice URL endpoint
func (c *Config) ValidateServerServiceParams() (*url.URL, error) {
	if c.ServerService.FacilityCode == "" {
		return nil, errors.Wrap(ErrConfig, "serverService facility code not defined")
	}

	if c.ServerService.Endpoint == "" {
		return nil, errors.Wrap(ErrConfig, "serverService endpoint not defined")
	}

	endpoint, err := url.Parse(c.ServerService.Endpoint)
	if err != nil {
		return nil, errors.Wrap(ErrConfig, "error in serverService endpoint URL: "+err.Error())
	}

	if os.Getenv(EnvVarServerserviceSkipOAuth) == "true" {
		return endpoint, nil
	}

	if c.ServerService.OidcProviderEndpoint == "" {
		return nil, errors.Wrap(ErrConfig, "serverService OIDC provider endpoint not defined")
	}

	if c.ServerService.AudienceEndpoint == "" {
		return nil, errors.Wrap(ErrConfig, "serverService Audience endpoint not defined")
	}

	if c.ServerService.ClientSecret == "" {
		return nil, errors.Wrap(ErrConfig, "serverService client secret not defined")
	}

	if c.ServerService.ClientID == "" {
		return nil, errors.Wrap(ErrConfig, "serverService client ID not defined")
	}

	return endpoint, nil
}