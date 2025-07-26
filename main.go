package main

import (
    "encoding/json"
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
    "time"

    "github.com/boltdb/bolt"
    "github.com/spf13/cobra"
)

type Transaction struct {
    Type     string  `json:"type"`
    Amount   float64 `json:"amount"`
    Category string  `json:"category"`
    Note     string  `json:"note"`
    Date     string  `json:"date"`
}

var db *bolt.DB

func main() {
    var rootCmd = &cobra.Command{Use: "gofinance"}
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(viewCmd)
    rootCmd.AddCommand(exportCmd)
    rootCmd.AddCommand(summaryCmd)

    var err error
    db, err = bolt.Open("gofinance.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("Transactions"))
        return err
    })

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a transaction",
    Run: func(cmd *cobra.Command, args []string) {
        t := Transaction{}

        fmt.Print("Type (income/expense): ")
        fmt.Scanln(&t.Type)

        fmt.Print("Amount: ")
        fmt.Scanln(&t.Amount)

        fmt.Print("Category: ")
        fmt.Scanln(&t.Category)

        fmt.Print("Note: ")
        fmt.Scanln(&t.Note)

        t.Date = time.Now().Format("2006-01-02 15:04:05")

        db.Update(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte("Transactions"))
            id, _ := b.NextSequence()
            key := strconv.Itoa(int(id))

            encoded, err := json.Marshal(t)
            if err != nil {
                return err
            }

            return b.Put([]byte(key), encoded)
        })

        fmt.Println("Transaction added successfully.")
    },
}

var viewCmd = &cobra.Command{
    Use:   "view",
    Short: "View all transactions",
    Run: func(cmd *cobra.Command, args []string) {
        db.View(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte("Transactions"))
            return b.ForEach(func(k, v []byte) error {
                var t Transaction
                json.Unmarshal(v, &t)
                fmt.Printf("ID: %s | Type: %s | ₹%.2f | %s | %s | %s
",
                    k, t.Type, t.Amount, t.Category, t.Note, t.Date)
                return nil
            })
        })
    },
}

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Export transactions to CSV",
    Run: func(cmd *cobra.Command, args []string) {
        file, err := os.Create("transactions.csv")
        if err != nil {
            log.Fatal("Failed to create CSV:", err)
        }
        defer file.Close()

        writer := csv.NewWriter(file)
        defer writer.Flush()

        writer.Write([]string{"ID", "Type", "Amount", "Category", "Note", "Date"})

        db.View(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte("Transactions"))
            return b.ForEach(func(k, v []byte) error {
                var t Transaction
                json.Unmarshal(v, &t)
                record := []string{
                    string(k),
                    t.Type,
                    fmt.Sprintf("%.2f", t.Amount),
                    t.Category,
                    t.Note,
                    t.Date,
                }
                return writer.Write(record)
            })
        })

        fmt.Println("Exported to transactions.csv")
    },
}

var summaryCmd = &cobra.Command{
    Use:   "summary",
    Short: "Show a summary of total income and expenses",
    Run: func(cmd *cobra.Command, args []string) {
        var income, expense float64
        db.View(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte("Transactions"))
            return b.ForEach(func(k, v []byte) error {
                var t Transaction
                json.Unmarshal(v, &t)
                if t.Type == "income" {
                    income += t.Amount
                } else if t.Type == "expense" {
                    expense += t.Amount
                }
                return nil
            })
        })

        fmt.Printf("💰 Total Income: ₹%.2f
", income)
        fmt.Printf("💸 Total Expenses: ₹%.2f
", expense)
        fmt.Printf("📊 Balance: ₹%.2f
", income-expense)
    },
}