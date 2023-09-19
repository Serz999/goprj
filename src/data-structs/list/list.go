package list

type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() (l *List) {
	return
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	return
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	return
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	return
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	return
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	return
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	return
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	return
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	return
}

// Clear очищает список
func (l *List) Clear() {

}

// Print выводит список в консоль
func (l *List) Print() {

}
