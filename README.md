````markdown
# go-readme: AI-Powered Readme and License Generator

## Project Overview

`go-readme` is a command-line application built in Go designed to automate the initial documentation process for projects. It generates an MIT License file and a comprehensive `README.md` by leveraging the Google Gemini API. This tool analyzes the project's directory structure and the content of its files to produce accurate and detailed documentation, significantly reducing manual effort during project setup. It aims to provide a solid foundation for new projects, making it easier for developers to quickly set up and document their repositories.

## Key Features

- **Automated MIT License Generation**: The application generates an `MIT LICENSE` file, dynamically populating it with the current year and the provided author's name. This ensures consistency and proper licensing from the start.
- **AI-Driven README Content Creation**: By integrating with the [Google Gemini API](https://deepmind.google/technologies/gemini/), `go-readme` intelligently analyzes the codebase and directory structure to produce rich, context-aware `README.md` content.
- **Dynamic Project Analysis**: It reads and processes the full contents of all files within the specified project directory, ensuring that the generated README accurately reflects the project's purpose, features, and technical details.
- **Structured Directory Tree Inclusion**: The tool automatically generates a visual representation of the project's directory structure, embedding it directly into the `README.md` to enhance clarity and navigability.

## Tech Stack & Dependencies

- **Go (Standard Library 1.24)**: The entire application is developed using Go, primarily relying on its standard library for file system operations, string manipulation, and command-line argument handling.
- **`github.com/google/generative-ai-go`**: This is the official Go client library for interacting with Google's Generative AI services, including the Gemini models, facilitating the AI-powered content generation.
- **`github.com/joho/godotenv`**: Used for loading environment variables from a `.env` file, allowing for flexible and secure management of API keys and other configurations.
- **`google.golang.org/api`**: A comprehensive Go client for Google APIs, providing underlying support for various Google services.
- **`google.golang.org/genai`**: A more direct and specialized client for the Google Gemini API, enabling fine-grained control over model interactions.

## Environment Variables & Configuration

The application requires a single environment variable to function correctly.

| Variable Name    | Purpose                                                     | Expected Data Type |
| :--------------- | :---------------------------------------------------------- | :----------------- |
| `GEMINI_API_KEY` | Your API key for authenticating with the Google Gemini API. | String             |

## Installation & Setup

1.  **Prerequisites:**
    - Ensure you have Go installed on your system. You can download the latest version from the [official Go website](https://go.dev/dl/).
    - Obtain a `GEMINI_API_KEY` from the [Google AI Studio](https://makersuite.google.com/app/apikey) or Google Cloud Console.

2.  **Clone the Repository:**

    ```bash
    git clone https://github.com/raziel-aleman/go-readme.git
    cd go-readme
    ```

3.  **Install Dependencies:**

    ```bash
    go mod tidy
    ```

4.  **Create `.env` File:**
    In the root directory of the cloned repository, create a file named `.env` and add your Gemini API key:
    ```
    GEMINI_API_KEY=your_actual_gemini_api_key
    ```
    Replace `your_actual_gemini_api_key` with the key you obtained. This file is excluded from version control by `.gitignore`.

## Usage Examples

To generate a `README.md` and `LICENSE` file for a project, navigate to the `go-readme` project's root directory and run the application with two arguments: the path to the target project's root directory and the author's name.

**Example 1: Generating documentation for the current `go-readme` project itself**

```bash
go run . . "Raziel Aleman Ramos"
```
````

This command will create (or overwrite) `LICENSE` and `README.md` files in the current directory (`.`), using "Raziel Aleman Ramos" as the copyright holder.

**Example 2: Generating documentation for another project located at `/path/to/my/new-project`**

```bash
go run . /path/to/my/new-project "Jane Doe"
```

This will create (or overwrite) `LICENSE` and `README.md` files in `/path/to/my/new-project`, attributing the license to "Jane Doe".

## Testing

This project does not include a dedicated unit or integration test suite. The core functionality, which involves interacting with the file system and external API, is typically verified through manual execution and inspection of the generated output. Users are encouraged to run the application as described in the Usage Examples and verify the contents of the generated `LICENSE` and `README.md` files.

## Project Structure

```
.
├── .env
├── .git/
├── .gitignore
├── LICENSE
├── README.md
├── directoryTree.go
├── genai.go
├── go.mod
└── go.sum
```

- **`.env`**: This file stores sensitive configuration data, such as the `GEMINI_API_KEY`, and is loaded at runtime. It is configured to be ignored by Git to prevent accidental exposure.
- **`.git/`**: The standard Git repository directory, containing all version control information and metadata for the project.
- **`.gitignore`**: Specifies files and directories that Git should intentionally ignore, such as `.env` and OS-specific files like `.DS_Store`.
- **`LICENSE`**: The generated MIT License file, which is created by the `main.go` application during its execution.
- **`README.md`**: The primary documentation file for the project, automatically generated by `main.go` using content provided by the Gemini API.
- **`directoryTree.go`**: Contains the Go logic responsible for recursively traversing a given directory path and generating a formatted string representation of its file and folder structure.
- **`genai.go`**: Houses the Go functions that handle the communication with the Google Gemini API, including loading the API key and sending prompts for content generation.
- **`go.mod`**: The Go module definition file, which specifies the project's module path and manages its direct and indirect dependencies.
- **`go.sum`**: A checksum file automatically maintained by Go modules to ensure the integrity and authenticity of downloaded dependencies.

```

```
