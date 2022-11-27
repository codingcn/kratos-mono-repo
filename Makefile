.PHONY: init
# init env
init:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) init'
.PHONY: api
# generate api
api:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: wire
# generate wire
wire:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'

.PHONY: docker
# docker build
docker:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) docker'

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: service
# create configs
service:
	kratos new app/$(name)/service --nomod
	kratos proto add api/$(name)/service/v1/$(name).proto && \
	kratos proto client api/$(name)/service/v1/$(name).proto && \
	kratos proto server api/$(name)/service/v1/$(name).proto -t app/$(name)/service/internal/service && \
	cd app/$(name)/service && echo "include ../../../app_makefile" > ./Makefile && touch README.MD && cd ../../../

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;