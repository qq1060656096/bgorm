package bgorm

import (
	"encoding/json"
	"sync"

	"gorm.io/gorm"
)

var (
	_                DbManager = (*MemoryDbManager)(nil)
	DefaultDbManager DbManager = NewMemoryDbManager()
)

// NewMemoryDbManager 创建一个新的 MemoryDbManager 实例
func NewMemoryDbManager() DbManager {
	return &MemoryDbManager{
		dbs:     make(map[string]*gorm.DB),
		dbSigns: make(map[string]string),
	}
}

// MemoryDbManager 管理多个数据库实例
type MemoryDbManager struct {
	dbs     map[string]*gorm.DB
	dbSigns map[string]string
	mutex   sync.RWMutex
}

// Register 注册数据库实例
func (m *MemoryDbManager) Register(name, sign string, db *gorm.DB) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.dbs[name] = db
	m.dbSigns[name] = sign
}

// Unregister 注销数据库实例
func (m *MemoryDbManager) Unregister(name string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.exists(name) {
		delete(m.dbs, name)
		delete(m.dbSigns, name)
		return true
	}
	return false
}

// Get 获取数据库实例
func (m *MemoryDbManager) Get(name string) (*gorm.DB, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if db, ok := m.dbs[name]; ok {
		return db, nil
	}
	return nil, errDbNotFound
}

// GetSign 获取数据库实例标识
func (m *MemoryDbManager) GetSign(name string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if sign, ok := m.dbSigns[name]; ok {
		return sign, nil
	}
	return "", errDbNotFound
}

// Exists 检查数据库实例是否存在
func (m *MemoryDbManager) Exists(name string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.exists(name)
}

// Count 获取数据库实例数量
func (m *MemoryDbManager) Count() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.dbs)
}

// String 返回所有数据库实例及其标识的JSON字符串
func (m *MemoryDbManager) String() string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	snapshot := make(map[string]string)
	for name := range m.dbs {
		snapshot[name] = m.dbSigns[name]
	}

	b, err := json.Marshal(snapshot)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// exists 是一个内部方法，用于检查数据库实例是否存在
func (m *MemoryDbManager) exists(name string) bool {
	_, ok := m.dbs[name]
	return ok
}
