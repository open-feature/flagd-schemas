BUF_VERSION=v1.6.0

guard-%:
	@ if [ "${${*}}" = "" ]; then \
        echo "Environment variable $* not set"; \
        exit 1; \
    fi

install-buf:
	go install github.com/bufbuild/buf/cmd/buf@${BUF_VERSION}

lint: guard-GOPATH
	cd protobuf && ${GOPATH}/bin/buf lint

gen-go: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.go.yaml

gen-go-server: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.go-server.yaml

gen-ts: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.ts.yaml

gen-java: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.java.yaml

gen-python: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.python.yaml

gen-csharp: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.csharp.yaml

gen-php: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.php.yaml
