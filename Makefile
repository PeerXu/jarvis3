all:

build_contrib_generate_container_impl:
	mkdir -p dist/contrib
	(cd contrib/container/generate_container_impl; go build -o ../../../dist/contrib/generate_container_impl)

clean_contrib_generate_container_impl:
	rm -rf dist/contrib/generate_container_impl

generate_container:
	dist/contrib/generate_container_impl

generate_agent:
	mkdir -p dist/agent
	(cd agent; gopherjs build -o ../dist/agent/main.js)

build_service:
	go build -o dist/jarvis3

clean_service:
	rm -rf dist/jarvis3
