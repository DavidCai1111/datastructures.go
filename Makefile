test:
	go test -v -race ./...

cover:
	rm -rf *.coverprofile
	go test -v -race -coverprofile=datastructures.coverprofile
	go test -v -race -coverprofile=bst.coverprofile ./binary-search-tree
	go test -v -race -coverprofile=mh.coverprofile ./min-heap
	go test -v -race -coverprofile=bt.coverprofile ./b-tree
	go test -v -race -coverprofile=avl.coverprofile ./avl-tree
	go test -v -race -coverprofile=pat.coverprofile ./pat-tree
	go test -v -race -coverprofile=pat.coverprofile ./sort
	gover
	go tool cover -html=gover.coverprofile
	rm -rf *.coverprofile

.PHONY: test cover