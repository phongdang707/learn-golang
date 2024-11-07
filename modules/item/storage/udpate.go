package storage

import (
	"context"
	item "todo/modules/item/model"
)

func (s *sqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, data *item.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
