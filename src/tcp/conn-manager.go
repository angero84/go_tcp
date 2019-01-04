package tcp

import (
	"sync"
	"errors"
	"fmt"
)

type ConnManager struct {

	conns 			map[uint64]*Conn
	connsMutex 		sync.Mutex
}

func NewConnManager() *ConnManager {

	return &ConnManager{
		conns:		make(map[uint64]*Conn),
	}
}

func (m *ConnManager) addConn( conn *Conn ) ( err error ) {

	m.connsMutex.Lock()
	defer m.connsMutex.Unlock()

	if _, exist := m.conns[conn.id] ; false == exist {
		m.conns[conn.id] = conn
	} else {
		err = errors.New(fmt.Sprintf("the connSeqId %d already exists", conn.id ) )
	}

	return
}

func (m *ConnManager) removeConn( conn *Conn ) ( err error ) {

	m.connsMutex.Lock()
	defer m.connsMutex.Unlock()

	if _, exist := m.conns[conn.id] ; true == exist {
		delete(m.conns, conn.id)
	} else {
		err = errors.New(fmt.Sprintf("the connSeqId %d does not exists", conn.id ) )
	}

	return
}

func (m *ConnManager) findConn( seqId uint64 ) ( conn *Conn ) {

	m.connsMutex.Lock()
	defer m.connsMutex.Unlock()

	conn, _ = m.conns[seqId]

	return
}

func (m *ConnManager) connCount() ( count int ) {
	m.connsMutex.Lock()
	defer m.connsMutex.Unlock()

	count = len(m.conns)

	return
}

