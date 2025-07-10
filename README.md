# Golang API Study

The purpose of this project was to study main API concepts in Golang, the ones that I focused was:

- CRUD Endpoints
- Database Conection
- Password encyption and decryption
- JWT validation
- File upload and storage
- Clear project structure
- Work with environment variables

## Setup and running steps

### With Docker
- Install Docker and Docker Compose, If you don't already have them installed

- Clone the repository

    ```
    git clone https://github.com/MatheusFullstackOverkill/golang-api
    ```

- Run the docker-compose file with the command:

    ```
    cd golang-api && docker-compose up
    ```

### Without Docker

#### System Requirements

- Golang (tested version: 1.22.5) 
- PostgreSQL (tested version: 12.22)
- openssl (tested version: 3.0.8)
- Enable make command

#### Steps
- Access folder
    ```
    cd golang-api
    ```
- Generate the keys for encyption and decryption of password in requests (optional, I already provided the keys for testing purposes)
    ```
    make generate_keys
    ```

    The private.pem should start with '-----BEGIN RSA PRIVATE KEY-----'. If you have problems generating the keys, you could use a website like https://cryptotools.net/rsagen, setting the key length to 4096, and copy the generated contents in to the private.pem and public.pem files.

- Install the dependencies
    ```
    go mod tidy && go install github.com/pressly/goose/v3/cmd/goose@latest
    ```

- Copy .env.example to .env and set the environment variables

- Run the migrations
    ```
    goose up
    ```

- Run the project
    - With make
        ```
        make run_api
        ```

    - Without make
        ```
        go run .
        ```

- Import the collection file: *collection/Golang API Study.postman_collection.json* to to see how to use the endpoints

## How to encrypt the password in the client side

In order to use the login and create user endpoints, you should send the password encrypted in the body. The Postman collection already has an encrypted key for testing purposes, and in the *client_example* folder there is an example of encryption, but here are more details:

- Generate the keys for encyption and decryption, you can check above how to do it

- Download the lib [JSEncrypt](https://www.npmjs.com/package/jsencrypt)

- Create an encrypt function like this:
    ```
    const encrypt = (value) => {
        const options = {
            default_key_size: "4096",
            default_public_exponent: "010001",
            log: false
        };
        var encrypt = new JSEncrypt(options);
        encrypt.setPublicKey(publicKey);
        let encrypted = encrypt.encrypt(value);

        return encrypted;
    };
    ```
    In *publicKey* you should use the contents of th *public.pem* file

- Encrypt the password
    ```
    encrypt('password0123');
    ```