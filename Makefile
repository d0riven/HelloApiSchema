DOCKER := docker
SED := $(DOCKER) run --rm -v $(PWD):/app -w /app ubuntu:latest sed

URL := http://localhost:8080

.PHONY: FORCE
FORCE:

.PHONY: clean
clean:
	-rm .swagger-editor.id

.PHONY: run
run: .swagger-editor.id
.swagger-editor.id:
	$(DOCKER) run --rm -d -p 8080:8080 --name swagger-editor swaggerapi/swagger-editor | tee -i $@

.PHONY: rm
rm:
	$(DOCKER) rm -f $$(cat .swagger-editor.id)
	rm .swagger-editor.id

.PHONY: validate/openapi generate/openapi/% api/go-openapi-server
generate: generate/server generate/client
generate/server: validate/openapi
	$(DOCKER) run --rm -v $(PWD):/app openapitools/openapi-generator-cli generate \
		-c /app/api/go-server-config.json \
		-i /app/api/openapi-schema/openapi.yaml \
		-g go-server \
		-o /app/server/golang
generate/client: validate/openapi
	$(DOCKER) run --rm -v $(PWD):/app openapitools/openapi-generator-cli generate \
		-c /app/api/go-server-config.json \
		-i /app/api/openapi-schema/openapi.yaml \
		-g typescript-fetch \
		-o /app/client/ts/src/api-client
validate/openapi: api/openapi-schema/openapi.yaml
	$(DOCKER) run --rm -v $(PWD):/app openapitools/openapi-generator-cli validate \
		-i /app/$<
