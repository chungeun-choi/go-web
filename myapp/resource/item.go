package resource

import (
	"encoding/json"
	"errors"
	"io"
)

type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var ItemMap map[int]*Item
var LastID int

func init() {
	ItemMap = make(map[int]*Item)
	LastID = 0
}

func GetItem(p int) (*Item, error) {
	value, ok := ItemMap[p]
	if !ok {
		err := errors.New("Item information is not in server")
		return nil, err
	}
	return value, nil
}

func CreateItem(p io.Reader) (int, error) {
	item := new(Item)
	err := json.NewDecoder(p).Decode(item)
	if err != nil {
		return 0, err
	}

	LastID++
	ItemMap[LastID] = item
	return LastID, err
}

func DeleteItem(n int) (int, error) {
	_, ok := ItemMap[n]
	if !ok {
		err := errors.New("Item information is not in server")
		return 0, err
	}

	delete(ItemMap, n)

	return n, nil
}

func UpdateItem(n int, p io.Reader) (*Item, error) {
	_, ok := ItemMap[n]
	if !ok {
		err := errors.New("Item information is not in server")
		return nil, err
	}
	delete(ItemMap, n)

	changeItem := new(Item)
	err := json.NewDecoder(p).Decode(changeItem)
	if err != nil {
		return nil, err
	}

	ItemMap[n] = changeItem
	return changeItem, nil
}
