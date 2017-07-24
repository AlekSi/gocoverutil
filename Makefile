all:
	go install -v ./...
	cd internal/test/package1 && go test -coverprofile=package1.out
	cd internal/test/package2 && go test -coverprofile=package2.out
	gocovermerge test -v \
		github.com/AlekSi/gocovermerge/internal/test/package1 \
		github.com/AlekSi/gocovermerge/internal/test/package2 \
		github.com/AlekSi/gocovermerge/...
	go tool cover -html=cover.out -o cover.html
