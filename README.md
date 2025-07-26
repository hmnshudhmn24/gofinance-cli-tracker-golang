
# 📊 GoFinance: Personal Expense Tracker CLI

**GoFinance** is a powerful and lightweight command-line application that allows users to manage personal finances with ease. Designed for simplicity and speed, it helps track income, expenses, generate summaries, and export records—all from your terminal. The tool is ideal for developers and CLI enthusiasts looking for a no-frills budget tracker with local storage.

---

## 🚀 Features

- ✅ Add income or expense entries interactively  
- 🗃️ View all saved transactions in a clean format  
- 📤 Export your data to CSV for external use  
- 📊 Generate summary reports with total income, expenses, and balance  
- 💾 Stores data locally using embedded BoltDB  
- ⚡ Fast and efficient experience with a user-friendly CLI interface

---

## 🧱 Tech Stack

- **Go** – Core programming language  
- **Cobra** – Framework for building CLI applications  
- **BoltDB** – Lightweight embedded key-value database for local persistence

---

## 🛠️ Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/yourusername/gofinance-cli-tracker.git
cd gofinance-cli-tracker
go build -o gofinance
./gofinance
```

---

## 📘 Usage Guide

### ➕ Add Transaction

```bash
./gofinance add
```

- You'll be prompted to enter:
  - Type: `income` or `expense`
  - Amount
  - Category (e.g., food, salary, rent)
  - Optional note

---

### 📂 View Transactions

```bash
./gofinance view
```

- Displays a list of all stored transactions.

---

### 📤 Export Transactions to CSV

```bash
./gofinance export
```

- Saves a `transactions.csv` file in the current directory.

---

### 📊 View Summary Report

```bash
./gofinance summary
```

- Shows:
  - Total income  
  - Total expenses  
  - Current balance

---

## 🧭 Future Enhancements

- 🔒 Add password protection or encryption for transaction data  
- ☁️ Cloud sync or optional integration with Google Drive/Dropbox  
- 📱 Build a TUI (Text-based UI) using libraries like `tview`  
- 📈 Graphical charts via terminal (e.g., income vs expenses trends)  
- 🔄 Monthly recurring transactions  
- 🗓️ Monthly/weekly budgeting and alerts  
- 🌐 REST API version to enable remote access and mobile integrations

---

## 🤝 Contributing

Feel free to fork the repository, suggest features, and open pull requests. Contributions are welcome!

---

## 📎 License

This project is licensed under the [MIT License](LICENSE). Open-source and free to use.
