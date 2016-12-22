test:
	go test -v -race ./...

cover:
	rm -rf *.coverprofile
	go test -v -race -coverprofile=datastructures.coverprofile
	go test -v -race -coverprofile=bst.coverprofile ./binary-search-tree
	go test -v -race -coverprofile=bh.coverprofile ./binary-heap
	gover
	go tool cover -html=gover.coverprofile

.PHONY: test cover