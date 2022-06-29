package internal

type Queue struct {
	items []string
}

func NewQueue(l int) *Queue {
	return &Queue{
		items: make([]string, 0, l),
	}
}

func (q *Queue) Enqueue(s string) {
	q.items = append(q.items, s)
}

func (q *Queue) Dequeue() string {
	if len(q.items) == 0 {
		return ""
	}

	deq := q.items[0]
	q.items = q.items[1:]
	return deq
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) GetData() []string {
	return q.items
}
