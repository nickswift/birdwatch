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
BIN_PUBL     = ${DIR_INSTALL}/publ
BIN_DELOREAN = ${DIR_INSTALL}/delorean

BINS = \
  ${BIN_PUBL} \
  ${BIN_DELOREAN}

# Install Dependencies
all: deps ${BINS}

deps:
	@echo "GO-GETTING DEPENDENCIES"
	@go get ${DEPS}

${BIN_PUBL}:
	go install ${SRC_PUBLICIST}

${BIN_DELOREAN}:
	go install ${SRC_DELOREAN}

clean:
	rm -f ${BIN_PUBL} ${BIN_DELOREAN}

fmt:
	gofmt -s -w .