# Contributing to go-arty

We'd love to accept your contributions to this project! There are just a few small guidelines you need to follow.

## Issues

[Issues](issues/new/) are always welcome!

## Pull Requests

**NOTE: We recommend you start by opening a new issue describing the bug or feature you're intending to fix. Even if you think it's relatively minor, it's helpful to know what people are working on.**

These rules must be followed for any contributions to be merged into master. A Git installation is required.

1. Fork this repo
1. Go get the original code:

  `go get github.com/target/go-arty`

1. Navigate to the original code:

  `$GOPATH/src/github.com/target/go-arty`

1. Add a remote branch pointing to your fork:

  `git remote add fork https://github.com/your_fork/go-arty`

1. Implement desired changes
1. Validate the changes meet your desired use case
1. Write tests around the changes you made
1. Update documentation
1. Please run the below commands:

```sh
# Generate necessary code
go generate github.com/target/go-arty/...

# Test the code
go test github.com/target/go-arty/...

# Format the code
go fmt github.com/target/go-arty/...

# Vet the code
go vet github.com/target/go-arty/...
```

1. Push to your fork:

  `git push fork master`

1. Open a pull request. Thank you for your contribution! A dialog will ensue.
