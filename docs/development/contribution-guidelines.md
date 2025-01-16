# Contribution Guidelines

Thank you for considering contributing to Kisumu Lang! Your involvement is vital to shaping the future of this project. Whether you’re fixing a bug, suggesting an enhancement, or implementing a new feature, we’re thrilled to have your input.

## Getting Started

1. **Read the Documentation:** Familiarize yourself with the [Kisumu Lang Specification](../specs/introduction.md) and the project roadmap to understand the language's goals and design principles.

2. **Check Existing Issues:** Look at the [issue tracker](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/issues) to see if your idea or bug report is already being discussed. If not, create a new issue.

3. **Fork the Repository:** Fork the [Kisumu Lang repository](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang.git) to your account and clone it locally.

   ```bash
   git clone https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/issues
   cd kisumu-lang
   ```

4. **Set Up Your Environment:**
   - Ensure you have Go 1.21 or higher installed.
   - Install any required dependencies:

     ```bash
     go mod tidy
     ```

5. **Follow the Code Style:** Adhere to the existing code style and structure. Use `gofmt` to format your code before submitting.

## How to Contribute

### Reporting Issues

1. **Search First:** Avoid duplicates by searching the [issue tracker](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/issues) before reporting.
2. **Create a Detailed Report:** Include the following:
   - A clear and concise description of the issue.
   - Steps to reproduce the issue.
   - Expected vs. actual behavior.
   - Screenshots or logs, if applicable.

### Suggesting Features

1. **Review the Roadmap:** Ensure your suggestion aligns with the project’s goals.
2. **Open an Issue:** Use the "Feature Request" template to outline your idea.
3. **Provide Justification:** Explain why the feature is beneficial and how it fits into the overall vision of Kisumu Lang.

### Making Code Contributions

1. **Create a Branch:** Use a descriptive name for your branch.

   ```bash
   git checkout -b feature/my-feature
   ```

2. **Write Clear Commit Messages:** Explain what you’ve done and why. For example:

   ```
   Feat(datatype): Add support for custom data types
   ```

3. **Test Your Changes:** Ensure all tests pass, and write new tests for your additions:

   ```bash
   go test ./...
   ```

4. **Submit a Pull Request (PR):**
   - Reference related issues in the PR description.
   - Describe your changes and why they’re necessary.
   - Mark your PR as a draft if it’s not ready for review.

5. **Engage in the Review Process:**
   - Be open to feedback and iterate on your PR based on reviewers’ suggestions.
   - Respond promptly to comments.

## Community Guidelines

- **Be Respectful:** Maintain a welcoming and inclusive environment.
- **Ask for Help:** If you’re stuck, don’t hesitate to ask questions in the [Discussions](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang) tab.
- **Stay Focused:** Ensure your contributions align with the project’s goals.

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](../../LICENSE).

## Contact Us

If you have any questions or need clarification, feel free to reach out via:
- [GitHub Discussions](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/discussions)
- [Discord](https://discord.gg/amrst3npC8)

