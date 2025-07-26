# 📊 GoFinance: Personal Expense Tracker CLI

**GoFinance** is a command-line tool written in Go that helps you track your income, expenses, and budgets. It uses `Cobra` for CLI commands and `BoltDB` for persistent local storage.

## 🚀 Features

- ✅ Add income or expense entries
- 🗃️ View all saved transactions
- 📤 Export your transactions to CSV
- 📊 Generate summary reports (total income, expenses, balance)
- 💾 Local persistent storage using BoltDB
- ⚡ Fast and efficient CLI experience

## 🧱 Tech Stack

- **Go**: Main programming language
- **Cobra**: CLI toolkit
- **BoltDB**: Embedded key-value database

## 🛠️ Installation

```bash
git clone https://github.com/yourusername/gofinance-cli-tracker.git
cd gofinance-cli-tracker
go build -o gofinance
./gofinance
```

## 📘 Usage

### ➕ Add Transaction

```bash
./gofinance add
```

Fill in the prompts like type (income/expense), amount, category, note.

### 📂 View Transactions

```bash
./gofinance view
```

### 📤 Export to CSV

```bash
./gofinance export
```

Saves `transactions.csv` to your current directory.

### 📊 Summary

```bash
./gofinance summary
```
