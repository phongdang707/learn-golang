package storage

import (
	"context"
	item "todo/modules/item/model"
)

func (s *sqlStorage) GetItemById(ctx context.Context, cond map[string]interface{}) (*item.Item, error) {
	var result item.Item
	if err := s.db.Where(cond).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
