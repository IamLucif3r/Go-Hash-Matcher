![](docs/go-hash-matcher-banner.gif)

<div align = "center"> <h1> Go-Hash-Matcher   </h1>	
A simple password cracker written in Go, designed to search for a specific hashed password in a wordlist.
</div>

<br>
<a href=https://twitter.com/IamLucif3r_>
	
<div align = "center">
	
![](https://img.shields.io/twitter/follow/IamLucif3r_?style=social) </a>    <a href=https://github.com/IamLucif3r> ![](https://img.shields.io/github/followers/IamLucif3r?label=Follow%20Me&style=social) </a> </div>

## Features

- Utilizes goroutines for concurrent password checking to improve performance.
- Supports SHA1 hashing with customizable salt.
- Provides flexibility for using different hashing algorithms.

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/IamLucif3r/Go-Hash-Matcher.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Go-Hash-Matcher
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


