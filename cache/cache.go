package cache

import (
	"sync"

	"github.com/boomer-goten/nats-streaming-test/db"
	"github.com/boomer-goten/nats-streaming-test/model"
)

type Cache struct {
	sync.RWMutex
	items map[string]model.Order
}

func NewCache() *Cache {
	items := make(map[string]model.Order)
	cache := Cache{items: items}
	return &cache
}

func (cache *Cache) GetByOrderUID(OrderUID string) (interface{}, bool) {
	cache.RLock()
	defer cache.RUnlock()
	item, found := cache.items[OrderUID]
	if !found {
		return nil, false
	}
	return item, true
}

func (cache *Cache) RestoreFromDB(db *db.DataBaseStorage) error {
	cache.Lock()
	defer cache.Unlock()
	var err error
	cache.items, err = db.GetOrders()
	return err
}

func (cache *Cache) Add(OrderUID string, value model.Order) {
	cache.Lock()
	defer cache.Unlock()
	cache.items[OrderUID] = value
}
