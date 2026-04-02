package hashmap

import "fmt"

type Entry struct {
	Key   string
	Value int
}

type HashMap struct {
	Bucket [][]Entry
	Size   int
}

func (h *HashMap) HashFunc(key string) int {

	hash := 0
	for _, ch := range key {
		hash += int(ch)
	}
	return hash % h.Size
}

func (h *HashMap) Put(k string, v int) {

	inx := h.HashFunc(k)

	for i, kv := range h.Bucket[inx] {
		if kv.Key == k {
			h.Bucket[inx][i].Value = v
			return
		}
	}

	h.Bucket[inx] = append(h.Bucket[inx], Entry{k, v})

}

func BuildHashMap(size int) *HashMap {

	return &HashMap{
		Bucket: make([][]Entry, size),
		Size:   size,
	}
}

func (h *HashMap) Get(key string) (int, bool) {

	inx := h.HashFunc(key)

	for _, ch := range h.Bucket[inx] {
		if ch.Key == key {
			return ch.Value, true
		}
	}
	return 0, false
}

func (h *HashMap) Remove(k string) {
	inx := h.HashFunc(k)
	bucket := h.Bucket[inx]

	for i, kv := range bucket {
		if kv.Key == k {
			h.Bucket[inx] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func (h *HashMap) Display() {

	for i, bucket := range h.Bucket {
		fmt.Println("Bucket ID: ",i)

		for _,kv:=range bucket{
			fmt.Printf("[%s:%d]",kv.Key,kv.Value)
		}
		fmt.Println()
	}

}

// Main function
func main() {
	hm := BuildHashMap(5)

	hm.Put("apple", 10)
	hm.Put("banana", 20)
	hm.Put("grape", 30)
	hm.Put("apple", 50) // update

	hm.Print()

	val, found := hm.Get("apple")
	fmt.Println("apple:", val, found)

	hm.Remove("banana")
	hm.Print()
}