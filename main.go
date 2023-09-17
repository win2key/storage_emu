package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleCommand(text string, disks *[]string) {
	// Handle commands
	switch text {
	case "test":
		fmt.Println("Test command")
	case "list disks":
		fmt.Println("Disks: ", *disks)
	case "add disk":
		*disks = append(*disks, fmt.Sprintf("Disk%d", len(*disks)+1))
	case "remove disk":
		if len(*disks) > 0 {
			*disks = (*disks)[:len(*disks)-1]
		}
	case "exit":
		os.Exit(0)
	case "help":
		cmd_help()
	default:
		fmt.Println("Command not recognized.")
	}
}

func main() {
	// fmt.Println("Arguments received:", os.Args) // Print arguments
	// Initialize "virtual" storage state
	var disks = []string{"Disk1", "Disk2"}

	var history []string
	var historyIndex = -1

	if len(os.Args) > 1 {
		// Detecting "-c" flag
		if os.Args[1] == "-c" && len(os.Args) > 2 {
			joinedArgs := strings.Join(os.Args[2:], " ")
			handleCommand(joinedArgs, &disks)
		} else {
			handleCommand(os.Args[1], &disks)
		}
	} else {
		// Interactive mode
		reader := bufio.NewReader(os.Stdin)
		for {
			// Show a prompt
			fmt.Print("FS> ")

			text := ""
			for {
				char, _, err := reader.ReadRune()
				if err != nil {
					fmt.Println("Read error:", err)
					return
				}

				if char == '\r' || char == '\n' {
					break
				} else if char == '\x1b' { // escape byte
					// _, _ = reader.ReadRune() // consume '['
					reader.ReadRune() // consume '['
					arrow, _, _ := reader.ReadRune()
					if arrow == 'A' { // up arrow
						if historyIndex > 0 {
							historyIndex--
						}
						if historyIndex >= 0 && historyIndex < len(history) {
							text = history[historyIndex]
						}
						break
					} else if arrow == 'B' { // down arrow
						if historyIndex < len(history)-1 {
							historyIndex++
						}
						if historyIndex >= 0 && historyIndex < len(history) {
							text = history[historyIndex]
						}
						break
					}
				} else {
					text += string(char)
				}
			}

			// Trim leading and trailing whitespaces
			text = strings.TrimSpace(text)

			if text != "" {
				// Add updated or new command to the end of the history
				history = append(history, text)
				historyIndex = len(history) - 1
			}

			if text == "" {
				// Skip empty lines
				continue
			}

			handleCommand(text, &disks)
		}
	}
}
