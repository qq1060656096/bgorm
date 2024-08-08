package bmdb_test

import (
	"fmt"
	"github.com/qq1060656096/bgorm/bmdb"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestOk(t *testing.T) {
	assert.True(t, true)
}

func TestDefaultDbManager(t *testing.T) {
	bmdb.DefaultDbManager.Register("test", "test.sign", &gorm.DB{})
	bmdb.DefaultDbManager.Register("test2", "test2.sign", &gorm.DB{})
	bmdb.DefaultDbManager.Register("test3", "test3.sign", &gorm.DB{})
	assert.LessOrEqual(t, 3, bmdb.DefaultDbManager.Count())
}

func TestNewMemoryDbManager(t *testing.T) {
	m := bmdb.NewMemoryDbManager()
	test1Db := &gorm.DB{}
	test2Db := &gorm.DB{}
	test3Db := &gorm.DB{}

	m.Register("test1", "test1.sign", test1Db)
	m.Register("test2", "test2.sign", test2Db)
	m.Register("test3", "test3.sign", test3Db)
	db, err := m.Get("test1")
	assert.Nil(t, err)
	assert.Equal(t, test1Db, db)

	exist := m.Exists("test1")
	assert.True(t, exist)

	exist = m.Exists("test3")
	assert.True(t, exist)

	exist = m.Exists("test.not.exist")
	assert.False(t, exist)

	assert.Equal(t, 3, m.Count())

	m.Unregister("test2")
	assert.Equal(t, 2, m.Count())
	exist = m.Exists("test2")
	assert.False(t, exist)

	sign, err := m.GetSign("test3")
	assert.Nil(t, err)
	assert.Equal(t, "test3.sign", sign)
}

func BenchmarkConcurrence(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrentTestDefaultDbManager()
		}
	})
}

func concurrentTestDefaultDbManager() {
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("test%d", i)
		sign := fmt.Sprintf("test%d.sign", i)
		db := &gorm.DB{}
		bmdb.DefaultDbManager.Register(name, sign, db)
		bmdb.DefaultDbManager.Get(name)
		bmdb.DefaultDbManager.Count()
		bmdb.DefaultDbManager.Unregister(name)
	}
}
