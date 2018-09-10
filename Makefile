
OS?=amd64
PLAT?=linux

GO=go
cmd=install

SNAKY_VERSION=1.0.1

BUILD_FLAGS=GOOS=$(PLAT) GOARCH=$(OS)

ECHO=echo -e

Q=@
ifeq ("$V",  "1")
Q=
endif

SNAKYDIRS=html/addon
SNAKYDIRS+=html/js/wssimple
SNAKYDIRS+=html/wssimple
SNAKYDIRS+=html/manager
SNAKYFILES=html/css/global.css html/css/jmdkina.css
SNAKYFILES+=html/css/resume.css html/css/shici.css html/css/db.css 
SNAKYFILES+=html/css/wssimple.css
SNAKYFILES+=html/jmdkina/add.html html/jmdkina/jmdkina.html 
SNAKYFILES+=html/shici/add.html html/shici/shici.html 
SNAKYFILES+=html/db/db.html 
SNAKYFILES+=html/resume/resume.html 
SNAKYFILES+=html/resume/resume_en.html
SNAKYFILES+=html/resume/resume_set.html
SNAKYFILES+=html/resume/template.json
SNAKYFILES+=html/resume/template.json.back
SNAKYFILES+=html/404.html 
SNAKYFILES+=html/js/jmdkina/add.js html/js/jmdkina/jmdkina.js 
SNAKYFILES+=html/js/shici/add.js html/js/shici/shici.js 
SNAKYFILES+=html/js/global.js  html/js/db.js
SNAKYFILES+=html/js/resume.js
SNAKYSOURCE=$(shell pwd)
SNAKYDST=$(shell pwd)/../bin/snaky-bin

FILES=github.com/astaxie/beego github.com/beego/bee 
FILES+= code.google.com/p/graphics-go/graphics
FILES+= goconfig/config golanger.com/log golanger.com/utils
FILES+= jk/jkcommon jk/jkimage jk/jklog jk/jkmath jk/jkprotocol
FILES+= github.com/beego/i18n
FILES+= github.com/tyranron/daemonigo
FILES+= jk/jkeasycrypto
#FILES+=" github.com/jeffallen/mqtt github.com/surgemq/surgemq github.com/surgemq/surgemq/service"
#FILES+=" github.com/deckarep/gosx-notifier"

# go get golang.org/x/mobile/cmd/gomobile
# gomobile init
# go get -d golang.org/x/mobile/example/basic
# gomobile build -target=android golang.org/x/mobile/example/basic
# gomobile install golang.org/x/mobile/example/basic
# gomobile build -target=ios golang.org/x/mobile/example/basic
# ios-deploy -b basic.app
# go get -d golang.org/x/mobile/example/bind/...
# git clone https://github.com/beego/i18n.git
# go get github.com/unidoc/unidoc/... ## pdf operation

help:
	$(Q)$(ECHO) make jkavdu/snaky
	$(Q)$(ECHO) PLAT=linux/windows/... OS=amd64/i386/arm64/...
	$(Q)$(ECHO)     base - build all base
	$(Q)$(ECHO)     jkavdu - av distribution
	$(Q)$(ECHO)     snaky - simpleserver for http
	$(Q)$(ECHO)	    copysnaky - copy snaky files to snaky-bin if nesessary

jkavdu: 
	$(BUILD_FLAGS) $(GO) $(cmd) jkprog/jkavdu

snaky: 
	$(BUILD_FLAGS) $(GO) build -ldflags "-X main.VERSION=$(SNAKY_VERSION) -X 'main.BUILD_TIME=$(shell date)' -X 'main.GOVERSION=$(shell go version)'" simpleserver/simpleserver
	mv simpleserver bin/snaky

copysnaky:
	$(Q) for it in $(SNAKYDIRS); do \
		cp -rfv $(SNAKYSOURCE)/$$it $(SNAKYDST)/$${it%/*}; \
		done
	$(Q) for it in $(SNAKYFILES); do \
		cp -rfv $(SNAKYSOURCE)/$$it $(SNAKYDST)/$$it; \
		done
	$(Q) cp -rfv $(SNAKYSOURCE)/bin/snaky $(SNAKYDST)/bin

base:
	$(Q) for it in $(FILES); do \
		echo go install $$it; \
		go install $$it ; \
	done
	$(Q)$(ECHO) Install done

