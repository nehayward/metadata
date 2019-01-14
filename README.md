
# Metadata

This is a simple API to manage application metadata.

![Version](https://img.shields.io/badge/version-0.1-brightgreen.svg)

## Installation

Clone the project and `go run main.go`.

## Endpoints

|URL | HTTP Method | Functionality | Example
|:---:|:---:|:---|:---|
|/upload | POST | Upload Application YAML metadata | curl -X POST -D @example.yaml -H "Content-type: text/x-yaml" "http://localhost:8080/upload"
|/search| GET | Search uploaded application titles, glob search | curl -X GET "http://localhost:8080/search?title=app*"

### License

This work is published under the MIT license.

Please see the [LICENSE](https://github.com/nehayward/metadata/blob/master/LICENSE) file for details.
