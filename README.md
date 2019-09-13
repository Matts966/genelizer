# genelizer 
[![CircleCI][circleci-badge]][circleci]
[![codecov][codecov-badge]][codecov]
[![godoc.org][godoc-badge]][godoc]
[![GolangCI][golangci-badge]][golangci]
[![LICENSE][license-badge]][license]
[![Go Report Card][go-report-card-badge]][go-report-card]

## Quick Start

```
$ git clone https://github.com/Matts966/genelizer.git && cd genelizer
```

Now you can edit `./config/sample.hcl` or add other files in `./config` directory to change analysis and

```
$ YOUR_BINARY_NAME=binary-name make
```

Then you can get your portable binary named `binary-name`!

Or install it in your path by

```
$ YOUR_BINARY_NAME=binary-name make install
```

Also Dockerfile is located for test use.

```
$ make docker
```

## Analyzer and Goroutine

The `golang.org/x/tools/go/analysis` package runs `analysis.Analyzer` concurrently per packages using goroutine and waitgroup.
See the code doing it [here](https://github.com/golang/tools/blob/be0da057c5e3c2df569a2c25cd280149b7d7e7d0/go/analysis/internal/checker/checker.go#L201).

For utilizing this feature of `golang.org/x/tools/go/analysis` , `genelizer` generates `analysis.Analyzer` for each rule in the config.

[circleci-badge]: https://circleci.com/gh/Matts966/genelizer.svg?style=svg
[circleci]: https://circleci.com/gh/Matts966/genelizer
[codecov]: https://codecov.io/gh/Matts966/genelizer/branch/master/graph/badge.svg
[codecov-badge]: https://codecov.io/gh/Matts966/genelizer
[godoc]: https://godoc.org/github.com/Matts966/genelizer
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org
[golangci]: https://golangci.com/r/github.com/Matts966/genelizer
[golangci-badge]: https://golangci.com/badges/github.com/Matts966/genelizer.svg
[license-badge]: https://img.shields.io/badge/License-MIT-yellow.svg
[license]: https://opensource.org/licenses/MIT
[go-report-card]: https://goreportcard.com/report/github.com/Matts966/genelizer
[go-report-card-badge]: https://goreportcard.com/badge/github.com/Matts966/genelizer
