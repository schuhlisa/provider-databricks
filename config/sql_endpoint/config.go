package sql_endpoint

import (
	"encoding/base64"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"

	"github.com/glalanne/provider-databricks/config/common"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

func sqlEndpointConnectionDetails(attr map[string]interface{}) (map[string][]byte, error) {
	jdbcUrl, ok := attr["jdbcUrl"].(string)
	if !ok {
		return nil, errors.New("cannot get jdbcUrl attribute")
	}

	hostname, err := common.GetField(attr, "odbcParams[0].hostname")
	if err != nil {
		return nil, err
	}

	hostnameBytes, err := base64.StdEncoding.DecodeString(hostname)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot serialize hostname key data")
	}

	// conn["odbcParams.hostname"] = []byte(a[0]["hostname"])
	// conn["odbcParams.path"] = []byte(a[0]["path"])
	// conn["odbcParams.port"] = []byte(a[0]["port"])
	// conn["odbcParams.protocol"] = []byte(a[0]["protocol"])

	conn := map[string][]byte{
		xpv1.ResourceCredentialsSecretEndpointKey: []byte(jdbcUrl),
		"hostname": hostnameBytes,
	}

	return conn, nil
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_sql_endpoint", func(r *config.Resource) {
		r.ShortGroup = "databricks"
		r.Sensitive.AdditionalConnectionDetailsFn = sqlEndpointConnectionDetails

	})
}
