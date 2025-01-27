# Gunakan image Golang yang sudah ada sebagai base image
FROM golang:1.23-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum dan install dependensi
COPY go.mod go.sum ./
RUN go mod tidy

# Salin kode sumber aplikasi ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Expose port yang akan digunakan oleh aplikasi
EXPOSE 3000

# Jalankan aplikasi setelah container berjalan
CMD ["./main"]
