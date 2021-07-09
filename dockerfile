FROM golang  
WORKDIR /work
ADD . .
RUN go test ./...
RUN go build /cmd -o /bin/myapp .
WORKDIR /
RUN rm -r /work
CMD ["/bin/myapp"]  
