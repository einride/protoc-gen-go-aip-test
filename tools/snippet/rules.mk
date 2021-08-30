snippet_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
snippet := $(snippet_cwd)/bin/snippet

$(snippet): $(snippet_cwd)/main.go
	$(info building snippet binary...)
	@go build -o $@ $(snippet_cwd)
