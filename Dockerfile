FROM debian:buster

# Install updates, Golang
RUN apt-get update && apt-get upgrade -y
RUN apt-get install ca-certificates golang -y

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Binrecon Prep
RUN mkdir ${GOPATH}/src/binrecon
WORKDIR ${GOPATH}/src/binrecon
COPY src ${GOPATH}/src/binrecon
RUN go get -d -v ./...
RUN go install .

# Run Binrecon
CMD ["binrecon"]