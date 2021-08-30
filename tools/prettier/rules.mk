prettier_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
prettier_version := 2.3.2
prettier_dir := $(prettier_cwd)/$(prettier_version)
prettier := $(prettier_dir)/node_modules/.bin/prettier

$(prettier):
	npm install --no-save --no-audit --prefix $(prettier_dir) prettier@$(prettier_version)
	chmod +x $@
	touch $@
