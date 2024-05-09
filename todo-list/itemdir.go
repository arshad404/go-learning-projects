package main

import "errors"

type Item struct {
	id   int
	name string
}

type ItemStore struct {
	items map[int]Item
}

type ItemI interface {
	add(Item) error
	remove(int) error
	edit(Item) error
	get(int) (Item, error)
}

func NewItemStore() *ItemStore {
	return &ItemStore{
		items: make(map[int]Item),
	}
}

func (is *ItemStore) Add(item Item) error {
	if _, exists := is.items[item.id]; exists {
		return errors.New("item already exists")
	}
	is.items[item.id] = item
	return nil
}

func (is *ItemStore) Remove(id int) error {
	if _, exists := is.items[id]; !exists {
		return errors.New("item not found")
	}
	delete(is.items, id)
	return nil
}

func (is *ItemStore) Edit(item Item) error {
	if _, exists := is.items[item.id]; !exists {
		return errors.New("item not found")
	}
	is.items[item.id] = item
	return nil
}

func (is *ItemStore) Get(id int) (Item, error) {
	item, exists := is.items[id]
	if !exists {
		return Item{}, errors.New("item not found")
	}
	return item, nil
}
