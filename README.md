# fasttrack-quiz
The task is to build a super simple quiz with a few questions and a few alternatives for each question. Each with one correct answer. 

Component	Purpose
serve (API)	Runs the quiz backend, exposes HTTP endpoints (/quiz-question, /submit).
start (CLI)	Acts as a client that talks to the API, collects user input, and submits it.
