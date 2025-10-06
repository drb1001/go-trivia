# Go Trivia

> [!NOTE]
> This project was designed as a learning project:  my first attempt and working with Go (with support from Copilot & ChatGPT).
> There is not much utility in the final program, except for learning the basics of Go. 

This is a simple command-line trivia game written in **Go**, fetching multiple-choice questions from the [Open Trivia Database](https://opentdb.com/).  
Players can answer questions, track consecutive correct streaks, and save high scores with their name. 


## Features

* Fetches multiple-choice trivia questions from Open Trivia DB.
* Uses a CLI interface.
* Tracks consecutive correct answers as a streak.
* Prompts for player name when a new high score is achieved.
* High scores are obfuscated locally with a secret key stored in a config file.
* Input validation: only `a`, `b`, `c`, or `d` accepted for answers.
* Replayable without restarting the program.


## Getting Started

### 1. Prerequisites

* Go 1.20+ installed (or the latest stable version)
* Powershell, Terminal, or any CLI capable of running Go programs

### 2. Clone the repository

```bash
git clone <repo-url>
cd go-trivia
```

### 3. Create the local configuration

The game requires a "**secret key**" for obfuscating the high scores file.   
The secret key is for *obfuscating* the highscore, not for *encryption*, so it is possible to easily reverse-engineer the obfuscation. 
If players really want to set a false highscore, they can (but they can also just google the answers!)

Create a config file with the secret key:

> You should replace `"ExampleSecretKey"` with any string you like. This key is **required** to run the game.


**macOS / Linux:**

```bash
echo '{"secret_key": "ExampleSecretKey"}' > config.json
```

**Windows PowerShell:**

```powershell
'{"secret_key":"ExampleSecretKey"}' | Out-File -Encoding utf8 "config.json"
```


### 4. Run the game

```bash
go run main.go
```


## How to Play

1. The game fetches a multiple-choice question from Open Trivia DB.
2. Each option is labeled **a, b, c, or d**.
3. Type your answer (case-insensitive) and press **Enter**.
4. Your streak of consecutive correct answers will be tracked.
5. If you achieve a **new high score**, you will be prompted to enter your **name**, which is saved with the score.
6. At the end of the game (wrong answer or completion), you can choose to **play again**.

### Example CLI:

```
Welcome to Go Trivia! 🎯
────────────────────────
Current High Score: 1 (by Alice)

Question: What is the capital of France?

  a) Madrid
  b) Paris
  c) Rome
  d) Berlin

Your answer (a/b/c/d): b
✅ Correct!
Current Streak: 1
```


## Development Workflow 

### Running Tests

Go has a built-in testing framework that makes it easy to run both unit tests and integration tests.

> [NOTE!]
> Running tests is safe and read-only — your actual game highscores are not affected by the test suite.

To run all unit tests:

```powershell
go test ./trivia -v
```

Example output:
```powershell

=== RUN   TestPrepare
--- PASS: TestPrepare (0.00s)
=== RUN   TestCheckAnswer
--- PASS: TestCheckAnswer (0.00s)
PASS
ok      go-trivia/trivia 0.012s
```


### Check Test Coverage (Optional)

To see which lines of code are covered by tests:

```powershell
go test -cover ./trivia
```

### File Structure

```
go-trivia/
│
├── main.go               # Entry point
├── go.mod                # go module properties
├── trivia/               # Game logic
│   ├── question.go       # Question struct + API fetching
│   └── highscore.go      # High score handling
│   └── api.go            # API fetching
├── data/                 # High score storage
│   └── highscore.json
├── README.md             # This file
├── config.json           # Obfuscation secret key
└── .gitignore
```


## Notes

* **Highscore obfuscation:** The `config.json` secret key is used to prevent casual editing of high scores. Changing the key will make existing scores unreadable.
* **Input validation:** Only `a`, `b`, `c`, or `d` are valid answers.
* **Replay:** After finishing a round, you can choose to play again without restarting the program.
* **Dependencies:** Only uses Go’s standard library; no third-party modules needed.
