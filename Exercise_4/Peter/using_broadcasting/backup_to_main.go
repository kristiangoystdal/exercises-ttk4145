package main

import (
    "fmt"
    "io/ioutil"
    "os"
	"os/exec"
    "strconv"
    "strings"
    "time"
)

func countAndLog(start int) {
    count := start+1
	fmt.Println("resuming from ", count-1)
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
	fmt.Println("--- Backup phase ---")
    for {
        // Get the file's info
        info, err := os.Stat("count.txt")
        if err != nil {
            fmt.Println(err)
            return
        }

        // Check if the file has been modified in the last 2 seconds
        if time.Since(info.ModTime()) > 2*time.Second {
			fmt.Println("... timed out")

			fmt.Println("--- Primary phase ---")

            // Open the file
            data, err := ioutil.ReadFile("count.txt")
            if err != nil {
                fmt.Println(err)
                return
            }

            // Remove the newline character and parse the number in the file
            num, err := strconv.Atoi(strings.TrimSpace(string(data)))
            if err != nil {
                fmt.Println(err)
                return
            }

// Declare cmd before using it
var cmd *exec.Cmd
fmt.Println("... Creating new backup")

// Open a new terminal window and start counting and logging in it
cmd = exec.Command("terminator", "-e", "bash -c '/usr/local/go/bin/go run ~/Shcool/exercises-ttk4145/Exercise_4/Peter/using_files/backup_to_main.go; exec bash'")
err = cmd.Run()
if err != nil {
	fmt.Println("Failed to run command:", err)
}
            go countAndLog(num)
        }
		
        // Sleep for a bit before checking the file again
        time.Sleep(1 * time.Second)
    }
}