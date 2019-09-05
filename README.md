# genelizer [![CircleCI](https://circleci.com/gh/Matts966/genelizer.svg?style=svg)](https://circleci.com/gh/Matts966/genelizer) [![codecov](https://codecov.io/gh/Matts966/genelizer/branch/master/graph/badge.svg)](https://codecov.io/gh/Matts966/genelizer)

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
