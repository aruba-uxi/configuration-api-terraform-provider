# Terraform Provider HPEUXI - [Terraform Provider HPEUXI](#terraform-provider-hpeuxi)
- [Terraform Provider HPEUXI - Terraform Provider HPEUXI](#terraform-provider-hpeuxi---terraform-provider-hpeuxi)
  - [Introduction](#introduction)
  - [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Usage](#usage)
    - [Ensure correct terraform version](#ensure-correct-terraform-version)
    - [Create manifest](#create-manifest)
    - [Export API client credentials as environment variables](#export-api-client-credentials-as-environment-variables)
    - [Initialize Terraform working directory:](#initialize-terraform-working-directory)
    - [Plan resources](#plan-resources)
    - [Deploy resources](#deploy-resources)
  - [Development](#development)
    - [.terraformrc for local development](#terraformrc-for-local-development)
    - [Justfile](#justfile)
    - [Installing the provider locally](#installing-the-provider-locally)
  - [Running Local Test Suite](#running-local-test-suite)
  - [Updating the client](#updating-the-client)
  - [Running Acceptance Tests](#running-acceptance-tests)
  - [Building and Distribution](#building-and-distribution)


## Introduction

A repository to hold the HPEUXI terraform provider.

## Installation

See the [INSTALL.md](INSTALL.md) file for instructions on how to install the binaries manually.

## Prerequisites

1. Terraform version >= v1.7.0 and 64-bit [install terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)
1. A Greenlake API client to authenticate against UXI configuration API.
1. Terraform basics. [Terraform Introduction](https://www.terraform.io/intro/index.html)

## Usage

### Ensure correct terraform version

Make sure you have the correct Terraform version installed; It must be v1.7.0 or later:

```shell
terraform version
```

### Create manifest

Create a Terraform file in your working directory, and reference the desired provider version from the terraform [registry](https://registry.terraform.io/providers/aruba-uxi/hpeuxi/latest). For more information please consult the [documentation](docs/):

```terraform
terraform {
    required_providers {
        hpeuxi = {
            source  = "aruba-uxi/hpeuxi"
            version = ">= 0.0.0-pre.alpha.3"
        }
    }
    required_version = ">= 1.7.0"
}

# create a group
resource "hpeuxi_group" "my_group" {
    name = "level_1"
}

# import a sensor
resource "hpeuxi_sensor" "my_sensor" {
    name = "my_sensor_name"

    lifecycle {
        prevent_destroy = true
    }
}

import {
    to = hpeuxi_sensor.my_sensor
    id = "my_sensor_id"
}

# assign the sensor to a group
resource "hpeuxi_sensor_group_assignment" "my_uxi_sensor_group_assignment" {
    sensor_id = hpeuxi_sensor.my_sensor.id
    group_id = hpeuxi_group.my_group.id
}
```

### Export API client credentials as environment variables

```bash
export GREENLAKE_UXI_CLIENT_ID=<greenlake_uxi_client_id>
export GREENLAKE_UXI_CLIENT_SECRET=<greenlake_uxi_client_secret>
```

### Initialize Terraform working directory:
```shell
> terraform init
```

### Plan resources
```shell
> terraform plan
```

### Deploy resources
```shell
> terraform apply
```

## Development

### .terraformrc for local development

Add the appropriate terraform provider dev override to your `~/.terraformrc` file to ensure that that terraform operations is performed against the local provider.

Example `~/.terraformrc` file
```
provider_installation {
  dev_overrides {
      "registry.terraform.io/arubauxi/hpeuxi" = "/Users/<user>/go/bin"
  }
  direct {}
}
```

### Justfile

Install [just](https://github.com/casey/just?tab=readme-ov-file#installation). This is used to run
the recipes defined in the `Justfile`. Run the following to see a helpful list of commands:

```shell
just help
```

### Installing the provider locally

```shell
just setup-run
```

## Running Local Test Suite

```shell
just test
```

## Updating the client

The client used for interfacing with the UXI backend is autogenerated using an API Spec and [openapitools](https://github.com/OpenAPITools/openapi-generator-cli). In order to generate a client, ensure that the desired OpenAPI spec is copied to `./pkg/config-api-client/api/openapi.yaml`. Thereafter run:

```shell
just generate-config-api-client
```

## Running Acceptance Tests

Set the `GREENLAKE_UXI_CLIENT_ID` and `GREENLAKE_UXI_CLIENT_SECRET` environment variables for the acceptance test customer. The client credentials must have sufficient permissions execute create, read, update and destroy operations for the UXI customer. Also set the `UXI_HOST` environment variable to target the staging
environment.

```bash
export GREENLAKE_UXI_CLIENT_ID=<client_id>
export GREENLAKE_UXI_CLIENT_SECRET=<client_secret>
export UXI_HOST=api.staging.capedev.io
```

```shell
just acceptance-tests
```

Note: it is possible to run acceptance tests against any UXI customer, if the configurations in
`test/live/config/config.go` are updated to values for the given customer.

## Building and Distribution

All builds must be signed by HPE before distribution.
Read the [Notices Report](public/README.md) and ensure a report is added to every build which is published
All open source sourcecode must be submitted to the HPE DSM to make this available for customers to download on request.
