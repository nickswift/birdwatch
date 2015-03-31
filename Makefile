# Birdwatch Twitter Tools -- Makefile
ME = github.com/nickswift498/birdwatch

# Dependencies
DEPS = \
  gopkg.in/yaml.v2 \
  github.com/chimeracoder/anaconda

# Sources
SRC_PUBLICIST = ${ME}/publ
SRC_DELOREAN  = ${ME}/delorean

# Build targets
DIR_INSTALL = ${GOPATH}/bin/build

# Binaries
BIN = ${DIR_INSTALL}/birdwatch

# Install Dependencies
all: deps ${BIN}

deps:
	@echo "GO-GETTING DEPENDENCIES"
	@go get ${DEPS}

${BIN}:
	go install

clean:
	rm -f ${BIN_PUBL} ${BIN_DELOREAN}

fmt:
	gofmt -s -w .