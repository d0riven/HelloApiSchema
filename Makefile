DOCKER := docker

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
generate/openapi: api/go-openapi-server/api/openapi.yaml
api/go-openapi-server/api/openapi.yaml: validate/openapi api/openapi-schema/openapi.yaml api/go-server-config.json
	$(DOCKER) run --rm -v $(PWD):/app openapitools/openapi-generator-cli generate \
		-c /app/api/go-server-config.json \
		-i /app/api/openapi-schema/openapi.yaml \
		-g go-server \
		-o /app/$(@:%/api/openapi.yaml=%) \
		-Dmodels
validate/openapi: api/openapi-schema/openapi.yaml
	$(DOCKER) run --rm -v $(PWD):/app openapitools/openapi-generator-cli validate \
	  -i /app/$<
