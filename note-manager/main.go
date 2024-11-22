package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	storage := NewFileStorage("notes.txt")
	notes := storage.Load()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nNote Manager")
		fmt.Println("1. Add Note")
		fmt.Println("2. Edit Note")
		fmt.Println("3. Delete Note")
		fmt.Println("4. List Notes")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter note content: ")
			scanner.Scan()
			content := scanner.Text()
			notes = append(notes, NewNote(content))
			storage.Save(notes)
			fmt.Println("Note added.")
		case "2":
			fmt.Print("Enter note ID to edit: ")
			var id int
			fmt.Scan(&id)
			if id < 1 || id > len(notes) {
				fmt.Println("Invalid note ID.")
				continue
			}
			fmt.Print("Enter new content: ")
			scanner.Scan()
			content := scanner.Text()
			notes[id-1].Content = content
			storage.Save(notes)
			fmt.Println("Note updated.")
		case "3":
			fmt.Print("Enter note ID to delete: ")
			var id int
			fmt.Scan(&id)
			if id < 1 || id > len(notes) {
				fmt.Println("Invalid note ID.")
				continue
			}
			notes = append(notes[:id-1], notes[id:]...)
			storage.Save(notes)
			fmt.Println("Note deleted.")
		case "4":
			fmt.Println("Notes:")
			for i, note := range notes {
				fmt.Printf("ID: %d, Content: %s\n", i+1, note.Content)
			}
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
