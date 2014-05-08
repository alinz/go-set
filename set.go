//The MIT License (MIT)

//Copyright (c) 2014 Ali Najafizadeh

//Permission is hereby granted, free of charge, to any person obtaining a copy of
//this software and associated documentation files (the "Software"), to deal in
//the Software without restriction, including without limitation the rights to
//use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
//the Software, and to permit persons to whom the Software is furnished to do so,
//subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
//FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
//COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
//IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
//CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package set

var existKey = struct{}{}

//Set is a data structure for stroing unqiue items
type Set struct {
	list map[interface{}]struct{}
}

//Has checks whether an item exists in set or not.
//The complexity of calling this method is O(1)
func (s *Set) Has(item interface{}) bool {
	if _, object := s.list[item]; !object {
		return false
	}
	return true
}

//Add inserts new or overwrites existing item into set.
//The complexity of calling this method is O(1)
func (s *Set) Add(item interface{}) {
	s.list[item] = existKey
}

//Remove an item from set. If the value doesn't exists the action doesn't
//do anything.
func (s *Set) Remove(item interface{}) {
	delete(s.list, item)
}

//RemoveAll removes all elements in existing set.
func (s *Set) RemoveAll() {
	s.list = make(map[interface{}]struct{})
}

//Iterate is a functional callback which calls the callback for every item in
//set.
func (s *Set) Iterate(callback func(interface{})) {
	for item := range s.list {
		callback(item)
	}
}

//Size returns the size of set.
func (s *Set) Size() int {
	return len(s.list)
}

//NewNonThreadSafeSet is creating a new set which is NOT thread safe.
func NewNonThreadSafeSet() Set {
	return Set{list: make(map[interface{}]struct{})}
}
