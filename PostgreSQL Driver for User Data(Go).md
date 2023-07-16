# PostgreSQL Driver for User Data

This project provides a Go driver for connecting to a PostgreSQL database and retrieving user data from the "user_table" table. The driver fetches the data and returns the results in JSON string format.

## Development Environment

To set up the development environment for this driver, you need to have the following installed:

- Go programming language (version 1.16 or later): [Download Go](https://golang.org/dl/)
- PostgreSQL database: [Download PostgreSQL](https://www.postgresql.org/download/)

Additionally, you need to install the following Go package for the PostgreSQL driver:


## Configuration

Before using the driver, you need to configure the connection details to your PostgreSQL database. Open the `main.go` file and locate the following line:

```go
db, err := sql.Open("postgres", "user=your_username password=your_password dbname=your_database sslmode=disable")



