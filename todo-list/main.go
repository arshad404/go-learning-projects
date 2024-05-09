package main

type todoInterface interface {
	addItem(Item) error
	getItem(int) (Item, error)
	removeItem(int) error
	editItem(Item) error
}

type TodoList struct {
	itemStore *ItemStore
}

func (tl TodoList) addItem(item Item) error {
	return tl.itemStore.Add(item)
}

func (tl TodoList) getItem(itemId int) (Item, error) {
	return tl.itemStore.Get(itemId)
}

func (tl TodoList) removeItem(itemId int) error {
	return tl.itemStore.Remove(itemId)
}

func (tl TodoList) editItem(item Item) error {
	return tl.itemStore.Edit(item)
}

func NewTodoList() *TodoList {
	todoList := &TodoList{
		itemStore: NewItemStore(),
	}
	return todoList
}

func main() {
	todoList := NewTodoList()
	item1 := Item{
		id:   1,
		name: "first item",
	}

	item2 := Item{
		id:   2,
		name: "second item",
	}

	todoList.addItem(item1)
	todoList.addItem(item2)

	item, err := todoList.getItem(1)
	if err == nil {
		println(item.name)
	} else {
		println(err)
	}

	todoList.removeItem(1)

	item, err = todoList.getItem(1)
	if err == nil {
		println(item.name)
	} else {
		println(err.Error())
	}
}
