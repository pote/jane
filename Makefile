SOURCES = *.go src/**/*.go
DEPS = $(firstword $(subst :, ,$(GOPATH)))/up-to-date
GPM ?= gpm

all: bin/jane
	./bin/jane

bin/jane:  bin $(SOURCES) $(DEPS) | $(dir $(PROGNAME))
	go build -o bin/jane

$(DEPS): Godeps | $(dir $(DEPS))
	$(GPM) get
	touch $@

##
# Directories
##

$(dir $(PROGNAME)) $(dir $(DEPS)) bin:
	mkdir -p $@
