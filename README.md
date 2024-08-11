# Book Management API

> This is my first go application!

Welcome to the Book Management API! This API provides a simple way to manage a collection of books. With endpoints for viewing, creating, checking in, and checking out books, it allows you to perform basic operations on a book collection.

## Endpoints

### 1. **Get All Books**

- **Method:** `GET`
- **Endpoint:** `/books`
- **Description:** Retrieves a list of all books in the collection.
- **Response:** A JSON array of book objects with `ID`, `Title`, `Author`, and `Quantity` fields.

### 2. **Get Book by ID**

- **Method:** `GET`
- **Endpoint:** `/books/:id`
- **Description:** Retrieves a single book by its unique ID.
- **URL Parameter:**
  - `id` (string): The ID of the book to retrieve.
- **Response:**
  - **Success:** A JSON object representing the book with fields `ID`, `Title`, `Author`, and `Quantity`.
  - **Error:** A JSON object with an error message if the book is not found.

### 3. **Create a New Book**

- **Method:** `POST`
- **Endpoint:** `/books`
- **Description:** Adds a new book to the collection.
- **Request Body:** A JSON object with the fields `ID`, `Title`, `Author`, and `Quantity`.
- **Response:**
  - **Success:** A JSON object representing the newly created book.
  - **Error:** A JSON object with an error message if the book cannot be created.

### 4. **Check In a Book**

- **Method:** `POST`
- **Endpoint:** `/checkin`
- **Description:** Increases the quantity of a book by 1 (i.e., checking in a book).
- **Query Parameter:**
  - `id` (string): The ID of the book to check in.
- **Response:**
  - **Success:** A JSON object representing the updated book with increased quantity.
  - **Error:** 
    - If the `id` parameter is missing, a JSON object with an error message.
    - If the book is not found, a JSON object with an error message.
    - If the book is not available (quantity <= 0), a JSON object with an error message.

### 5. **Check Out a Book**

- **Method:** `POST`
- **Endpoint:** `/checkout`
- **Description:** Decreases the quantity of a book by 1 (i.e., checking out a book).
- **Query Parameter:**
  - `id` (string): The ID of the book to check out.
- **Response:**
  - **Success:** A JSON object representing the updated book with decreased quantity.
  - **Error:**
    - If the `id` parameter is missing, a JSON object with an error message.
    - If the book is not found, a JSON object with an error message.

## Example Usage

### Get All Books

```bash
curl -X GET http://localhost:8080/books
```

### Get Book by ID

```bash
curl -X GET http://localhost:8080/books/1
```

### Create a New Book

```bash
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"ID": "11", "Title": "New Book", "Author": "Author Name", "Quantity": 5}'
```

### Check In a Book

```bash
curl -X POST http://localhost:8080/checkin?id=1
```

### Check Out a Book

```bash
curl -X POST http://localhost:8080/checkout?id=1
```

## Running the Server

To start the server, run the following command in your terminal:

```bash
go run main.go
```

The server will start on `localhost:8080`. Make sure you have the [Gin](https://github.com/gin-gonic/gin) package installed in your Go environment.

Enjoy managing your book collection with this API! 📚