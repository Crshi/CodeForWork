package main

type LRUCache struct {
	Datas map[int]*Node
	Head  *Node
	End   *Node
	Cap   int
}

type Node struct {
	Key   int
	Value int
	Last  *Node
	Next  *Node
}

func Constructor(capacity int) LRUCache {
	data := LRUCache{
		Datas: make(map[int]*Node, capacity),
		Cap:   capacity,
		Head:  &Node{},
		End:   &Node{},
	}

	data.Head.Next = data.End
	data.End.Last = data.Head
	return data
}

func (this *LRUCache) Get(key int) int {
	data, ok := this.Datas[key]

	if ok {
		if this.Head.Next != data {
			this.remove(data)
			this.setHeader(data)
		}
		return data.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	data, ok := this.Datas[key]

	if ok {
		if this.Head.Next != data {
			this.remove(data)
			this.setHeader(data)
		}
		data.Value = value
	} else {
		if len(this.Datas) >= this.Cap {
			deleteNode := this.End.Last
			this.remove(deleteNode)
			delete(this.Datas, deleteNode.Key)
		}

		node := &Node{
			Key:   key,
			Value: value,
		}

		this.setHeader(node)
		this.Datas[key] = node
	}
}

func (this *LRUCache) setHeader(data *Node) {
	this.Head.Next, data.Last, data.Next, this.Head.Next.Last = data, this.Head, this.Head.Next, data
}

func (this *LRUCache) remove(data *Node) {
	data.Last.Next, data.Next.Last = data.Next, data.Last
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // 返回  1
	cache.Put(3, 3) // 该操作会使得密钥 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	cache.Put(4, 4) // 该操作会使得密钥 1 作废
	cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回  3
	cache.Get(4)    // 返回  4
}

// [null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]
// ["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
