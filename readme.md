# Book API in Go

This is a Book API created using Go. It provides various endpoints for managing books.

## Endpoints

- **Get All Books**: Retrieve a list of all books.

  - **URL**: `localhost:8080/books`
  - **HTTP Method**: GET

- **Get Book by ID**: Retrieve a book by its unique identifier.

  - **URL**: `localhost:8080/books/<id>`
  - **HTTP Method**: GET

- **Create a New Book**: Add a new book to the database.

  - **URL**: `localhost:8080/book`
  - **HTTP Method**: POST

- **Update Book by ID**: Modify an existing book by its unique identifier.

  - **URL**: `localhost:8080/books/<id>`
  - **HTTP Method**: PUT

- **Delete a Book by ID**: Remove a book from the database by its unique identifier.

  - **URL**: `localhost:8080/book/<id>`
  - **HTTP Method**: DELETE

## Usage

Here's how you can use these endpoints:

1. **Get All Books**

   ```bash
   curl -X GET localhost:8080/books
   ```

2. **Get Book by ID**

   ```bash
   curl -X GET localhost:8080/books/<id>
   ```

3. **Create a New Book**

   ```bash
   curl -X POST -H "Content-Type: application/json" -d <book data here> localhost:8080/book
   ```

4. **Update Book by ID**

   ```bash
   curl -X PUT -H "Content-Type: application/json" -d <updated book data here> localhost:8080/books/<id>

   ```

5. **Delete a Book by ID**
   ```bash
   curl -X DELETE localhost:8080/book/<id>
   ```
