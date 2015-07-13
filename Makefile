default: test

deps:
	go get golang.org/x/crypto/openpgp/clearsign
	go get github.com/stretchr/testify/assert

test:
	./gen-key.sh
	go test
