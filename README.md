# Go Project Documentation

This repository contains the Go-based backend for our project. It is designed to handle various backend functionalities including API endpoints, data processing, and integration with other components of the system.

## Project Structure

- `/backend/cmd`: Main applications for this project.
  - `/backend/cmd/main.go`: The HTTP server application.
  - `/backend/configs`: Configuration related files.
    - `/backend/configs/config.yml`: Main configuration file.
- `/backend/internal`: Internal application and library code.
  - `/backend/internal/config`: Configuration related code.
  - `/backend/internal/handler`: HTTP handlers.
  - `/backend/internal/lib`: Shared libraries.
  - `/backend/internal/models`: Data models.
  - `/backend/internal/repository`: Database interaction logic.
  - `/backend/internal/server`: Server setup and initialization.
  - `/backend/internal/service`: Business logic.
- `/backend/migrations`: Database migration files.
- `/backend/scripts`: Various scripts for setup and maintenance.
- `/backend/static`: Static files served by the application.
- `.env.example`: Example environment configuration file.
- `.gitignore`: Git ignore rules.
- `Dockerfile`: Docker configuration for the application.
- `README.md`: This documentation file.
- `docker-compose.yml`: Docker Compose configuration file.
- `go.mod`: Module dependencies.
- `go.sum`: Checksums for module contents.
- `run_migrations.sh`: Script to run database migrations.

## Technologies Used

- **Go**: The primary programming language used.
- **Libraries**:
  - `github.com/gin-gonic/gin`: HTTP web framework.
  - `github.com/go-redis/redis/v8`: Redis client for Go.
  - `github.com/joho/godotenv`: Load environment variables from `.env` files.
  - `github.com/jmoiron/sqlx`: Extensions to database/sql.
  - `github.com/lib/pq`: PostgreSQL driver for Go.

## Integration with Other Components

This repository is part of a larger ecosystem with the following components:

- **Frontend**: [Frontend Repository](https://github.com/your-organization/frontend)
- **Database**: Utilizes PostgreSQL for data storage.
- **Internal Services + ML**: [Internal Backend Repository](https://github.com/hackathon5-7/internal-backend-ml/tree/dev)
