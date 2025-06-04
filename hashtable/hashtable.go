package hashtable

type HashTable struct {
	table []*Entry
	size  int
	count int
}

type Entry struct {
	key   string
	value interface{}
	next  *Entry
}

func New(size int) *HashTable {
	return &HashTable{
		table: make([]*Entry, size),
		size:  size,
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(key[i])
	}
	return (hash & 0x7FFFFFFF) % ht.size
}

func (ht *HashTable) Put(key string, value interface{}) {
	index := ht.hash(key)
	entry := ht.table[index]

	// Check if key already exists
	for entry != nil {
		if entry.key == key {
			entry.value = value
			return
		}
		entry = entry.next
	}

	// Create new entry
	newEntry := &Entry{
		key:   key,
		value: value,
		next:  ht.table[index],
	}
	ht.table[index] = newEntry
	ht.count++
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.hash(key)
	entry := ht.table[index]

	for entry != nil {
		if entry.key == key {
			return entry.value, true
		}
		entry = entry.next
	}

	return nil, false
}

func (ht *HashTable) Delete(key string) bool {
	index := ht.hash(key)
	entry := ht.table[index]

	if entry == nil {
		return false
	}

	if entry.key == key {
		ht.table[index] = entry.next
		ht.count--
		return true
	}

	for entry.next != nil {
		if entry.next.key == key {
			entry.next = entry.next.next
			ht.count--
			return true
		}
		entry = entry.next
	}

	return false
}

func (ht *HashTable) Size() int {
	return ht.count
}
