# go-readme: Simple Readme and License Generator

## Description

`go-readme` is a simple Go terminal application for readme and license generation (MIT License) with the [Gemini API](https://deepmind.google/technologies/gemini/). This project serves as a starting point for developers looking for a lightweight CLI application to generate the initial documentation for any project. Future releases will expand on this foundation.

## Technologies and Libraries

*   **`Go (Standard Library)`:** The project primarily uses the Go standard library.
*   **`go-dotenv`:** The project uses the `go-dotenv` library to load environment variables from `.env` file. [go-dotenv](https://github.com/joho/godotenv)

## Project Structure

```
.
├── .env
├── .git
├── .gitignore
├── LICENSE
├── README.md
├── directoryTree.go
├── genai.go
├── go.mod
└── go.sum
```

*   `.env`:  This file contains environment variables used by the application.  (Note: This file *should not* be committed to version control in a production environment; it's included here for demonstration.)
*   `.git`: This directory contains all the version control information, used by git.
*   `.gitignore`: Specifies intentionally untracked files that Git should ignore.
*   `genai.go`: Contains the main application logic.
*   `directoryTree.go`: Shows the project's directory structure.
*   `go.mod`:  Go module file for dependency management.
*   `go.sum`:  Go module checksum file.

## How to Use

1.  **Prerequisites:** Ensure you have Go installed on your system.  You can download it from the official Go website: [https://go.dev/](https://go.dev/)

2.  **Clone the Repository:**
    ```bash
    git clone https://github.com/raziel-aleman/go-readme.git
    cd go-readme
    ```

3.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Create `.env` File:**  Create a `.env` file in the root directory and add your Gemini API key.  For example:
    ```
    GEMINI_API_KEY=your_api_key
    ```

5.  **Run the Application:**
    ```bash
    go run . <path-to-project-root-directory> 'author-name'
    ```
