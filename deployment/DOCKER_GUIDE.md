# Hướng dẫn Cấu hình Docker (Docker Guide)

Tài liệu này giải thích chi tiết cách viết và ý nghĩa của các file Docker trong dự án này, giúp bạn có thể tự cấu hình cho các dự án sau này.

## 1. Dockerfile (Công thức nấu ăn)

`Dockerfile` là file chứa các bước để đóng gói code của bạn thành một **Image** (món ăn đóng hộp).

### A. Dockerfile Frontend (`deployment/Dockerfile.frontend`)

```dockerfile
# 1. Chọn Base Image (Nền móng)
# "node:20-alpine" là bản Linux siêu nhẹ đã cài sẵn Node.js v20
FROM node:20-alpine AS builder

# 2. Tạo thư mục làm việc
WORKDIR /app

# 3. Copy file cài đặt trước (để tận dụng cache của Docker)
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

# 4. Copy toàn bộ code vào
COPY frontend/ .

# 5. Nhận biến môi trường lúc Build (Quan trọng cho Next.js)
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL

# 6. Build ra bản production
RUN npm run build
```

**Nguyên tắc:**
*   **FROM**: Luôn bắt đầu bằng việc chọn hệ điều hành (Node, Golang, Python...).
*   **COPY**: Chép file từ máy tính của bạn vào trong Image.
*   **RUN**: Chạy lệnh (cài thư viện, build code).
*   **ARG/ENV**: Nhận biến cấu hình.

---

### B. Dockerfile Backend (`deployment/Dockerfile.backend`)

```dockerfile
# Giai đoạn 1: Build (Dùng ảnh Golang đầy đủ để compile)
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api  # Build ra file chạy tên là 'main'

# Giai đoạn 2: Run (Dùng ảnh Alpine siêu nhẹ để chạy)
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main . # Chỉ lấy file 'main' từ giai đoạn trước
CMD ["./main"] # Lệnh chạy mặc định
```

**Tại sao cần 2 giai đoạn (Multi-stage build)?**
*   Để giảm dung lượng Image. Giai đoạn 1 cần bộ cài Go (rất nặng) để build. Giai đoạn 2 chỉ cần file chạy (rất nhẹ).
*   Kết quả: Image giảm từ ~1GB xuống còn ~20MB.

---

## 2. Docker Compose (Quản lý nhà hàng)

`docker-compose.yml` là file để quản lý và chạy nhiều Container cùng lúc.

### Cấu trúc cơ bản:

```yaml
version: '3.8'

services: # Danh sách các món ăn (services)

  # --- Món 1: Database (MongoDB) ---
  mongodb:
    image: mongo:6.0 # Dùng ảnh có sẵn trên mạng, không cần build
    restart: always  # Tự sống lại nếu bị lỗi
    volumes:
      - mongodb_data:/data/db # Lưu dữ liệu ra ngoài ổ cứng (để không mất khi tắt)
    networks:
      - app-network # Kết nối vào mạng nội bộ

  # --- Món 2: Backend API ---
  backend-api:
    build:
      context: ../ # Thư mục chứa code
      dockerfile: deployment/Dockerfile.backend # Đường dẫn tới công thức
    ports:
      - "8080:8080" # Mở cổng 8080 ra ngoài
    environment: # Cấu hình biến môi trường
      - MONGODB_URI=mongodb://mongodb:27017 # Gọi Database bằng tên service "mongodb"
    depends_on:
      - mongodb # Chờ Database chạy xong mới được chạy

  # --- Món 3: Frontend ---
  frontend:
    build:
      args: # Truyền biến lúc build (cho cái ARG trong Dockerfile)
        NEXT_PUBLIC_API_URL: https://cryptocheck.click
    ports:
      - "3000:3000"

networks:
  app-network: # Tạo một mạng LAN ảo để các container nhìn thấy nhau

volumes:
  mongodb_data: # Khai báo kho chứa dữ liệu bền vững
```

## Tóm tắt quy trình tự làm:

1.  **Viết Dockerfile**: Cho từng service (Backend, Frontend). Mục tiêu là "Làm sao để code này chạy được trên một máy tính trắng trơn?".
2.  **Viết docker-compose.yml**: Để lắp ghép các service lại.
    *   Khai báo Database (dùng image có sẵn).
    *   Khai báo App của mình (dùng `build`).
    *   Gắn kết chúng bằng `networks` và `environment`.
3.  **Chạy thử**: `docker compose up --build`. 🐳
