# Goland API Study

The pourpose of this project was to study main API concepts in Golang, the ones that I focused was:

- CRUD Endpoints
- Database Conection
- Password encyption and decryption
- JWT validation
- File upload and storage
- Clear project structure
- Work with environment variables

## System Requirements
- Goland (tested version: 1.22.5) 
- PostgreSQL (tested version: 12.22)

## Setup and running steps
- Create the database and run the contents of the migrations file: *migrations/migrations.sql*

- Generate the keys for encyption and decryption
    ```
    make generate_keys
    ```
    It's mainly for the login password (in a frontend you would use the public key for it's encryption before sending the login request)

- Install the dependencies
    ```
    go mod tidy
    ```
- Copy .env.example to .env and set the environment variables

- Run the project
    ```
    make run_api
    ```