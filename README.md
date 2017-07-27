# gocovermerge

[![Release](https://github-release-version.herokuapp.com/github/AlekSi/gocovermerge/release.svg?style=flat)](https://github.com/AlekSi/gocovermerge/releases/latest)
[![Travis CI](https://travis-ci.org/AlekSi/gocovermerge.svg?branch=master)](https://travis-ci.org/AlekSi/gocovermerge)
[![AppVeyor](https://ci.appveyor.com/api/projects/status/bxcbywwapyvsprju/branch/master?svg=true)](https://ci.appveyor.com/project/AlekSi/gocovermerge)
[![Codecov](https://codecov.io/gh/AlekSi/gocovermerge/branch/master/graph/badge.svg)](https://codecov.io/gh/AlekSi/gocovermerge)
[![Coveralls](https://coveralls.io/repos/github/AlekSi/gocovermerge/badge.svg?branch=master)](https://coveralls.io/github/AlekSi/gocovermerge)
[![Go Report Card](https://goreportcard.com/badge/AlekSi/gocovermerge)](https://goreportcard.com/report/AlekSi/gocovermerge)


Install it with `go get`:
```
go get -u github.com/AlekSi/gocovermerge
```

gocovermerge contains two commands: merge and test.

Merge command merges several go coverage profiles into a single file.
Run `gocovermerge merge -h` for usage information. Example:
```
gocovermerge -coverprofile=cover.out merge internal/test/package1/package1.out internal/test/package2/package2.out
```

Test command runs `go test -cover` with correct flags and merges profiles.
Packages list is passed as arguments; they may contain `...` patterns.
The list is expanded, sorted and duplicates are removed.
`go test -coverpkg` flag is set automatically to include all packages.
If tests are failing, gocovermerge exits with a correct exit code.
Run `gocovermerge test -h` for usage information. Example:
```
gocovermerge -coverprofile=cover.out test -v -covermode=count github.com/AlekSi/gocovermerge/internal/test/...
```
