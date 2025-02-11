# FastTrackQuiz - Quiz API & CLI

The task is to build a super simple quiz with a few questions and a few alternatives for each question. Each with one correct answer. 

> **Note:**  
> If you use Docker Compose, you don't need Go installed on your machine.  
> For a direct installation, you'll need Go (v1.23.4 or later) installed.

---

## Prerequisites

### For Docker-based Installation (Recommended)
- **Docker** and **Docker Compose** installed on your machine.

### For Normal (Direct) Installation
- **Go** (v1.23.4 or later) installed on your machine.

---

## Quick Start

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
