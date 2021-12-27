buf_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
buf := $(buf_cwd)/bin/buf

buf_version := 1.0.0-rc10

arch = $(shell uname -s)-$(shell uname -m)

buf_bin_url := https://github.com/bufbuild/buf/releases/download/v$(buf_version)/buf-$(arch)

$(buf): $(buf_cwd)/rules.mk
	$(info [buf] feching $(buf_version) binary...)
	@mkdir -p $(dir $@)
	@curl -sSL $(buf_bin_url) -o $@
	@chmod +x $@
	@touch $@
