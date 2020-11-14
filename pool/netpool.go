package pool

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/514366607/tools/logger"
)

// NetConnectPool 网络连接池
type NetConnectPool struct {
	ip        string
	port      int
	poolCount int
	pool      chan net.Conn
	sync.RWMutex
}

// NewNetPool 初始化网络连接池
// Get 完后记得Put回去！
func NewNetPool(ip string, port, poolCount int) (*NetConnectPool, error) {
	var newPool = &NetConnectPool{
		ip:        ip,
		port:      port,
		poolCount: poolCount,
		pool:      make(chan net.Conn, poolCount),
	}

	for i := 1; i <= poolCount; i++ {
		conn, err := newPool.connect()
		if err != nil {
			return nil, err
		}
		newPool.pool <- conn
	}

	return newPool, nil
}

func (p *NetConnectPool) connect() (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", p.ip, p.port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Get 取得连接
func (p *NetConnectPool) Get() (conn net.Conn) {
	p.Lock()
	defer p.Unlock()

	select {
	case conn = <-p.pool:
		return
	case <-time.After(time.Second):
		logger.Get().Sugar().Info("连接池繁忙，生成一个新连接")
		var err error
		conn, err = p.connect()
		if err != nil {
			logger.Get().Sugar().Error("连接失败", err)
			return p.Get()
		}
		return
	}
}

// Put 放回连接池
func (p *NetConnectPool) Put(conn net.Conn) {
	p.Lock()
	defer p.Unlock()

	if len(p.pool) < p.poolCount {
		p.pool <- conn
	} else {
		// 关闭新生成的连接
		err := conn.Close()
		if err != nil {
			logger.Get().Sugar().Error(err)
		}
	}
	return
}

// FreeCount 空闲连接数
func (p *NetConnectPool) FreeCount() int {
	p.RLock()
	defer p.RUnlock()
	return len(p.pool)
}

// Closes 关闭所有连接
func (p *NetConnectPool) Closes() {
	p.Lock()
	defer p.Unlock()

	close(p.pool)
	for conn := range p.pool {
		err := conn.Close()
		if err != nil {
			logger.Get().Sugar().Error(err)
		}
	}
	return
}
