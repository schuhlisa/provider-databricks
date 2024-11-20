
# Provider Databricks

<div style="text-align: center;">

![CI](https://github.com/glalanne/provider-databricks/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/glalanne/provider-databricks/all.svg)](https://github.com/glalanne/provider-databricks/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/glalanne/provider-databricks)](https://goreportcard.com/report/github.com/glalanne/provider-databricks)
[![Contributors](https://img.shields.io/github/contributors/glalanne/provider-databricks)](https://github.com/glalanne/provider-databricks/graphs/contributors)


</div>

`provider-databricks` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
[Databricks](https://registry.terraform.io/providers/databricks/databricks/latest/docs).

Warning:
`This provider is not ready to be used for production and might be missing resources`

Most of the testing have been done on [Azure Databricks](https://azure.microsoft.com/en-ca/products/databricks)

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/glalanne/provider-databricks):
```
up ctp provider install glalanne/provider-databricks:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-databricks
spec:
  package: glalanne/provider-databricks:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/glalanne/provider-databricks).

## Exemples

A few exemples are provided [here](./examples/) to show case how to configure the provider, and how to use some resources

## Resources
Here is the list of all supported resources from the [Terraform provider](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs):

- [databricks_alert](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/alert)
- [databricks_budget](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/budget)
- [databricks_catalog](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/catalog)
- [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/cluster)
- [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/cluster_policy)
- [databricks_connection](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/connection)
- [databricks_external_location](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/external_location)
- [databricks_git_credential](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/git_credential)
- [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/grants)
- [databricks_group](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/group)
- [databricks_group_member](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/group_member)
- [databricks_group_role](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/group_role)
- [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/instance_pool)
- [databricks_ip_access_list](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/ip_access_list)
- [databricks_job](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/job)
- [databricks_notebook](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/notebook)
- [databricks_permission_assignment](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/permission_assignment)
- [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/permissions)
- [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/pipeline)
- [databricks_query](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/query)
- [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/schema)
- [databricks_secret](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/secret)
- [databricks_secret_scope](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/secret_scope)
- [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/service_principal)
- [databricks_service_principal_role](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/service_principal_role)
- [databricks_sql_alert](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_alert)
- [databricks_sql_dashboard](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_dashboard)
- [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_endpoint)
- [databricks_sql_global_config](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_global_config)
- [databricks_sql_permissions](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_permissions)
- [databricks_sql_query](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_query)
- [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/sql_table)
- [databricks_token](https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/token)

## Developing

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/glalanne/provider-databricks/issues).
