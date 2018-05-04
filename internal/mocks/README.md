# mocks [![GoDoc](https://godoc.org/github.com/ddspog/mongo/internal/mocks?status.svg)](https://godoc.org/github.com/ddspog/mongo/internal/mocks) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/mongo/internal/mocks)](https://goreportcard.com/report/github.com/ddspog/mongo/internal/mocks) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)

## Overview

Package mocks defines mocks of mgo objects.

This package creates mocks objects that implements the same functions
as the mgo objects. It implements interface defined on
[github.com/ddspog/mongo/elements](https://github.com/ddspog/mongo/tree/master/elements) package.

## Usage

The package encourages to use the mock helper object MockMGOSetup, to call all the mocking methods returning the desired values. It was made to reduce number of lines when using the package.