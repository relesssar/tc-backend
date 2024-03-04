package tc

import "database/sql"
import "encoding/json"

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}

/*
//Заметка, только для списков
type NoteForList struct {
	Id         string `json:"id" db:"id"`
	CategoryId string `json:"category_id" db:"category_id"`
	//UserId     string `json:"user_id" db:"user_id"`
	UserId string `json:"_" db:"user_id"`
	Name   string `json:"name" db:"name"`
	//Text       string `json:"text" db:"text"`
	Text string `json:"_" db:"text"`
}

//Заметка, вся информация
type Note struct {
	Id         string `json:"id" db:"id"`
	CategoryId string `json:"category_id" db:"category_id"`
	UserId     string `json:"user_id" db:"user_id"`
	Name       string `json:"name" db:"name"`
	Text       string `json:"text" db:"text"`
}

//Создание заметки
type InsertNoteInput struct {
	CategoryId string `json:"category_id" db:"category_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
	Name       string `json:"name" db:"name" binding:"required" example:"Новая Заметка"`
	Text       string `json:"text" db:"text"`
}

//Обновление Заметки
type UpdateNoteInput struct {
	CategoryId *string `json:"category_id"`
	Name       *string `json:"name"`
	Text       *string `json:"text`
}

func (i UpdateNoteInput) Validate() error {
	if i.CategoryId == nil && i.Name == nil && i.Text == nil {
		return errors.New("Все значения пустые")
	}
	return nil
}*/
