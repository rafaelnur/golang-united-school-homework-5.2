package cache

import "time"

type Cache struct {
	Data map[string]Data
}

type Data struct {
	Value    string
	Deadline time.Time
}

func NewCache() Cache {
	d := make(map[string]Data)
	return Cache{Data: d}
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.Data[key]
	if time.Now().After(value.Deadline) {
		return "", false
	}
	return value.Value, ok
}

func (c Cache) Put(key, value string) {
	d := Data{Value: value, Deadline: time.Now().Add(10 * time.Minute)}
	c.Data[key] = d
}

func (c Cache) Keys() []string {
	keys := []string{}
	now := time.Now()
	for key, data := range c.Data {
		if now.Before(data.Deadline) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	d := Data{Value: value, Deadline: deadline}
	c.Data[key] = d
}
