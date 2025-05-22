# Gunakan versi Go 1.23 (Debian-based)
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy go.mod & go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]
