# Gunakan official Golang image sebagai base
FROM golang:1.24-alpine

# Set environment agar build Go lebih ringan
ENV CGO_ENABLED=0 GOOS=linux

# Set working directory di dalam container
WORKDIR /app

# Salin dependency files terlebih dahulu agar bisa cache build
COPY go.mod go.sum ./
RUN go mod tidy

# Salin semua source code ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Jalankan aplikasi Go
CMD ["./main"]
