
VERSION ?= v$(shell cat .version)


tidy:
	go mod tidy

version:
	git add . && git commit -m "chroe: $(VERSION)" && git tag $(VERSION) && git push origin $(VERSION)
