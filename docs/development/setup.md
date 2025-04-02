# Setting Up the Kisumu Development Environment

Welcome to the Kisumu programming language development community! Follow this guide to set up your environment and start contributing.

## **1. Prerequisites**

Ensure you have the following tools installed:

- **Go (1.21 or higher):** [Download Go](https://golang.org/dl/).
- **Git:** [Download Git](https://git-scm.com/).
- **Editor/IDE:** Visual Studio Code (recommended) or any editor of your choice.
- **Make:** Pre-installed on Linux/macOS. Windows users can use [Make for Windows](https://gnuwin32.sourceforge.net/packages/make.htm).
- **golangci-lint:** Install the linter for Go projects:
  ```bash
  $ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

## **2. Clone the Repository**

```bash
# Clone the Kisumu Lang repository
$ git clone https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang.git

# Navigate into the project directory
$ cd kisumu-lang
```

## **3. Initialize the Project**

Kisumu uses Go modules for dependency management. Initialize the project as follows:

```bash
# Verify Go is installed
$ go version

# Initialize the Go module (if not already done)
$ go mod tidy
```

## **4. Install Dependencies**

Install any necessary packages:

```bash
$ make install
```

> **Note:** If `make` is not installed, refer to the "Prerequisites" section to set it up.

## **5. Run Tests**

Ensure everything works by running the unit tests:

```bash
$ make test
```

## **6. Build the Project**

Compile the Kisumu language interpreter:

```bash
$ make build
```

The compiled binary will be placed in the `bin/` directory.

## **7. Run the REPL**

Start the Kisumu Read-Eval-Print Loop (REPL) for interactive experimentation:

```bash
$ ./bin/kisumu-repl
```

## **8. Linting and Formatting**

Kisumu uses `golangci-lint` for static analysis and linting. Run the following commands to ensure code quality:

### **Run Linters**

```bash
$ make lint
```

### **Format Code**

```bash
$ make format
```

> **Note:** The CI/CD pipeline will also run linting and formatting checks. Ensure your code passes these checks before submitting a pull request.

## **9. Contributing**

### **Development Workflow:**

1. **Fork the Repository:** Create your own copy of the project.
2. **Create a Feature Branch:**
   ```bash
   $ git checkout -b feature/awesome-feature
   ```
3. **Write Code and Tests:** Ensure your changes are thoroughly tested.
4. **Run Lint and Formatting Checks:**
   ```bash
   $ make pre-commit
   ```
5. **Submit a Pull Request:** Push your changes and open a PR for review.

### **Commit Guidelines:**

- Use descriptive commit messages, e.g., `Feat(literals): Add support for integer literals.`
- Follow the [contribution guidelines](../development/contribution-guidelines.md).

## **10. CI/CD Pipeline**

Kisumu uses GitHub Actions for continuous integration and deployment. The pipeline performs the following tasks:

- **Build:** Compiles the project on multiple platforms (Linux, macOS, Windows).
- **Test:** Runs unit tests and generates a coverage report.
- **Lint:** Checks code quality using `golangci-lint`.
- **Format:** Ensures code is properly formatted.

You can view the pipeline status on the [Actions tab](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/actions) in the repository.

## **11. Need Help?**

If you encounter any issues:

- Check the [FAQ](docs/faq.md) (coming soon).
- Open a GitHub [issue](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/issues).
- Join our community discussions on [Discord](https://discord.gg/amrst3npC8).

## **12. Additional Resources**

- [Go Documentation](https://golang.org/doc/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [golangci-lint Documentation](https://golangci-lint.run/)
