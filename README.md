# Go-Hash-Matcher

A simple password cracker written in Go, designed to search for a specific hashed password in a wordlist.

## Features

- Utilizes goroutines for concurrent password checking to improve performance.
- Supports SHA1 hashing with customizable salt.
- Provides flexibility for using different hashing algorithms.

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/IamLucif3r/go-password-cracker.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-password-cracker
    ```

3. Build the project:

    ```bash
    go build main.go
    ```

4. Run the executable:

    ```bash
    ./main
    ```

    Replace `main` with the name of the compiled executable if it's different.

5. The program will read passwords from the `passwords.txt` file and check for a specific hashed password.

## Configuration

- Adjust the `hashType` variable to use a different hashing algorithm (e.g., SHA256).
- Modify the `salt` variable to change the salt used during hashing.

## Dependencies

The project uses only standard Go libraries and does not require additional dependencies.

## Contributing

Feel free to contribute by opening issues, providing feedback, or submitting pull requests.


