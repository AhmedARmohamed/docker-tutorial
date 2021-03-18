FROM golang:latest

# Add Mainter Info
LABEL maintainer="Ahmed Mohamed <mahmedrashid128@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app


# Copy go mod and sum files
COPY go.mod .

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Coy the source from the current directory to the Working Directory inside the container
COPY . .


# Build the Go app
RUN go build -o main .

# Expose port 8084 to the outside world
EXPOSE 8084

# Command to run the executable
CMD [ "./main" ]