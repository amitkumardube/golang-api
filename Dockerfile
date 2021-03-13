# choose the base image
FROM golang:1.16
# Add the maintainer
maintainer "kshitiz DUBE golang developer from INDIA."
# Working Directory for code
WORKDIR /src
# Add code under working directory
ADD main.go .
ADD go.mod .
ADD go.sum .
# Build the go code
RUN go build -o /bin/main
# Expose port 8080 from container
EXPOSE 8080
# This gets executed when container is started
ENTRYPOINT ["/bin/main"]