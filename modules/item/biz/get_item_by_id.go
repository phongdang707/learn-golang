package biz

import (
	"context"
	item "todo/modules/item/model"
)

type GetItemStorage interface {
	GetItemById(ctx context.Context, cond map[string]interface{}) (*item.Item, error)
}

type getItemByIdBiz struct {
	store GetItemStorage
}

func NewGetItemByIdBiz(store GetItemStorage) *getItemByIdBiz {
	return &getItemByIdBiz{store: store}
}

func (b *getItemByIdBiz) GetItemById(ctx context.Context, id int) (*item.Item, error) {
	item, err := b.store.GetItemById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return item, nil
}

