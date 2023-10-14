# Grammar Test Web Application

This web application allows users to take a grammar test by answering questions and provides feedback on their
performance.

## Table of Contents

- [Features](#Features)
- [Installation](#Installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

# Features

- Users can take a grammar test by answering questions presented to them.
- The application provides feedback on the correctness of the answers after submission.
- Statistics are displayed upon completion, including the total number of questions, points scored, and the percentage
  of correct answers.

# Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/visaruttt/english-exercise.git
    ```
2. Install the necessary dependencies using a package manager:
    ```bash
    cd grammar-test-webapp
    go mod download
    ```
3. Run the application:
    ```bash
    go run main.go
    ```

# Usage
1. Access the web application by navigating to http://localhost:8080 in your web browser, where PORT is the port number
specified in the application configuration.
2. Start the grammar test and answer the questions presented to you.
3. After completing the test, review your results, including the total number of questions, points scored, and the
percentage of correct answers.

# Contributing
If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix:
    ```bash
   git checkout -b feature/new-feature
   ```
3. Make your changes and commit them:
    ```bash
   make prepare  # install pre-commit hooks.
   make generate # tidy Go modules.
   git add .
   git commit -m "Add new feature"
    ```
4. Push to the branch:
    ```bash
    git push origin feature/new-feature
    ```
5. Submit a pull request.
