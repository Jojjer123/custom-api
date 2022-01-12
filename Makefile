include ./common.mk

# Tools
.PHONY: coverage
coverage: overalls | $(GOVERALLS) ; $(info $(M) running coveralls) @ ## run coveralls (PROJECT)
	$Q $(GOVERALLS) -coverprofile=overalls.coverprofile -service=travis-ci

.PHONY: clean
clean:: ; $(info $(M) custom-api clean) @ ## clean (ADDITIONAL)
	@rm -rf  build/_output
	rm -rf pkg

.PHONY: images
images: docker-$(PRJ_NAME) ; $(info $(M) building images...) @ ## build all docker images (ADDITIONAL)

.PHONY: images-push
images-push: images $(DOCKER_LOGIN) ; $(info $(M) pushing images...) @ ## push docker images (PROJECT)
	docker push onosproject/$(PRJ_NAME):$(PRJ_VERSION)

.PHONY: kind
kind: images ; $(info $(M) add images to kind cluster...) @ ## add images to kind (ADDITIONAL)
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/$(PRJ_NAME):$(PRJ_VERSION)

generate:
	cd api; buf generate

run:
	cd cmd; go run main.go