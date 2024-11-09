package sql_endpoint

import (
	"encoding/base64"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"

	"github.com/glalanne/provider-databricks/config/common"
)

func sqlEndpointConnectionDetails(attr map[string]interface{}) (map[string][]byte, error) {

	jdbcUrl, err := common.GetField(attr, "jdbcUrl")
	if err != nil {
		return nil, err
	}

	jdbcUrlBytes, err := base64.StdEncoding.DecodeString(jdbcUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot serialize jdbcUrl key data")
	}

	hostname, err := common.GetField(attr, "odbcParams[0].hostname")
	if err != nil {
		return nil, err
	}

	hostnameBytes, err := base64.StdEncoding.DecodeString(hostname)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot serialize hostname key data")
	}

	path, err := common.GetField(attr, "odbcParams[0].path")
	if err != nil {
		return nil, err
	}

	pathBytes, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot serialize path key data")
	}

	// conn["odbcParams.hostname"] = []byte(a[0]["hostname"])
	// conn["odbcParams.path"] = []byte(a[0]["path"])
	// conn["odbcParams.port"] = []byte(a[0]["port"])
	// conn["odbcParams.protocol"] = []byte(a[0]["protocol"])

	conn := map[string][]byte{
		// xpv1.ResourceCredentialsSecretEndpointKey: []byte(jdbcUrl),
		"hostname": hostnameBytes,
		"path":     pathBytes,
		"jdbcUrl":  jdbcUrlBytes,
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
