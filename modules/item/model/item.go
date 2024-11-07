package item

import (
	"errors"
	"strings"
	"time"
	"todo/common"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
)

type Item struct {
	common.SQLModel
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status;default:News"`
}

func (Item) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status;default:News"`
}

func (TodoItemCreation) TableName() string {
	return Item{}.TableName()
}

func (item *TodoItemCreation) Validate() error {
	item.Title = strings.TrimSpace(item.Title)
	if item.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string     `json:"status" gorm:"column:status;default:News"`
}

func (TodoItemUpdate) TableName() string {
	return Item{}.TableName()
}
type TodoItemDelete struct {
	Id int `json:"id" gorm:"column:id"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Status      *string     `json:"status" gorm:"column:status;default:News"`
}

func (TodoItemDelete) TableName() string {
	return Item{}.TableName()
}
