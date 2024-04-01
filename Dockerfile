# Base Image
FROM golang:alpine  

# Work Directory
WORKDIR /app

# Copy source code
COPY . /app

# Build your executable
RUN go build -o sse-client

#EXPOSE 8003

# Define the startup command (if interactive)
CMD ["./sse-client"] 
