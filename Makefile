PLUGIN_VERSIONS=1 10

build_plugins:
	@for ver in $(PLUGIN_VERSIONS); do \
		go build -buildmode=plugin -o tmp/plugin$$ver.so modules/plugins$$ver/plugin.go; \
	done

build_host:
	@go build -o tmp/main main.go

.PHONY: run
run: build_host build_plugins
	@cd tmp;\
	./main