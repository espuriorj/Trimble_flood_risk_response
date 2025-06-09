# Contributing to Flood Risk Management Service

Thank you for your interest in contributing to the Flood Risk Management Service! We welcome contributions from the community to improve the project. This document outlines the guidelines for contributing.

## How to Contribute
1. **Fork the Repository**:
   - Fork the repository to your own GitHub account.
   - Clone your fork to your local machine:
     ```bash
     git clone <your-fork-url>
     cd floodRiskManagementService
     ```

2. **Create a Branch**:
   - Create a new branch for your feature or bug fix:
     ```bash
     git checkout -b feature/your-feature-name
     ```

3. **Make Changes**:
   - Follow the project's coding style (e.g., Go idiomatic conventions).
   - Ensure your changes align with the project's structure and purpose.
   - Add or update tests as necessary.

4. **Test Your Changes**:
   - Run existing tests to ensure nothing is broken:
     ```bash
     go test ./...
     ```
   - If you add new functionality, include corresponding tests.

5. **Commit Your Changes**:
   - Write clear, concise commit messages following the [Conventional Commits](https://www.conventionalcommits.org/) format, e.g.:
     ```
     feat: add endpoint for flood risk analysis
     fix: resolve database connection timeout issue
     ```
   - Commit your changes:
     ```bash
     git commit -m "feat: your descriptive message"
     ```

6. **Push and Create a Pull Request**:
   - Push your branch to your fork:
     ```bash
     git push origin feature/your-feature-name
     ```
   - Open a pull request (PR) against the `main` branch of the original repository.
   - Provide a clear description of your changes in the PR, including:
     - What the change does
     - Why it is needed
     - Any relevant context or issues it addresses

7. **Code Review**:
   - Respond to feedback from maintainers during the PR review process.
   - Make necessary updates to your branch and push them to keep the PR up-to-date.

## Development Setup
- **Prerequisites**: Ensure you have Go (1.20+), PostgreSQL (13+), `golang-migrate/migrate` CLI, and Make installed.
- **Environment**: Set up a `.env` file as described in [README.md](README.md#environment-variables).
- **Database Migrations**: Run `make migrate` to apply schema changes.
- **Running Locally**: Use `go run main.go` to start the server.

## Code Style
- Follow Go's standard coding conventions (e.g., `gofmt` for formatting).
- Write clear, self-documenting code with appropriate comments.
- Ensure error handling is robust and consistent with the `internal/errors` package.
- Use meaningful variable and function names that reflect their purpose.

## Testing
- Add unit tests for new functionality in the appropriate package (e.g., `internal/handlers`).
- Ensure tests cover both happy paths and edge cases.
- Use the `testing` package and follow Go's testing conventions.

## Reporting Issues
- Use the GitHub Issues page to report bugs or suggest features.
- Provide a clear description, steps to reproduce (for bugs), and any relevant logs or screenshots.

## Community Guidelines
- Be respectful and collaborative in all interactions.
- Follow the project's [Code of Conduct](CODE_OF_CONDUCT.md) (if applicable).
- Avoid submitting changes that are out of scope without prior discussion.

## Getting Help
- For questions or discussions, open an issue or reach out via the project's communication channels (if provided).
- Refer to [README.md](README.md) for setup and usage details.

Thank you for contributing to the Flood Risk Management Service!
