GLOLANGCI_LINT_VERSION=v1.30.0

test: .FORCE
	go test -v ./test/

.PHONY: .FORCE

install-test:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin $(GLOLANGCI_LINT_VERSION)

lint:
	$(GOPATH)/bin/golangci-lint run -v
