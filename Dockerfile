#Step 1: Use the official Go image to build the binary
FROM golang:1.23.2-alpine AS build

#Step 2: Set the working directory inside the container
WORKDIR /app

#Copy the go.mod and go.sum files first to leverage Docker cache
# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

#Step 3: Copy your Go project files to the container
COPY . .

#Step 4: Download dependencies
RUN go mod tidy

#Step 5: Build the Go binary
RUN go build -o myapp .

#Step 6: Use a lightweight image to run the binary
FROM golang:1.23.2-alpine

#Step 7: Set the working directory in the runtime container
WORKDIR /app

#Step 8: Copy the binary from the build stage to this stage
COPY --from=build /app/myapp .


#Step 9: Expose the application port (if needed)
#EXPOSE 8080

#Step 10: Run the Go application
CMD ["tail", "-f", "/dev/null"]
