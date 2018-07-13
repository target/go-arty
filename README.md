# go-arty

[![GoDoc](https://godoc.org/github.com/target/go-arty?status.svg)](https://godoc.org/github.com/target/go-arty)
[![Go Report Card](https://goreportcard.com/badge/target/go-arty)](https://goreportcard.com/report/target/go-arty)
[![Build Status](https://travis-ci.org/target/go-arty.svg?branch=master)](https://travis-ci.org/target/go-arty)

go-arty is a Go client library for accessing the [Artifactory](https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API) and [Xray](https://www.jfrog.com/confluence/display/XRAY/Xray+REST+API) API.

## Usage

```go
import "github.com/target/go-arty"
```

Construct a new Artifactory client, then use the various services on the client to access different parts of the Artifactory API. For example:

```go
client, _ := artifactory.NewClient("artifactory.company.com", nil)

// list all repositories from the artifactory server
repos, _, err := client.Repositories.GetAll()
```

### Authentication

The go-arty library allows you to pass basic auth or a [API Key](https://www.jfrog.com/confluence/display/RTF/Updating+Your+Profile#UpdatingYourProfile-APIKey).

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

## Road map

This library was initially developed for internal applications at Target, so API methods will likely be added in the order that they are required.

## Contributing

We always welcome new PRs! See [_Contributing_](CONTRIBUTING.md) for further instructions.

## Bugs and Feature Requests

Found something that doesn't seem right or have a feature request? [Please open a new issue](issues/new/).

## Copyright and License

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE)

Copyright (c) 2018 Target Brands, Inc.
