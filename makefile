buildDir := build

compile:
	go build ./...
race:
	go test -v -race ./...
test: 
	go test -v -cover ./...
coverage:$(buildDir)/cover.out
	go tool cover -func=$< | sed -E 's%github.com/.*/shrub/%%' | column -t
coverage-html:$(buildDir)/cover.html

$(buildDir):
	mkdir -p $@
$(buildDir)/cover.out:$(buildDir)
	go test -coverprofile $@ -cover ./...
$(buildDir)/cover.html:$(buildDir)/cover.out
	go tool cover -html=$< -o $@
