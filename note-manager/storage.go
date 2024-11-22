package main

import (
	"bufio"
	"os"
)

type Storage interface {
	Load() []Note
	Save(notes []Note)
}

type FileStorage struct {
	filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (fs *FileStorage) Load() []Note {
	file, err := os.Open(fs.filename)
	if err != nil {
		return []Note{} // Возвращаем пустой список заметок, если файл не найден
	}
	defer file.Close()

	var notes []Note
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, NewNote(scanner.Text()))
	}

	return notes
}

func (fs *FileStorage) Save(notes []Note) {
	file, err := os.Create(fs.filename)
	if err != nil {
		return // Если возникла ошибка, не сохраняем
	}
	defer file.Close()

	for _, note := range notes {
		file.WriteString(note.Content + "\n")
	}
}
