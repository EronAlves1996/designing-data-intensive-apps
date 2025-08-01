package main

import "sync"

type Pair[K any, V any] struct {
	k K
	v V
}

type MapReducer[T any, K comparable, V any, S any] struct {
	PoolNumber  int
	MapperIn    (chan *T)
	MapperOut   (chan *Pair[K, *V])
	Mapper      (func(*T) (K, *V))
	Reducer     (func(*S, K, []*V) *S)
	ReducerBase *S
}

func (m *MapReducer[T, K, V, S]) Accept(v []*T) {
	m.process()
	go func() {
		for _, v := range v {
			m.MapperIn <- v
		}
		close(m.MapperIn)
	}()
}

func (m *MapReducer[T, K, V, S]) process() {
	var wg sync.WaitGroup
	wg.Add(m.PoolNumber)
	for range m.PoolNumber {
		go func() {
			for val := range m.MapperIn {
				k, v := m.Mapper(val)
				p := Pair[K, *V]{
					k, v,
				}
				m.MapperOut <- &p
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(m.MapperOut)
	}()
}

func (m *MapReducer[T, K, V, S]) aggregate() map[K][]*V {
	mapp := make(map[K][]*V)
	for v := range m.MapperOut {
		_, exists := mapp[v.k]
		if !exists {
			mapp[v.k] = []*V{v.v}
		} else {
			mapp[v.k] = append(mapp[v.k], v.v)
		}
	}
	return mapp
}

func (m *MapReducer[T, K, V, S]) Get() *S {
	a := m.aggregate()
	s := m.ReducerBase
	for k, v := range a {
		s = m.Reducer(s, k, v)
	}
	return s
}
