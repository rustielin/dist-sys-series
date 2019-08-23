FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
# XXX: for now, do this manually, as we copy in the entire project 
# RUN go build -o host2 .

# spin forever, connect manually
CMD ["tail", "-f", "/dev/null"]
