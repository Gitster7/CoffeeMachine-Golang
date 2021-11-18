FROM golang:1.15.6-alpine AS build

WORKDIR /go/src/github.com/kOstovic/CoffeeMachine
COPY . ..

RUN CGO_ENABLED=0 go build -o /bin/coffeemachine cmd/CoffeeMachineRestApiV1/main.go 

FROM scratch
COPY --from=build /bin/coffeemachine /bin/coffeemachine
ENTRYPOINT ["/bin/coffeemachine"]