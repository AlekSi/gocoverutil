# gocovermerge

[![Build Status](https://travis-ci.org/AlekSi/gocovermerge.svg?branch=master)](https://travis-ci.org/AlekSi/gocovermerge)
[![codecov](https://codecov.io/gh/AlekSi/gocovermerge/branch/master/graph/badge.svg)](https://codecov.io/gh/AlekSi/gocovermerge)
[![coveralls](https://coveralls.io/repos/github/AlekSi/gocovermerge/badge.svg?branch=master)](https://coveralls.io/github/AlekSi/gocovermerge)

Install it with `go get`:
```
go get -u github.com/AlekSi/gocovermerge
```

gocovermerge contains two commands: merge and test.

Merge command merges several go coverage profiles into single file.
Run `gocovermerge merge -h` for usage information.

Test command runs go test -cover with correct flags and merges profiles.
Run `gocovermerge test -h` for usage information.
