# Go Project Documentation

This repository contains the Go-based backend for our project. It is designed to handle various backend functionalities including API endpoints, data processing, and integration with other components of the system.

## Project Structure

- `/cmd`: Main applications for this project.
  - `/cmd/api`: The HTTP server application.
- `/internal`: Internal application and library code.
  - `/internal/config`: Configuration related code.
  - `/internal/handlers`: HTTP handlers.
  - `/internal/middleware`: Middleware components.
  - `/internal/models`: Data models.
  - `/internal/repository`: Database interaction logic.
  - `/internal/services`: Business logic.
- `/pkg`: External packages that can be used by other applications.
- `go.mod`: Module dependencies.
- `go.sum`: Checksums for module contents.

## Technologies Used

- **Go**: The primary programming language used.
- **Libraries**:
  - `gorilla/mux`: HTTP request router and dispatcher.
  - `jinzhu/gorm`: ORM library for Go.
  - `spf13/viper`: Configuration management.
  - `sirupsen/logrus`: Structured logger for Go.

## Integration with Other Components

This repository is part of a larger ecosystem with the following components:

- **Frontend**: [Frontend Repository](https://github.com/your-organization/frontend)
- **Database**: Utilizes PostgreSQL for data storage.
- **Internal Services + ML**: [Internal Backend Repository](https://github.com/hackathon5-7/internal-backend-ml/tree/dev)
