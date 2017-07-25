all: test

install:
	go install -v ./...

test: install
	cd internal/test/package1 && go test -coverprofile=package1.out -covermode=count
	cd internal/test/package2 && go test -coverprofile=package2.out -covermode=count
	gocovermerge -coverprofile=cover.out test -v -covermode=count \
		github.com/AlekSi/gocovermerge/internal/test/package1 \
		github.com/AlekSi/gocovermerge/internal/test/package2 \
		github.com/AlekSi/gocovermerge/internal/test/...
	go tool cover -html=cover.out -o cover.html

merge: install
	gocovermerge -coverprofile=cover.out merge internal/test/package1/package1.out internal/test/package2/package2.out
