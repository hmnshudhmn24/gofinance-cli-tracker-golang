# 📊 GoFinance: Personal Expense Tracker CLI

description: >
  GoFinance is a powerful command-line tool to track income and expenses, manage budgets, and generate reports with ease. Built using Go, Cobra CLI, and BoltDB, it provides a smooth local finance management experience right from your terminal.

badges:
  - "Built with Go 🧬"
  - "Cobra CLI 🐍"
  - "BoltDB Storage 🗄️"
  - "100% Offline Support 🚫🌐"

features:
  - 🚀 Fast and intuitive CLI interface
  - 💵 Add income and expenses with category, amount, and notes
  - 📅 View entries by day, month, or custom range
  - 📊 Generate summaries and reports (total, by category, etc.)
  - 💾 Store all data locally using BoltDB
  - 📤 Export transactions to CSV or JSON
  - 🎯 Budget tracking by category
  - 🧹 Clear data or remove entries by ID

usage:
  - `gofinance add income 5000 --category salary --note "July Salary"`
  - `gofinance add expense 1200 --category groceries --note "Walmart"`
  - `gofinance list --month July`
  - `gofinance report`
  - `gofinance export --format csv`

installation:
  - Install Go from https://golang.org
  - Clone the repository:
    ```bash
    git clone https://github.com/yourusername/gofinance-cli-tracker-golang
    cd gofinance-cli-tracker-golang
    go build -o gofinance
    ./gofinance
    ```

project_structure:
  - `main.go` – Entry point and CLI command routing
  - `db.go` – Handles BoltDB transactions and storage
  - `commands/` – Cobra CLI command definitions (add, list, export, etc.)
  - `models/` – Data models for transactions
  - `exporter/` – CSV and JSON export functions
  - `README.md` – Project documentation

tech_stack:
  - Language: Go
  - CLI Framework: Cobra
  - Storage: BoltDB
  - Output: Terminal + CSV/JSON files

future_plans:
  - 🔒 Password protection or encrypted storage
  - 🌐 Optional sync to cloud/remote
  - 📱 GUI/REST API version
  - 📦 Docker container for portable use

license: MIT

author: Himanshu (GitHub: yourusername)
