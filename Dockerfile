FROM golang:1.16-alpine
WORKDIR /app
COPY . .
ARG DEFAULT_PORT 
ENV port $DEFAULT_PORT
ENV DB_USER r
ENV DB_PASSWORD p
EXPOSE $port
RUN go mod vendor
RUN go build -o /user-service
CMD ["/user-service"]