.PHONY: build push

REGISTRY_DOMAIN := jun06t

build:
	docker build -f Dockerfile -t ${REGISTRY_DOMAIN}/probe:latest .

push:
	docker push ${REGISTRY_DOMAIN}/probe

