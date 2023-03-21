# REST API for managing books

## The project uses skills:
- Development of Web Applications on Go, following the REST API design.
- Working with the framework <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- The Pure Architecture approach in building the application structure. Dependency injection technique.
- Working with the PostgreSQL database.
- Application configuration using the library <a href="https://github.com/spf13/viper">spf13/viper</a>. Working with environment variables.
- Working with the database using the library <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Writing SQL queries.

### Для запуска приложения:

```bash
go run cmd/main.go
```

If the application is launched for the first time, you need to apply migrations to the database:
```bash
migrate -path ./schema -database 'postgres://postgres:password@host:5436/postgres?sslmode=disable' up
```