MODULES=activity l10n responder sorting audited location roles transition exchange media_library seo validations i18n qor serializable_meta worker inflection slug publish admin widget render cache bindatafs action_bar filebox notification wildcard_router gomerchant app log

QORTHEME=activity i18n l10n location media_library publish seo serializable_meta slug sorting widget worker

DB_NAME=database.dev.yml

OSNAME=$(shell uname)

GO=$(shell which go)

CUR_TIME=$(shell date '+%Y-%m-%d_%H:%M:%S')
# Program version
VERSION=$(shell cat VERSION)

# Binary name for bintray
BIN_NAME=$(shell basename $(abspath ./))
BIN_NAME_CLI=qor-cli

# Project name for bintray
PROJECT_NAME=$(shell basename $(abspath ./))
PROJECT_DIR=$(shell pwd)

# Project url used for builds
# examples: github.com, bitbucket.org
REPO_HOST_URL=github.com.org

# Grab the current commit
GIT_COMMIT="$(shell git rev-parse HEAD)"

# Check if there are uncommited changes
GIT_DIRTY="$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)"

QOR_DIR="${PROJECT_DIR}/dist/lib/src"
DIST_PUBLIC="${PROJECT_DIR}/dist/public"
QOR_REPO="${PROJECT_DIR}/../"

# Add the godep path to the GOPATH
#GOPATH=$(shell godep path):$(shell echo $$GOPATH)

default: help

help:
	@echo "..............................................................."
	@echo "Project: $(PROJECT_NAME) | current dir: $(PROJECT_DIR)"
	@echo "version: $(VERSION) GIT_DIRTY: $(GIT_DIRTY)\n"
	@#echo "Autocomplete exec -> PROG=$(BIN_NAME) source ./autocomplete/bash_autocomplete\n"
	@echo "make init    - Load godep"
	@echo "make save    - Save project libs"
	@echo "make git     - Git pull libs"
	@echo "make gor     - Download QOR repo"
	@echo "make install - Install packages"
	@echo "make clean   - Clean .orig, .log files"
	@echo "make run     - Run project debug mode"
	@echo "make seed    - Run project seeds"
	@echo "make cli     - Build qor-cli"
	@echo "make build   - Build for current OS project"
	@echo "make release - Build release project"
	@echo "make arm     - Build release project for ARM"
	@echo "make docs    - Project documentation"
	@echo "...............................................................\n"

init:
	@go get github.com/tools/godep

save:
	@godep save

install:
	@#go get -v -u github.com/constabulary/gb/...
	@#go get -v -u github.com/kr/godep
	@go get -v -u github.com/dgrijalva/jwt-go
	@go get -v -u github.com/antonholmquist/jason
	@go get -v -u github.com/go-resty/resty
	@go get -v -u github.com/gin-gonic/gin
	@go get -v -u github.com/itsjamie/gin-cors
	@#go get -v -u github.com/gin-gonic/contrib/jwt
	@go get -v -u github.com/codegangsta/cli
	@go get -v -u github.com/azumads/faker
	@go get -v -u github.com/jteeuwen/go-bindata/...
	@go get -v -u github.com/apertoire/mlog
	@go get -v -u github.com/microcosm-cc/bluemonday
	@go get -v -u github.com/jinzhu/gorm/dialects/mysql
	@go get -v -u github.com/jinzhu/gorm/dialects/postgres
	@go get -v -u github.com/jinzhu/gorm/dialects/sqlite
	@go get -v -u github.com/smartystreets/goconvey/convey
	@go get -v -u github.com/shopspring/decimal
	@go get -v -u github.com/tealeg/xlsx
	@go get -v -u github.com/justinas/nosurf
	@go get -v -u github.com/asaskevich/govalidator
	@#go get -v -u

qor:
	@#go get -v ./...
	@#echo ${PROJECT_DIR}
	@for a in $(MODULES); do echo "-> $$a"; cd ${PROJECT_DIR}/../ && test -e ./$$a || git clone https://github.com/qor/$$a.git; done

git:
	@for a in $(MODULES); do test ! -e ../$$a || echo "-> $$a"; test -e ../$$a || echo "Is NOT repo: https://github.com/qor/$$a"; test ! -e ../$$a || cd ../$$a && test ! -e ../$$a || git pull; done

