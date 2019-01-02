# go-arty

[![GoDoc](https://godoc.org/github.com/target/go-arty?status.svg)](https://godoc.org/github.com/target/go-arty)
[![Go Report Card](https://goreportcard.com/badge/target/go-arty)](https://goreportcard.com/report/target/go-arty)
[![Coverage Status](https://coveralls.io/repos/target/go-arty/badge.svg?branch=master)](https://coveralls.io/r/target/go-arty?branch=master)
[![Build Status](https://travis-ci.org/target/go-arty.svg?branch=master)](https://travis-ci.org/target/go-arty)

go-arty is a Go client library for accessing the [Artifactory](https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API) and [Xray](https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API) API.

## Artifactory

### Usage

```go
import "github.com/target/go-arty/artifactory"
```

Construct a new Artifactory client, then use the various services on the client to access different parts of the Artifactory API. For example:

```go
client, _ := artifactory.NewClient("artifactory.company.com", nil)

// list all users from the artifactory server
users, _, err := client.Users.GetAllSecurity()
```

### Authentication

The `artifactory` package allows you to pass basic auth or an [API Key](https://www.jfrog.com/confluence/display/RTF/Updating+Your+Profile#UpdatingYourProfile-APIKey).

Example using basic auth:

```go
client, _ := artifactory.NewClient("artifactory.company.com", nil)

client.Authentication.SetBasicAuth("username", "password")
```

Example using API Key:

```go
client, _ := artifactory.NewClient("artifactory.company.com", nil)

client.Authentication.SetTokenAuth("token")
```

## Xray

### Usage

```go
import "github.com/target/go-arty/xray"
```

Construct a new Xray client, then use the various services on the client to access different parts of the Xray API. For example:

```go
client, _ := xray.NewClient("artifactory.company.com", nil)

// list all users from the xray server
users, _, err := client.Users.GetAll()
```

### Authentication

The `xray` package allows you to pass basic auth or a [token](https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-Authentication).

**NOTE: To get the token for Xray, you have to hit an API endpoint that returns the token. See [the docs for more info](https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API#XrayRESTAPI-GetToken).**

Example using basic auth:

```go
client, _ := xray.NewClient("xray.company.com", nil)

client.Authentication.SetBasicAuth("username", "password")
```

Example using token:

```go
client, _ := xray.NewClient("xray.company.com", nil)

client.Authentication.SetTokenAuth("token")
```

## Creating/Updating Resources

All structs in this library use pointer values for all non-repeated fields. This allows distinguishing between unset fields and those set to a zero-value. Helper functions have been provided to easily create these pointers for string, bool, and int values. For example:

```go
// create a new user named "admin"
user := &artifactory.SecurityUser{
	Name:     artifactory.String("admin"),
	Email:    artifactory.String("admin@company.com"),
	Password: artifactory.String("secretPassword"),
	Admin:    artifactory.Bool(true),
}

client.Users.CreateSecurity(user)
```

Users who have worked with protocol buffers should find this pattern familiar.

## Versioning

In general, `go-arty` follows [semantic versioning](https://semver.org/) as closely as we can for tagging releases of the package. For self-contained libraries, the application of semantic versioning is relatively straightforward and generally understood. But because `go-arty` is a client library for the Artifactory API and the Xray API, which both change behavior frequently, we've adopted the following versioning policy:

* We increment the major version with any incompatible change to either package (`artifactory` or `xray`) in this library, including changes to the exported Go API surface or behavior of the API.
* We increment the minor version with any backwards-compatible changes to functionality.
* We increment the patch version with any backwards-compatible bug fixes.

## Road map

This library was initially developed for internal applications at Target, so API methods will likely be added in the order that they are required.

## Contributing

We always welcome new PRs! See [_Contributing_](.github/CONTRIBUTING.md) for further instructions.

## Bugs and Feature Requests

Found something that doesn't seem right or have a feature request? [Please open a new issue](../../issues/new/).

## Copyright and License

[![license](https://img.shields.io/crates/l/gl.svg)](LICENSE)

Copyright (c) 2018 Target Brands, Inc.
