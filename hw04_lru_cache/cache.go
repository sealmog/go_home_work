package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		element.Value.(*cacheItem).value = value
		return true
	}

	if c.queue.Len() > c.capacity {
		c.purge()
	}

	item := &cacheItem{
		key:   key,
		value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.key] = element

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	element, exists := c.items[key]
	if !exists {
		return nil, false
	}
	return element.Value.(*cacheItem).value, true
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
}

func (c *lruCache) purge() {
	if last := c.queue.Back(); last != nil {
		c.queue.Remove(last)
		delete(c.items, last.Value.(*cacheItem).key)
	}
}

/*
2) Реализация кэша на основе ранее написанного списка

Необходимо реализовать следующий интерфейс Cache:

    Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу.
    Get(key Key) (interface{}, bool) // Получить значение из кэша по ключу.
    Clear() // Очистить кэш.

Структура кэша:

    ёмкость (количество сохраняемых в кэше элементов)
    очередь [последних используемых элементов] на основе двусвязного списка
    словарь, отображающий ключ (строка) на элемент очереди

Элемент кэша хранит в себе ключ, по которому он лежит в словаре, и само значение.
Для чего это нужно понятно из алгоритма работы кэша (см. ниже).

Сложность операций Set/Get должна быть O(1), при желании Clear тоже можно сделать О(1).

Алгоритм работы кэша:

    при добавлении элемента:
        если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди;
        если элемента нет в словаре, то добавить в словарь и в начало очереди
          (при этом, если размер очереди больше ёмкости кэша,
          то необходимо удалить последний элемент из очереди и его значение из словаря);
        возвращаемое значение - флаг, присутствовал ли элемент в кэше.
    при получении элемента:
        если элемент присутствует в словаре, то переместить элемент в начало очереди и вернуть его значение и true;
        если элемента нет в словаре, то вернуть nil и false (работа с кешом похожа на работу с map)

Ожидаются следующие тесты:

    на логику выталкивания элементов из-за размера очереди
      (например: n = 3, добавили 4 элемента - 1й из кэша вытолкнулся);
    на логику выталкивания давно используемых элементов
      (например: n = 3, добавили 3 элемента, обратились несколько раз к разным элементам:
      изменили значение, получили значение и пр. - добавили 4й элемент, из первой тройки вытолкнется тот элемент,
      что был затронут наиболее давно).
*/