view:
	@test ! -e ./dist || rm -R ./dist
	@mkdir -p ${DIST_PUBLIC}/admin/assets

	@mkdir -p ${QOR_DIR}/github.com/qor/admin/
	@cp -R ${QOR_REPO}admin/views ${QOR_DIR}/github.com/qor/admin/
	@cp -R ${QOR_DIR}/github.com/qor/admin/views/assets ${DIST_PUBLIC}/admin/

	@#mkdir -p ${QOR_DIR}/github.com/qor/media_library/
	@#cp -R ${QOR_REPO}l10n/views ${QOR_DIR}/github.com/qor/l10n/
	@#cp -R ${QOR_DIR}/github.com/qor/l10n/views/themes/l10n/assets ${DIST_PUBLIC}/admin/

	@for a in $(QORTHEME); do echo "--> $$a"; mkdir -p ${QOR_DIR}/github.com/qor/$$a; cp -R ${QOR_REPO}$$a/views ${QOR_DIR}/github.com/qor/$$a/; test ! -e ${QOR_DIR}/github.com/qor/$$a/views/themes/$$a/assets || cp -R ${QOR_DIR}/github.com/qor/$$a/views/themes/$$a/assets ${DIST_PUBLIC}/admin/; done

assets:
	@mkdir -p ./dist/config
	@cp ./config/database.yml ./dist/config/
	@mkdir -p ./dist/app/views
	@cp -R app/views/*.tmpl  ./dist/app/views/
	@cp -R app/views/qor  ./dist/app/views/
	@cp -R app/views/qor/assets  ./dist/public/admin/

release: clean view assets
	@mkdir -p ./dist/bin
	@cp -R ./public ./dist/
	@#go-bindata -nomemcopy ../qor/admin/views/...
	@echo "building release ${BIN_NAME} ${VERSION}"
	@GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o ./dist/$(BIN_NAME) main.go
	@echo "building release ${BIN_NAME_CLI} ${VERSION}"
	@GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o ./dist/bin/$(BIN_NAME_CLI) cmd/cli.go
	@chmod 0755 ./dist/bin/$(BIN_NAME_CLI)

arm: clean view assets
	@mkdir -p ./dist/bin
	@cp -R ./public ./dist/
	@#go-bindata -nomemcopy ../qor/admin/views/...
	@echo "building release ${BIN_NAME} ${VERSION}"
	@#CGO_ENABLED=0
	@GOOS=linux GOARCH=arm GOARM=7 go build -a -tags 'icu libsqlite3 linux netgo' -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o ./qor-server main.go
	@echo "building release ${BIN_NAME_CLI} ${VERSION}"
	@GOOS=linux GOARCH=arm GOARM=7 go build -a -tags 'icu libsqlite3 linux netgo' -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o ./dist/bin/$(BIN_NAME_CLI) cmd/cli.go
	@chmod 0755 ./dist/bin/$(BIN_NAME_CLI)

clean:
	@test ! -e ./${BIN_NAME} || rm ./${BIN_NAME}
	@git gc --prune=0 --aggressive
	@find . -name "*.orig" -type f -delete
	@find . -name "*.log" -type f -delete
	@test ! -e ./dist || rm -R ./dist

seed:
	@echo "...............................................................\n"
	@echo $(PROJECT_NAME) seed
	@echo ...............................................................
	@go run db/seeds/main.go

run:
	@echo "...............................................................\n"
	@echo Project: $(PROJECT_NAME) Path: ${PROJECT_DIR}
	@echo Open in browser:
	@echo	"	 http://localhost:7000/\n"
	@echo ...............................................................
	@QORCONFIG=${PROJECT_DIR}/config/${DB_NAME} go run main.go

test:
	@#QORCONFIG=${PROJECT_DIR}/config/${DB_NAME} GIN_MODE=release go test -v ./app/controllers/*_test.go
	@QORCONFIG=${PROJECT_DIR}/config/${DB_NAME} GIN_MODE=release go test -v ./app/models/*_test.go
	@#QORCONFIG=${PROJECT_DIR}/config/${DB_NAME} go test -v ./...
	@#API_PATH=$(PROJECT_DIR) ginkgo -v -r

build: clean
	@echo "Building ${BIN_NAME} ${VERSION}"
	@CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o $(BIN_NAME) main.go
	# @echo "Building ${BIN_NAME_CLI} ${VERSION}"
	# @CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -X cmd.BuildTime=${CUR_TIME} -X cmd.Version=${VERSION} -X cmd.GitHash=${GIT_COMMIT}' -o $(BIN_NAME_CLI) cmd/cli.go
	# @chmod 0755 ./$(BIN_NAME_CLI)


cli: clean
	@echo "Building cli ${VERSION}"
	@go build -a -tags netgo -ldflags '-w -X cmd.BuildTime=${CUR_TIME} -X cmd.Version=${VERSION} -X cmd.GitHash=${GIT_COMMIT}' -o $(BIN_NAME_CLI) cmd/cli.go
	@chmod 0755 ./$(BIN_NAME_CLI)
	@echo "PROG=$(BIN_NAME_CLI) source ./scripts/bash_autocomplete"
	@echo "export QORCONFIG=config/${DB_NAME}"
	@echo "export DEBUG=false"
	@echo "RUN: ./$(BIN_NAME_CLI)"

docs:
	godoc -http=:6060 -index

