DOCKER := docker
SWAGGER_EDITOR := $(DOCKER) run --rm -d -p 8180:8080 \
	--name swagger-editor swaggerapi/swagger-editor
SWAGGER_UI := $(DOCKER) run --rm -d -p 8280:8080 \
	-v $(PWD)/api/openapi-schema/openapi.yaml:/app/openapi.yaml \
	-e SWAGGER_JSON=/app/openapi.yaml \
	--name swagger-ui swaggerapi/swagger-ui
OPENAPI_GENERATOR_CLI := $(DOCKER) run --rm -v $(PWD):/app openapitools/openapi-generator-cli

.PHONY: FORCE
FORCE:

.PHONY: clean
clean: rm

.PHONY: run run/swagger-editor run/swagger-ui
run: run/swagger-editor run/swagger-ui
run/swagger-editor: .swagger-editor.id
.swagger-editor.id:
	$(SWAGGER_EDITOR) | tee -i $@
run/swagger-ui: .swagger-ui.id
.swagger-ui.id:
	$(SWAGGER_UI) | tee -i $@

.PHONY: rm rm/swagger-editor rm/swagger-ui
# TODO: --rm オプションをつけているのでコンテナIDを使わずstopでコンテナ名を指定すればコンテナID不要になる
rm: rm/swagger-editor rm/swagger-ui
rm/swagger-editor:
	-$(DOCKER) rm -f $$(cat .swagger-editor.id)
	-rm .swagger-editor.id
rm/swagger-ui:
	-$(DOCKER) rm -f $$(cat .swagger-ui.id)
	-rm .swagger-ui.id

.PHONY: generate generate/server generate/client
generate: generate/server generate/client
generate/server: validate/openapi
	$(OPENAPI_GENERATOR_CLI) generate \
		-c /app/api/go-server-config.json \
		-i /app/api/openapi-schema/openapi.yaml \
		-g go-server \
		-o /app/server/golang/openapi
generate/client: validate/openapi
	$(OPENAPI_GENERATOR_CLI) generate \
		-c /app/api/go-server-config.json \
		-i /app/api/openapi-schema/openapi.yaml \
		-g typescript-fetch \
		-o /app/client/ts/openapi/src/api-client

.PHONY: validate/openapi
validate/openapi: api/openapi-schema/openapi.yaml
	$(OPENAPI_GENERATOR_CLI) validate \
		-i /app/$<
