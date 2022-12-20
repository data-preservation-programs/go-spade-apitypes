.PHONY: $(MAKECMDGOALS)

build: webapi

webapi: genapitypes

genapitypes:
	go generate ./types.go
