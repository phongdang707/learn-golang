package biz

import (
	"context"
	item "todo/modules/item/model"
)

// Handler -> Biz [ -> Repo ] -> Storage

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *item.TodoItemCreation) error
}

type createNewItemBiz struct {
	store CreateItemStorage
}

func NewCreateNewItemBiz(store CreateItemStorage) *createNewItemBiz {
	return &createNewItemBiz{store: store}
}

func (biz *createNewItemBiz) CreateNewItem(ctx context.Context, data *item.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if errors := biz.store.CreateItem(ctx, data); errors != nil {
		return errors
	}
	return nil
}
