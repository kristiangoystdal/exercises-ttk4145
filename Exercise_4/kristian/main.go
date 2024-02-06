package main

import (
	"fmt"
	"os"
	"os/exec"
	"Exercise_4/pack"
)

func main() {
	// Set an environment variable for the main terminal ID
	os.Setenv("TERMINAL_ID", "1")
	fmt.Println("Running in terminal ID 1")

	// Prepare to start a new terminal with a different terminal ID
    if(os.Getenv("TERMINAL_ID")=="1"){
        cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "set TERMINAL_ID=2 && C:\\Users\\krisg\\OneDrive - NTNU\\2024 Vår\\Sanntidsprogrammering\\exercises-ttk4145\\Exercise_4\\kristian\\main.exe")
        if err := cmd.Run(); err != nil {
            panic(err)
        }
    }
	

	// If pack.Program() needs to wait or you have more code that depends on it, adjust accordingly
	// For demonstration, we directly call pack.Program() after initiating the new terminal to show terminal ID 1 usage
	pack.Program()
}
