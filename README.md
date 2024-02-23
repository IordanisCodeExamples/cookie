# Cookie Log Analyzer

This Go application analyzes cookie log files to find the most active cookies for a given date.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Ensure you have Go installed on your machine. This project was built using Go [specify version, e.g., 1.22]. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

### Installing

To start using the Cookie Log Analyzer, clone the repository to your local machine and navigate to the project directory:

```bash
go mod tidy
```

```bash
go mod vendor
```

### Running the Tests

To run the automated tests for this system, use the following command:

```bash
go test ./...
```

This command will recursively run all tests in the project directory, ensuring that each component functions as expected.

### Executing the Script

To execute the script and analyze a cookie log file, run the following command:

```bash
go run cmd/cookie/main.go -f <path to cookie log file> -d <date in YYYY-MM-DD format>
```

For example, to analyze the file `examples/one.csv` for the most active cookies on December 9, 2018, use:

```bash
go run cmd/cookie/main.go -f examples/one.csv -d 2018-12-09
```
