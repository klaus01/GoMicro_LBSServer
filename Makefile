.PHONY: proto build

proto:
	for d in api srv; do \
		for f in $$d/**/proto/*.proto; do \
			protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done

lint:
	./bin/lint.sh

build:
	./bin/build.sh

# run:
# 	docker-compose build
# 	docker-compose up
