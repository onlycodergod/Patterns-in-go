package pattern

import "fmt"

/*
	Реализовать паттерн "стратегия".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/Strategy_pattern


Пример реализации: кэширование в памяти.
Размер кэша ограничен памятью. При заполнении, некоторые записи нужно убирать для освобождения места через один из алгоритмов (FIFO, LFU, LRU).
Необходимо отделить кэш от алгоритмов для возможности их замены "на ходу". Класс кэша не должен изменяться при добавлении нового алгоритма.
Решение:
Создадим семейства алгоритмов, каждый из которых имеет свой класс.
Все классы применяют одинаковый интерфейс, что делает алгоритмы взаимозаменяемыми внутри этого семейства. Назовем этот общий интерфейс "evictionAlgo".
Основной класс кэша будет включать в себя "evictionAlgo". Вместо прямой реализации всех типов алгоритмов вытеснения внутри самого себя, наш класс будет передавать их в "evictionAlgo"
Поскольку это интерфейс, можно непосредственно во время выполнения программы менять алгоритм на LRU, FIFO, LFU без изменений в классе кэша.
*/

// Интерфейс стратегии
// Определяет интерфейс, общий для всех вариаций алгоритма. Контекст использует этот интерфейс для вызова алгоритма.
// Для контекста не важно, какая именно вариация алгоритма будет выбрана, так как все они имеют одинаковый интерфейс.
type evictionAlgo interface {
	evict(c *cache)
}

// Конкретная стратегия. Реализует вариацию алгоритма FIFO
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by FIFO strategy")
}

// Конкретная стратегия. Реализует вариацию алгоритма LRU
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by LRU strategy")
}

// Конкретная стратегия. Реализует вариацию алгоритма LFU
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by LFU strategy")
}

// Контекст хранит ссылку на объект конкретной стратегии, работая с ним через общий интерфейс стратегий.
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Клиент
// Создаёт объект конкретной стратегии и передаёт его в конструктор контекста.
// Должен иметь возможность заменить стратегию на лету, используя сеттер. Благодаря этому, контекст не будет знать о том, какая именно стратегия сейчас выбрана.
func StrategyPatternRun() {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
}
