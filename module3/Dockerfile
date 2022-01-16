From golang as build
RUN mkdir /go_builder
COPY ./module2/* /go_builder/
WORKDIR /go_builder/
RUN go build server.go

From ubuntu 
RUN mkdir /http_server
COPY --from=build /go_builder/server /http_server/
WORKDIR /http_server/
CMD ./server
