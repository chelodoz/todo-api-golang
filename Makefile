check_swagger_install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_swagger_install
	swagger generate spec -o ./docs/swagger.yaml --scan-models
	swagger generate spec -o ./third_party/swagger-ui-4.11.1/swagger.json --scan-models

check_mockery_install:
	which mockery || go install github.com/vektra/mockery/v2@latest
mocks: check_mockery_install
	cd internal && mockery --all

build:
	docker compose --env-file ./config/app.env up --build