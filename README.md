# go-weasel

[![GoDoc](https://godoc.org/github.com/OxygenTrader/go-weasel?status.svg)](https://godoc.org/github.com/OxygenTrader/go-weasel)
[![Go Report Card](https://goreportcard.com/badge/github.com/OxygenTrader/go-weasel)](https://goreportcard.com/report/github.com/OxygenTrader/go-weasel)

[Weasel program](https://en.wikipedia.org/wiki/Weasel_program).

## Installation

```
$ go install github.com/OxygenTrader/go-weasel@latest
```

## Usage

```
$ go-weasel -h | -help | --help
$ go-weasel [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-sample STRING` &mdash; target string (default: `METHINKS IT IS LIKE A WEASEL`);
- `-rate FLOAT` &mdash; character mutate rate (default: `0.05`);
- `-count INTEGER` &mdash; population size (default: `100`).

## Output Example

```
0 HWICMFDFWUMHEXRVDIOPPERZCV F
10 JXICMNNF IM OSMVDKEKPIREAS F
20 MOICINKF IM IS LDKEVAIREAS F
30 MZIHINKA IM IS LDKEVA REASXL
40 MZIHINKS IT IS LGKEVA REASLL
50 METHINKS IT IS LGKEVA WEASEL
54 METHINKS IT IS LIKE A WEASEL
```

## License

The MIT License (MIT)

Copyright &copy; 2022 OxygenTrader
