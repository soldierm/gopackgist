package sync

import "sync"

//多线程安全Map结构体
type SynchronizedMap struct {
	lock *sync.RWMutex
	data map[interface{}]interface{}
}

//加读锁的Get
func (syncMap *SynchronizedMap) Get(key interface{}) (value interface{}) {
	defer syncMap.lock.RUnlock()
	syncMap.lock.RLock()
	return syncMap.data[key]
}

//加读写锁的Set
func (syncMap *SynchronizedMap) Set(key interface{}, value interface{}) {
	defer syncMap.lock.Unlock()
	syncMap.lock.Lock()
	syncMap.data[key] = value
}

//加读写锁的Del
func (syncMap *SynchronizedMap) Del(key interface{}) {
	defer syncMap.lock.Unlock()
	syncMap.lock.Lock()
	delete(syncMap.data, key)
}

//加读写锁的遍历
func (syncMap *SynchronizedMap) Each(call func(key interface{}, value interface{})) {
	defer syncMap.lock.RUnlock()
	syncMap.lock.RLock()
	for key, value := range syncMap.data {
		call(key, value)
	}
}

//生成一个多线程安全Map
func NewSyncMap() (syncMap *SynchronizedMap) {
	return &SynchronizedMap{
		lock: new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}
