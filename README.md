# Bookstore-API-GO-MySQL

This project is a simple RESTful API for managing a bookstore, built with `Go`, `Gorilla Mux` for routing, and `GORM` for
database interaction.

## Features

- Create Book: Add a new book to the bookstore.
- Retrieve Books: Fetch all books or a specific book by its ID.
- Update Book: Modify details of an existing book.
- Delete Book: Remove a book from the bookstore.

## Project Structure

* `main.go`: Entry point for the application.
* `routes`: Defines API routes using Gorilla Mux.
* `controllers`: Contains business logic for handling requests.
* `models`: Defines the `Book` structure and database operations.
* `utils`: Utility functions like request body parsing.
* `config`: Database connection setup.

## Endpoints

| Method | Endpoint       | Description               |
|--------|----------------|---------------------------|
| POST   | /book          | Create a new book         |
| GET    | /book          | Retrieve all books        |
| GET    | /book/{bookId} | Retrieve a book by its ID |
| PUT    | /book/{bookId} | Update a book by its ID   |
| DELETE | /book/{bookId} | Delete a book by its ID   |

### Dependencies

* `Gorilla Mux` - Routing library.
* `GORM` - ORM for database interaction.
