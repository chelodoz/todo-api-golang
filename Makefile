check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_install
	GO111MODULE=off swagger generate spec -o ./docs/swagger.yaml --scan-models
	GO111MODULE=off swagger generate spec -o ./third_party/swagger-ui-4.11.1/swagger.json --scan-models