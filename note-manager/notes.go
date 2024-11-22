package main

type Note struct {
	Content string
}

func NewNote(content string) Note {
	return Note{Content: content}
}
