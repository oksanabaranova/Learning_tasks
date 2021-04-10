FROM golang:1.16-alpine AS build

WORKDIR /app/
#COPY main.go go.* /app/
COPY . /app
RUN go build

FROM alpine
WORKDIR /app/ 
COPY script.sh /app 
COPY --from=build /app/task5_yaml /app/task5_yaml
ENTRYPOINT ["/app/script.sh"] 