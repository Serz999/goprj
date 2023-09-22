package list

import (
	"fmt"
)

type List struct {
	len       int64
	firstNode *Node
}

// NewList создает новый список
func NewList() (l *List) {
	l = &List{0, nil}
	return
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {  
	len = l.len	
	return
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {

	l.len++

	if l.firstNode == nil {	
		id = 0
		l.firstNode = &Node{id, value, nil}
		return
	}

	node := l.firstNode
	for node.next != nil {
		node = node.next	
	}	
	
	id = node.index + 1
	node.next = &Node{id, value, nil}

	return
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if l.Len() == 0 {
		return
	}

	nodeWasRemoved := false

	if id == 0 {
		l.firstNode = l.firstNode.next
		l.firstNode.index--
		nodeWasRemoved = true
	}

	node := l.firstNode.next
	prevNode := l.firstNode

	for node != nil {
		if nodeWasRemoved {
			node.index--	
		} else if node.index == id {
			prevNode.next = node.next	
			nodeWasRemoved = true
			l.len--
		} 
		prevNode = prevNode.next 
		node = node.next
	}
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	node := l.firstNode
	for node != nil {
		if node.value == value {
			l.RemoveByIndex(node.index)
			break
		}
		node = node.next
	}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	
	for node := l.firstNode; node != nil; node = node.next {
		if node.value == value {
			l.RemoveByValue(node.value)
		}	
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	
	for node := l.firstNode; node != nil; node = node.next {
		if node.index == id {
			return node.value, true	
		}	
	}

	return 0, false
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	
	for node := l.firstNode; node != nil; node = node.next {
		if node.value == value {
			return node.index, true	
		}	
	}

	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	ids = nil
	ok = false
	
	for node := l.firstNode; node != nil; node = node.next {
		if node.value == value {
			ids = append(ids, node.index)	
			ok = true
		}	
	}

	return
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	values = nil
	ok = false
	
	for node := l.firstNode; node != nil; node = node.next {
		values = append(values, node.value)	
		ok = true
	}

	return
}

// Clear очищает список
func (l *List) Clear() {
	l.firstNode = nil	
	l.len = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	values, _ := l.GetAll()
	fmt.Printf("%v", values)		
}
