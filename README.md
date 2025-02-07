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
- ***create the database and run the contents of the migrations file: migrations/migrations.sql***
- ***make generate_keys***
    to generate the keys for encyption and decryption, mainly for the login password (in a frontend you would use the public key it's for encryption before sending the login request)
- ***go mod tidy***
- ***copy .env.example to .env and set the environment variables***
- ***make run_api***