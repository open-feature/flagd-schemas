BUF_VERSION=v1.6.0

guard-%:
	@ if [ "${${*}}" = "" ]; then \
        echo "Environment variable $* not set"; \
        exit 1; \
    fi

install-buf:
	go install github.com/bufbuild/buf/cmd/buf@${BUF_VERSION}

install-yq-linux:
	wget https://github.com/mikefarah/yq/releases/download/v4.2.0/yq_linux_amd64 -O /usr/bin/yq &&\
		chmod +x /usr/bin/yq

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

gen-ruby: install-buf guard-GOPATH
	${GOPATH}/bin/buf generate buf.build/open-feature/flagd --template protobuf/buf.gen.ruby.yaml

gen-schema-yaml:
	yq eval -j json/flagd-definitions.json | yq -P > json/flagd-definitions.yaml
