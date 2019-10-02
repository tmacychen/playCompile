package main

type List struct {
	cells []interface{}
	pos   int
}

func NewList() *List {
	return &List{}
}

func (l *List) Add(a interface{}) {
	l.cells = append(l.cells, a)
}
func (l *List) Get(p int) interface{} {
	if p >= len(l.cells) {
		return nil
	}
	return l.cells[p]
}

func (l *List) Read() interface{} {
	pos := l.pos
	if pos < l.Len() {
		l.pos++
	}
	return l.cells[pos]
}

func (l *List) Unread() interface{} {
	if l.pos > 0 {
		l.pos--
	}
	return l.cells[l.pos]
}
func (l *List) GetCells() []interface{} {
	return l.cells
}
func (l *List) Len() int {
	return len(l.cells)
}
func (l *List) GetPos() int {
	return l.pos
}
