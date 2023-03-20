# Go REST API using Gin
=====================

This is a simple Go program that sets up a REST API using the Gin framework. The API allows clients to perform CRUD (Create, Read, Update, Delete) operations on a collection of "event" records.

## Getting started

### Prerequisites

To run this program, you need to have Go installed on your system. You can download and install Go from the [official website](https://golang.org/dl/).

### Installing

To install this program, clone this repository to your local machine:

### Running the program

Navigate to the project directory and run the following command to start the server:

```go run main.go```

By default, the server will listen on `localhost:8080`.

### Usage

#### GET `/events/show`

Returns a list of all events as JSON.

#### GET `/events/show/:ID`

Locates the event whose `ID` value matches the `id` parameter sent by the client, then returns that event as a response.

#### POST `/events/addJ`

Adds an event from JSON received in the request body.

Example JSON body:

```{"id": 2, "x": 2, "y": 2}```

#### POST `/events/addQ`

Adds an event from query received in URL.

Example URL:

#### PUT `/events/update`

Updates an event with the given `ID`.

Example JSON body:

```{"id": 1, "x": 10, "y": 20}```

#### DELETE `/events/delete`

Deletes an event with the given `ID`.

Example JSON body:

```{"id": 1}```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) - Web framework used in this project.
 
---

© 2023 Your Name. All Rights Reserved.
