**Standard REST API GO Server Template**

This repository is a pet project, where a small part of the knowledge that I gained during practice on real projects is collected.

# Some of the Areas This Go Application Touches

- **Architecture**
  - Modular design using [`github.com/labstack/echo/v4`](https://pkg.go.dev/github.com/labstack/echo/v4) for HTTP routing and middleware.

- **Package Structure**
  - Organized into meaningful packages with dependencies managed via `go.mod`.

- **Building the Application**
  - Built with Go 1.22.4, utilizing packages like [`github.com/jackc/pgx/v5`](https://pkg.go.dev/github.com/jackc/pgx/v5) for PostgreSQL interaction and [`github.com/jmoiron/sqlx`](https://pkg.go.dev/github.com/jmoiron/sqlx) for SQL handling.

- **Testing**
  - Testing framework and strategies are not explicitly detailed in `go.mod`; recommended to use Go’s built-in testing support and additional libraries as needed.

- **Configuration**
  - Configuration management via [`gopkg.in/yaml.v2`](https://pkg.go.dev/gopkg.in/yaml.v2) for YAML file parsing.

- **Running the Application**
  - Designed to run with support for Docker and containerized environments (details not included in `go.mod`, but can be integrated with Dockerfiles).

- **Developer Environment/Experience**
  - Development supported by Go modules, with dependencies managed in `go.mod`, ensuring reproducible builds and a consistent development setup.

- **Telemetry**
  - Telemetry and monitoring are not directly included; integration with telemetry solutions would involve additional libraries or middleware as needed.

## Features

- Configuration (using gopkg.in/yaml.v2)
- Logging (using github.com/labstack/gommon)
- Error handling (using github.com/labstack/echo/v4)
- Metrics and tracing (using Prometheus and Jaeger; support via github.com/labstack/echo/v4)
- Health checks (using github.com/labstack/echo/v4/middleware)
- Graceful shutdown (using github.com/labstack/echo/v4)
- Support for multiple server/daemon instances (using github.com/labstack/echo/v4)
- Messaging (using github.com/jackc/pgx/v5)
- PostgreSQL database connection (using github.com/jackc/pgx/v5)
- JWT authentication (using github.com/golang-jwt/jwt)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
