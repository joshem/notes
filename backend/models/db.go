package models

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	db *gorm.DB
}

func New(username, password, name, address, port string) (
	Storage, error) {

	conn := getConnectString(username, name, address, port)
	if len(password) > 0 {
		conn += fmt.Sprintf(" password=%s", password)
	}

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) AddNote(note Note) error {
	return d.db.Create(note).Error
}

func (d *Database) GetNote(noteId string, userId uint) (Note, error) {
	var note Note
	d.db.Where("id = ? AND user_id = ?", noteId, userId).First(&note)

	if note.ID == 0 {
		return Note{}, errors.New(
			fmt.Sprintf("Note %s not found", noteId))
	}

	return note, nil
}

func (d *Database) UpdateNote(old, updated Note) error {
	if updated.Content != "" {
		old.Content = updated.Content
	}

	if updated.Title != "" {
		old.Content = updated.Title
	}

	updated.UpdatedAt = time.Now()

	return d.db.Save(&updated).Error
}

func (d *Database) DeleteNote(note Note) error {
	return d.db.Delete(&note).Error

}
func (d *Database) AddTag(tag Tag) error {
	return d.db.Create(&tag).Error

}

func (d *Database) GetTag(tagId string) (Tag, error) {
	var tag Tag
	d.db.Where("id = ?", tagId).First(&tag)
	if tag.Id == 0 {
		return Tag{}, errors.New(fmt.Sprintf("Tag %s not found", tag))
	}

	return tag, nil
}

func (d *Database) UpdateTag(old, new Tag) error {

	old.Name = new.Name

	return d.db.Save(&old).Error
}

func (d *Database) DeleteTag(tag Tag) error {
	return d.db.Delete(tag).Error
}

func (d *Database) LinkTagToNote(note Note, tag Tag) (NoteTag, error) {
	noteTag := NoteTag{note.ID, tag.Id}
	if d.db.Create(&noteTag).Error != nil {

	}
	return noteTag, nil
}

func getConnectString(username, name, address, port string) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		address, port, username, name)

}
