
VERSION ?= v$(shell cat .version)

version:
	git add . && git commit -m "chroe: $(VERSION)" && git tag $(VERSION) && git push origin $(VERSION)
