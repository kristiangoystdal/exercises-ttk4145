package main

import (
    "fmt"
    "time"
    "os"
    "os/exec"
    "strconv"
)

func countAndLog() {
    count := 0
    for {
        fmt.Println(count)

        // Open the file in write mode, or create it if it doesn't exist
        f, err := os.OpenFile("count.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Write the count to the file
        if _, err := f.WriteString(strconv.Itoa(count) + "\n"); err != nil {
            fmt.Println(err)
            f.Close()
            return
        }

        if err := f.Close(); err != nil {
            fmt.Println(err)
            return
        }

        count++
        time.Sleep(1 * time.Second)
    }
}

func main() {
    fmt.Println("--- Primary phase ---")
    time.Sleep(1 * time.Second)
    // Start counting and logging in the current terminal
    go countAndLog()

    // Open a new terminal window and start counting and logging in it
    cmd := exec.Command("terminator", "-e", "bash -c '/usr/local/go/bin/go run ~/Shcool/exercises-ttk4145/Exercise_4/Peter/backup_to_main.go; exec bash'")
    err := cmd.Run()
    if err != nil {
        fmt.Println("Failed to run command:", err)
    }

    // Prevent the main function from exiting immediately,
    // which would kill the goroutine running countAndLog
    select {}
}