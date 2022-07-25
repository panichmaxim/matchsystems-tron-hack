CI_REGISTRY_IMAGE := registry.gitlab.com/rubin-dev/api
CI_COMMIT_REF_SLUG := dev
BUILDKIT_BUILD := buildctl build \
	--progress=plain \
	--frontend=dockerfile.v0 \
	--local context=. \
	--local dockerfile=. \
	--opt build-arg:CI_COMMIT_SHORT_SHA=${CI_COMMIT_SHORT_SHA}

.PHONY: docker-all
docker-all: docker-app

.PHONY: docker-app
docker-app:
	$(BUILDKIT_BUILD) \
		--opt filename=./app.dockerfile \
		--import-cache type=registry,ref=${CI_REGISTRY_IMAGE}/app:${CI_COMMIT_REF_SLUG}-buildcache \
		--export-cache type=registry,ref=${CI_REGISTRY_IMAGE}/app:${CI_COMMIT_REF_SLUG}-buildcache,push=true \
		--output type=image,\"name=${CI_REGISTRY_IMAGE}/app:${CI_COMMIT_REF_SLUG},${CI_REGISTRY_IMAGE}/app:${CI_COMMIT_SHORT_SHA}\",push=true
