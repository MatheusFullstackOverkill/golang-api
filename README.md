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

- Import the collection file: *collection/Golang API Study.postman_collection.json* to to see how to use the endpoints

## How to encrypt the password in the client side

In order to use the login and create user endpoints, you should send the password encrypted in the body, this is one example of how you would go about it:

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
        value = window.btoa(value);
        let encrypted: any = encrypt.encrypt(value);

        return encrypted;
    };
    ```
    In *publicKey* you should use the contents of th *public.pem* file

- Encrypt the password
    ```
    encrypt('password0123');
    ```