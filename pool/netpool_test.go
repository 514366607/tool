package pool

import (
	"net"
	"testing"

	"github.com/514366607/tools/logger"
)

func init() {
	logger.SetLevel(10)
}

var (
	testPort      = 80
	testPoolCount = 10
)

func TestGet(t *testing.T) {
	pool, err := NewNetPool("192.168.0.10", testPort, testPoolCount)
	if err != nil {
		t.Error(err)
		return
	}
	defer pool.Closes()

	conn := pool.Get()
	pool.Put(conn)

	if pool.FreeCount() != 10 {
		t.Error("数据不对")
		return
	}

}

func TestBusy(t *testing.T) {
	pool, err := NewNetPool("192.168.0.10", testPort, testPoolCount)
	if err != nil {
		t.Error(err)
		return
	}
	defer pool.Closes()

	var testpool = make(chan net.Conn, testPoolCount+1)
	for i := 1; i <= testPoolCount+1; i++ {
		testpool <- pool.Get()
	}
	if len(testpool) != 11 {
		t.Error("取出数据不对", len(testpool))
		return
	}

	close(testpool)
	for c := range testpool {
		pool.Put(c)
	}

	if pool.FreeCount() != 10 {
		t.Error("数据不对", pool.FreeCount())
		return
	}
}
