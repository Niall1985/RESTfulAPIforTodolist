# Todo List RESTful API

This API allows you to manage a todo list. You can add, delete, and toggle the status of todo items.

## Usage

### Prerequisites

- Go installed on your system
- Postman or curl installed for API testing

### Running the Application

1. Clone this repository to your local machine.

2. Navigate to the project directory in your terminal.

3. Run the following command to start the server:

    ```sh
    go run main.go
    ```

4. The server will start running on `localhost:8000`.

### Endpoints

#### 1. Get Todo List

- **URL:** `GET /todo_list`
- **Description:** Retrieve the list of todo items.
- **Example:**
    ```sh
    curl http://localhost:8000/todo_list
    ```

#### 2. Add Todo Item

- **URL:** `POST /todo_list`
- **Description:** Add a new todo item to the list.
- **Example:**
    ```sh
    curl -X POST http://localhost:8000/todo_list -d '{"id":"4","task":"New Task","status":false}' -H "Content-Type: application/json"
    ```

#### 3. Delete Todo Item

- **URL:** `DELETE /todo_list/:id`
- **Description:** Delete a todo item by its ID.
- **Example:**
    ```sh
    curl -X DELETE http://localhost:8000/todo_list/1
    ```

#### 4. Toggle Todo Item Status

- **URL:** `PUT /todo_list/:id/toggle`
- **Description:** Toggle the status of a todo item (from false to true or vice versa) by its ID.
- **Example:**
    ```sh
    curl -X PUT http://localhost:8000/todo_list/1/toggle
    ```

### PowerShell Commands

Here are equivalent PowerShell commands for each endpoint:

```powershell
# Get Todo List
Invoke-RestMethod -Uri http://localhost:8000/todo_list -Method Get

# Add Todo Item
Invoke-RestMethod -Uri http://localhost:8000/todo_list -Method Post -ContentType 'application/json' -Body '{"id":"4","task":"New Task","status":false}'

# Delete Todo Item
Invoke-RestMethod -Uri http://localhost:8000/todo_list/1 -Method Delete

# Toggle Todo Item Status
Invoke-RestMethod -Uri http://localhost:8000/todo_list/1/toggle -Method Put

```

## License
This project has been licensed under the MIT lisence

