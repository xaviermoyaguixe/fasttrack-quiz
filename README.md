# FastTrackQuiz - Quiz API & CLI

The task is to build a super simple quiz with a few questions and a few alternatives for each question. Each with one correct answer. 

> **Note:**  
> If you use Docker Compose, you don't need Go installed on your machine.  
> For a direct installation, you'll need Go (v1.23.4 or later) installed.

## REST API Endpoints

The API is deliberately kept simple yet functional, consisting of just two main endpoints:

### GET `/api/v1/quiz/questions`
- **Purpose:** Retrieves a list of quiz questions along with the available answers.
- **Use Case:** When a user accesses the quiz, they first call this endpoint to fetch all the questions.
- **Response:** Typically returns a JSON object containing an array of question objects, each with its possible answers.

### POST `/api/v1/quiz/submit`
- **Purpose:** Accepts the user's answers to the quiz questions.
- **Use Case:** After the user completes the quiz, they submit their answers to this endpoint.
- **Response:** Returns a JSON object containing the quiz result, including the number of correct answers and a comparative metric (for example, "better than 60% of all quizzers").

---

## Prerequisites

### For Docker-based Installation (Recommended)
- **Docker** and **Docker Compose** installed on your machine.

### For Normal (Direct) Installation
- **Go** (v1.23.4 or later) installed on your machine.

---

## Quick Start - Make sure to go step by step

### Using Docker Compose

This method provides a consistent, containerized environment—**no Go installation is required.**

1. **Build & Start the Server**

   Run the following command in the project root:

   ```sh
   docker-compose up --build
   ```
2. **In a new terminal runs the CLI**

   Run the following command in the project root:

   ```sh
   docker run -it --rm --network=fasttrack-quiz_quiz-network -e QUIZ_API_URL=http://server:3000 quiz-cli start-client
   ```
3. **Stop services & Docker**

   Run the following command in the project root:

   ```sh
   docker-compose down
   ```
### Normal Installation (Without Docker)

This method provides a consistent, containerized environment—**no Go installation is required.**

1. **Build the Binary**

   Run the following command in the project root:

   ```sh
   go build -o quiz-cli .
   ```
2. **Build & Start the Services**

   Run the following command in the project root:

   ```sh
   ./quiz-cli start-server
   ```
3. **In a new terminal run the CLI**

   Run the following command in the project root:

   ```sh
   ./quiz-cli start-client
   ```
---
### Testing

1. **Locally test**

   Run the following command in the project root:

   ```sh
   go test ./...
   ```
