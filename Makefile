test:
	go test -v -race ./...

cover:
	rm -rf *.coverprofile
	go test -v -race -coverprofile=datastructures.coverprofile
	go test -v -race -coverprofile=bst.coverprofile ./binary-search-tree
	go test -v -race -coverprofile=mh.coverprofile ./min-heap
	go test -v -race -coverprofile=bt.coverprofile ./b-tree
	gover
	go tool cover -html=gover.coverprofile

.PHONY: test cover