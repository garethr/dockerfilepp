A simple library for building Dockerfile pre-processors.

[![Build
Status](https://travis-ci.org/garethr/dockerfilepp.svg)](https://travis-ci.org/garethr/dockerfilepp)
[![Go Report
Card](https://goreportcard.com/badge/github.com/garethr/dockerfilepp)](https://goreportcard.com/report/github.com/garethr/dockerfilepp)
[![GoDoc](https://godoc.org/github.com/garethr/dockerfilepp?status.svg)](https://godoc.org/github.com/garethr/dockerfilepp)

Dockerfile pre-processors aim to make extending Dockerfile with new
domain-specific instructions easy. Simply build your own pre-processor
and pipe Dockerfiles in on stdin.

Here is a simple hello world usage of the API:

```go
import (
    "github.com/garethr/dockerfilepp"
)

replacements := map[string]string{
    "CUSTOM_INSTALL": "RUN apt-get install thing",
    "CUSTOM_ECHO": "RUN echo hello",
}
dockerfilepp.Process(replacements, "include comprehensive help text here")
```

As a example implementation see [dockerfilepp-puppet](https://github.com/garethr/dockerfilepp-puppet)
which uses local files compiled into the binary rather than hard coded
variables.
