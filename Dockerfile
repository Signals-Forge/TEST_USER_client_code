# Base Image
FROM golang:alpine  

# Work Directory
WORKDIR /app

# Copy source code
COPY . /app

# Build your executable
RUN go build -o sse-client


# Define the startup command (if interactive)
CMD ["./sse-client"] 
