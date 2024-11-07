package biz

import (
	"context"
	"errors"
	item "todo/modules/item/model"
)

// Handler -> Biz [ -> Repo ] -> Storage

type UpdateItemStorage interface {
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *item.TodoItemUpdate) error
	GetItemById(ctx context.Context, cond map[string]interface{}) (*item.Item, error)
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: store}
}

func (biz *updateItemBiz) UpdateItem(ctx context.Context, cond map[string]interface{}, data *item.TodoItemUpdate) error {
	item, err := biz.store.GetItemById(ctx, cond)
	if err != nil {
		return err
	}
	if item.Status == "Deleted" {
		return errors.New("item is deleted")
	}
	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": item.Id}, data); err != nil {
		return err
	}
	return nil
}
