# genelizer

## Quick Start

```
$ git clone https://github.com/Matts966/genelizer.git && cd genelizer
```

Now you can edit `./config/sample.hcl` or add other files in `./config` directory to change analysis.

```
$ go get -u github.com/gobuffalo/packr/v2/packr2
$ YOUR_BINARY_NAME=binary-name make
```

Then you can get your portable binary!

You can also get the binary like this if `packr2` is installed.

```
$ YOUR_BINARY_NAME=binary-name make
```

Or install it in your path by

```
$ YOUR_BINARY_NAME=binary-name make
```

## Analyzer and Goroutine

The `golang.org/x/tools/go/analysis` package runs `analysis.Analyzer` concurrently per packages using goroutine and waitgroup.
See the code doing it [here](https://github.com/golang/tools/blob/be0da057c5e3c2df569a2c25cd280149b7d7e7d0/go/analysis/internal/checker/checker.go#L201).

For utilizing this feature of `golang.org/x/tools/go/analysis` , `genelizer` generates `analysis.Analyzer` for each rule in the config.
