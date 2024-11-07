package storage

import (
	"context"

	item "todo/modules/item/model"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *item.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
