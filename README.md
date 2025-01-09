# GenAI: A Simple Generative AI Image App

This project is a basic demonstration of generating images using a generative AI model (details below). It's designed to be a simple, understandable example for developers familiar with Go and web development basics.  This is the first release, and we're eager to get your feedback!


## Interesting Techniques

* **Go Web Development:** The project uses Go's `net/http` package to create a simple web server.  [Learn more about `net/http` on MDN](https://developer.mozilla.org/en-US/docs/Learn/Server-side/First_steps/Setup) (Note: MDN doesn't have dedicated Go docs, this links to general server-side concepts).
* **Environment Variables:** Configuration is managed using environment variables loaded from a `.env` file ([`.env`](./.env)). This allows for easy changes without modifying the code.
* **Go Modules:** The project uses Go modules for dependency management.  The `go.mod` and `go.sum` files manage project dependencies.


## Non-Obvious Technologies/Libraries

This project currently doesn't use any complex external libraries beyond the standard Go library.  Future releases may include integrations with specific image generation APIs.


## Project Structure

```
├── go.mod
├── go.sum
├── .env
├── LICENSE
├── README.md
├── genai.go
├── directoryTree.go
├── main.go
```

* **`genai.go`**: Contains the core logic for interacting with the generative AI model (currently a placeholder).  This will be expanded in future releases.
* **`directoryTree.go`**: (If implemented)  This file would contain code to dynamically generate the directory structure (similar to what is done here in the README).
* **`.env`**: Stores environment variables (API keys, etc.).  Keep this file out of version control!


## How to Use

1. **Clone the repository:** `git clone <repository_url>`
2. **Install Go:** Make sure you have Go installed on your system. [Download Go](https://go.dev/dl/)
3. **Install dependencies:** `go mod tidy`
4. **Set environment variables:** Create a `.env` file and populate it with any necessary API keys or configuration settings (if required in future releases).  The structure of the `.env` file will be documented in future releases if needed.
5. **Run the application:** `go run main.go`


## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

[MIT License](./LICENSE)

