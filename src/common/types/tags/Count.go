package tags

func (o *Tag)Count() int{
	mutex.Lock()
	defer mutex.Unlock()

	return len(*o)
}