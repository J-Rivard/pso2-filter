FROM golang

# Set the Current Working Directory inside the container
WORKDIR /app/pso2-filter

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build ./cmd/pso2-filter/main.go

EXPOSE 8080

CMD [ "./main" ]