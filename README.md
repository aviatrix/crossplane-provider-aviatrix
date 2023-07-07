# Provider Aviatrix

This repository is a work in progress.  In the future, it will be the site 
of a fully supported crossplane provider for configuring Aviatrix Controllers.
It is not currently ready for production use, and should be considered to be
a pre-alpha level quality.

`provider-aviatrix` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Aviatrix API.

## Getting Started

This section is under construction, and not yet functional.

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/aviatrix/provider-aviatrix):
```
up ctp provider install aviatrix/provider-aviatrix:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aviatrix
spec:
  package: aviatrix/provider-aviatrix:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/aviatrix/provider-aviatrix).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
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

Run Unit/Integ Tests:
```console
make test
```
(see internal/README.md for explanation of tests)

Run tests against kind cluster and live controller:
```console
UPTEST_EXAMPLE_LIST=$(find examples/*.yaml -type f| tr '\n' ',') make e2e
```
This will run the e2e suite against all examples.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/aviatrix/provider-aviatrix/issues).
