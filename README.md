# Gin + MySQL

This project is a combination of two tutorials: [one](https://go.dev/doc/tutorial/web-service-gin) and 
[two](https://go.dev/doc/tutorial/database-access).

### Launch
```bash
docker compose up --build
go run main.go
```

### Swagger UI
Swagger UI is located [here](http://localhost:8080/swagger/index.html).

Bearer token is `secret_key`.

To update Swagger docs enter the following line:
```bash
swag init
```
