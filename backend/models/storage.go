package models

// Storage represents a storage application for a notes and tag
// tracker.
type Storage interface {
	AddNote(note Note) error
	GetNote(noteId string, userId uint) (Note, error)
	UpdateNote(old, updated Note) error
	DeleteNote(note Note) error

	AddTag(tag Tag) error
	GetTag(tagId string) (Tag, error)
	UpdateTag(old, new Tag) error
	DeleteTag(tag Tag) error

	LinkTagToNote(note Note, tag Tag) (NoteTag, error)
}
