<p align="center"><img width=100% src="/main/resources/Banner/banner.png"></p>

![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Basic Overview

Backend API based on Golang and MySQL implementing CRUD operations and following the N-Tier Architecture

## Tech Stack

- MySQL Database
- GORM
- JSON Marshalling/Unmarshalling
- Gorilla Mux

## Dependencies

- No dependencies needed.

## How to use

- Clone the repo.
- Setup your username/password through `config/app.go` file then `Connect` function.
- Main.go already built but you can go to cmd/terminal -> path folder for project -> type `go build`
- Using cmd/terminal type `go run main.go` to run the server, server can be tested through Postman HTTP Requests.
- Navigate to:

```text
http://localhost:9010/book/
```

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/MostafaAbdelkarim/Go-Bookstore/issues) to report any bugs or file feature requests.

### Pending Features

- Non for the foreseeable future as this is for educational purposes.
