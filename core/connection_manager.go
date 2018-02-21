package core

import (
	"sync"
)

// ConnectionManager ...
type ConnectionManager struct {
	conns map[*Connection]bool
	mutex sync.Mutex
}

func (cm *ConnectionManager) add(c *Connection) {
	cm.mutex.Lock()
	cm.conns[c] = true
	cm.mutex.Unlock()
}

func (cm *ConnectionManager) remove(c *Connection) {
	cm.mutex.Lock()
	delete(cm.conns, c)
	cm.mutex.Unlock()
}

func (cm *ConnectionManager) count() int {
	return len(cm.conns)
}
