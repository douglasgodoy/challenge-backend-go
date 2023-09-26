FROM golang:1.21.1


WORKDIR /api
COPY . .

RUN cd src/
RUN go install 
RUN go get -d 

CMD ["app"]
