# Note Manager

This is a console application on Go for managing notes. You can add, edit, delete and display your notes. Notes are saved in a text file so that they are saved between app launches.

## Installation

1. Make sure that Go is installed on your computer. [Install Go](https://golang.org/doc/install ).
2. Clone the repository or create a `note-manager` folder and add the files `main.go`, `notes.go`, and `storage.go'.
3. In the terminal, go to the `note-manager` folder.

```bash
cd path/to/note-manager
```
## Launch

Launch the application using the following command:

```bash
go run .\main.go .\storage.go .\notes.go
```
## Usage

Select the desired menu item to manage your notes:

1. Add a note
2. Edit the note
3. Delete the note
4. Display a list of notes
5. Exit
