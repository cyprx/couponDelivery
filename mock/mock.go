package mock

import (
	"log"
)

type MockCouponIterator struct {
	data    []string
	current int
}

func NewMockCouponIterator() *MockCouponIterator {
	data := []string{"testVal1", "testVale2"}
	return &MockCouponIterator{data: data, current: -1}
}

func (m *MockCouponIterator) Next() bool {
	m.current += 1
	if m.current < len(m.data) {
		return true
	}
	return false
}

func (m *MockCouponIterator) Scan(dest ...interface{}) error {
	for _, d := range dest {
		d = m.data[m.current]
		log.Println(d)
	}
	return nil
}
