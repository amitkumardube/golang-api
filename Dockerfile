# choose the base image
FROM golang:1.16
# Add the maintainer
maintainer "kshitiz DUBE golang developer from INDIA."
# Working Directory for code
WORKDIR /src
# Add code under working directory
ADD main.go .
# Build the go code
RUN go build -o /bin/main
# This gets executed when container is started
ENTRYPOINT ["/bin/main"]