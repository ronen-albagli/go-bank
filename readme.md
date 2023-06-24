# Go Clean Bank
####  This is a sample application demonstrating the use of Clean Architecture principles in a GoLang project. The project structure follows the recommended guidelines for organizing code and separating concerns.

### Project Structure
The project structure is organized as follows:

```go
app/                 # Application layer
    adapter/         # Adapters for external dependencies (e.g., MongoDB, SQS)
    entities/        # Entities representing the core business logic
    types/           # Project-specific types and models

config/              # Application configuration

ports/               # Inbound layer (interfaces for interacting with the application)

usagecases/          # Use cases or handlers for executing application operations

webserver/           # Web server layer (using the Gin framework)
```

### Installation
Clone the repository: ```go git clone <repository-url>```

Install dependencies: ```go go mod download```

### Configuration
The application's configuration can be found in the config/ directory. Update the necessary configuration files to match your environment settings, such as database connection details, API keys, and other application-specific settings.

### Building and Running
To build and run the application, use the following command:

```go
go run main.go

```

## Testing

To run the tests, execute the following command:

```go
go test ./...
```

### Usage
The application provides a set of use cases or handlers that can be invoked to interact with the core business logic. The usagecases/ directory contains these handlers, which utilize the configured gateways and repositories to perform operations such as inserting or retrieving data from the database.

To interact with the application, you can make HTTP requests to the web server endpoints defined in the webserver/ directory. Refer to the API documentation or code comments for more details on the available endpoints and their expected request/response formats.

### Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

### License
This project is licensed under the MIT License.

Feel free to further customize this README file according to your specific project requirements.