.PHONY: build-local codegen

build-local:
	docker compose up -d --build

OPENAPI_DOCS_DIR := docs/openapi
OPENAPI_OUTPUT_DIR := api

YAML_FILES := $(wildcard $(OPENAPI_DOCS_DIR)/*.yaml)

# openapiの*.yamlからapi/*/gen/*.gen.goを生成
codegen:
	@for file in $(YAML_FILES); do \
		dir=$$(dirname $$file | sed 's|$(OPENAPI_DOCS_DIR)|$(OPENAPI_OUTPUT_DIR)|'); \
		base=$$(basename $$file .yaml); \
		mkdir -p $$dir/$$base/gen; \
		oapi-codegen -package gen $$file > $$dir/$$base/gen/$$base.gen.go; \
	done