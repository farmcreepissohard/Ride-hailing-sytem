```
my-gin-project/
├── cmd/
│   └── api/
│   └── main.go # Điểm bắt đầu (entry point) của ứng dụng
├── internal/ # Code nội bộ, các project khác không thể import được
│   ├── config/ # Load và lưu trữ cấu hình (đọc từ .env, yaml)
│   ├── models/ # Định nghĩa các struct (Entities, DTOs)
│   ├── repositories/ # Giao tiếp với Database (CRUD operations)
│   ├── services/ # Chứa Business Logic cốt lõi
│   ├── controllers/ # Nhận Request từ Gin, gọi Service, trả về Response
│   ├── middlewares/ # Các logic đứng giữa Request & Controller (Auth, Logging)
│   └── routers/ # Định nghĩa các endpoints và gắn Controller/Middleware vào
├── pkg/ # Code dùng chung có thể chia sẻ cho các project khác (utils, helpers)
├── migrations/ # Các file script để tạo và cập nhật schema Database
├── .env # Biến môi trường (không push lên git)
├── go.mod # Quản lý dependencies
└── go.sum # Checksum của các dependencies
```

go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
