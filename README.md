# 🎯 FastTrackQuiz - Quiz API & CLI
## Usage

When server and client are up and running, the quiz will start.
Only need to follow the instructions to complete all the answers.

# ⚡️ Normal quick start

### **1️⃣ Build the binary**
```sh
go build -o quiz-cli .
```

### **1️⃣ Run the API Server**
```sh
./quiz-cli start-server
```

### **2️⃣ Run the CLI Client**
```sh
./quiz-cli start-client
```

# ⚡️ Make quick start - Necessary ``make`` installed
```Linux: sudo apt-get install make```

```macOS: brew install make```


### **1️⃣ Run the API Server**
```sh
make run-server
```

### **2️⃣ Run the CLI Client**
```sh
make run-client
```
# Testing
```sh
go test ./...
```

# Testing with ``make``
```sh
make test
```

## Conclusion

I tried to keep this test as simple as possible, but since this is a backend engineering test, I decided to include some things:

```
Deep Logging (slog.Logger) - Helps track API calls and submissions.
Use of Locks (sync.Mutex) - Ensures concurrency safety if scaling concurrent users.
Graceful Shutdown - Handles termination signals properly.
```