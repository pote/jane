SOURCES = *.go src/**/*.go
DEPS = $(firstword $(subst :, ,$(GOPATH)))/up-to-date
GPM ?= gpm

all: bin/jane

bin/jane:  bin $(SOURCES) $(DEPS) | $(dir $(PROGNAME))
	go build -o bin/jane

$(DEPS): Godeps | $(dir $(DEPS))
	$(GPM) get
	touch $@

config.mk:
	@./configure

install: bin/jane
	install -d $(prefix)/bin
	install -m 0755 bin/jane /usr/local/bin

cross-compile: clean pkg
	script/cross-compile

clean:
	rm -rf pkg/

##
# Directories
##

$(dir $(PROGNAME)) $(dir $(DEPS)) bin pkg:
	mkdir -p $@
