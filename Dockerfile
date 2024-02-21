# Use the pandoc/extra image as the base
FROM pandoc/extra:latest

# Install Go
RUN apk add --update --no-cache git go
# RUN apk add --update --no-cache go git make musl-dev curl

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:$PATH
RUN mkdir -p ${GOPATH}/src ${GOBIN}

# Set the working directory
WORKDIR /app

# install reviews binary
RUN go install github.com/rahji/reviews@latest

# clone reviews repo into /app and create data directories
RUN git clone https://github.com/rahji/reviews.git
RUN mkdir -p /app/data/student/markdown /app/data/student/pdf 
RUN mkdir -p /app/data/director/markdown /app/data/director/pdf

# Set the entrypoint for the container
ENTRYPOINT ["/app/reviews/bin/reviewrunner.sh"]

