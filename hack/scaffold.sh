#!/bin/bash

# Directory containing the files
resources=$(jq --raw-output '.provider_schemas."registry.terraform.io/databricks/databricks".resource_schemas | keys | .[]' ./config/schema.json)
root_folder="./config"

for resource in $resources
do
  resource_name=$(echo $resource | sed 's/databricks_//')

  if [ ! -d "$root_folder/$resource_name" ]; then
    echo "Creating: $resource_name"
    mkdir "$root_folder/$resource_name"
    cat > "$root_folder/$resource_name/config.go"  << EOF
package $resource_name

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("$resource", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
EOF
    echo "\"$resource\":               config.IdentifierFromProvider," >> ./.work/external_name.txt
    echo "$resource_name.Configure," >> ./.work/provider.txt

    echo "Open ./.work/external_name.txt and add it into ./config/external_name.go"
    echo "Open ./.work/provider.txt and add it into ./config/provider.go"
  else 
    echo "Config exists for: $resource_name"
  fi
done
