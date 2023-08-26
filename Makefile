.PHONY: init
# init env
init:
	find app -maxdepth 1 -mindepth 1 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) init'
.PHONY: api
# generate api
api:
	find app -maxdepth 1 -mindepth 1 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: proto
# generate proto
proto:
	find app -maxdepth 1 -mindepth 1 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'

.PHONY: docker
# docker build
docker:
	find app -maxdepth 1 -mindepth 1 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) docker'

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: service
# create configs
service:
	kratos new app/$(name) --nomod
	kratos proto add api/$(name)/v1/$(name).proto && \
	kratos proto client api/$(name)/v1/$(name).proto && \
	kratos proto server api/$(name)/v1/$(name).proto -t app/$(name)/internal/service && \
	cd app/$(name) && echo "include ../../app_makefile" > ./Makefile && touch README.md && cd ../../

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;