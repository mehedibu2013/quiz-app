
---

## 🧠 Quiz App in Go

A simple command-line quiz app written in Go that reads questions from a CSV file, asks them one-by-one, and tracks user score with an optional time limit.

Perfect for learning **Go concurrency**, **goroutines**, **channels**, **CSV parsing**, and **timers**!

---

## 🚀 Features

- ✅ Read quiz questions from a CSV file
- ✅ Track correct/incorrect answers
- ⏱️ Optional time limit (default: 30 seconds)
- 📝 Gracefully handle malformed input and time expiration
- 🧪 Works on Windows, macOS, and Linux

---

## 📦 Requirements

- [Go](https://go.dev/dl/) installed (1.20+ recommended)
- A terminal or PowerShell
- `problems.csv` file with quiz data

---

## ▶️ How to Run

### Step 1: Clone the repo (optional)

```bash
git clone https://github.com/your-username/quiz-app.git
cd quiz-app
```

### Step 2: Run the program

```bash
go run main.go
```

To customize settings:

```bash
go run main.go -file myquiz.csv -limit 60
```

---

## 🛠️ Flags

| Flag     | Description                      | Default Value |
|----------|----------------------------------|---------------|
| `-file`  | Path to quiz CSV file            | `problems.csv`|
| `-limit` | Time limit in seconds for quiz   | `30`          |

---

## 📄 CSV File Format

Each line should have a question and its answer separated by a comma:

```
5+5,10
7+3,10
What color is the sky?,blue
"Who wrote ""1984""?",George Orwell
```

> ✅ Quoted fields with commas or quotes are supported using Go’s standard CSV parser.

---

## 🎮 Sample Output

```
Press ENTER to start the quiz (you have 30 seconds)...
Question #1: 5+5 = 10
Question #2: 7+3 = 10
Question #3: What color is the sky? = blue
✅ You scored 3 out of 3.
```

If time runs out:

```
⏰ Time is up!
✅ You scored 2 out of 3.
```

---

## 💡 Want to Improve It?

Here are some ideas for new features:

| Feature | Description |
|--------|-------------|
| 🔢 Case-insensitive matching | e.g., accept `"BLUE"` as correct for `"blue"` |
| 🎨 ANSI color output | Green for correct answers, red for incorrect |
| ⏱️ Live countdown display | Show remaining time while quiz runs |
| 📝 Save results to file | Log scores after each quiz |
| 🧠 Shuffle questions | Randomize order every time |


---

## 📌 License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

---

## 👥 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for improvements.

---

## 🙌 Credits

Built as part of a step-by-step tutorial to understand:
- Go concurrency (`goroutines`, `channels`)
- CSV parsing
- Timer-based programs

---
