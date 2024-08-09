# bgorm
business gorm

```shell
# 测试
go test -count=2 -v -race ./...
```

```go
package main

import (
    "crypto/md5"
    "encoding/json"
    "fmt"
    "github.com/qq1060656096/bgorm"
)

func main() {
    mdb := bgorm.DefaultDbManager
    dsn := "root:root@tcp(127.0.0.1:3306)/test_data_1?charset=utf8mb4&parseTime=True&loc=Local"
    dsn2 := "root:root@tcp(127.0.0.1:3306)/test_data_2?charset=utf8mb4&parseTime=True&loc=Local"
    // database 1
    db := bgorm.MustDbOpen(bgorm.DriverTypeMysql, dsn)
    dbSign := fmt.Sprintf("%s", md5.Sum([]byte(dsn)))
    mdb.Register("test_data_1", dbSign, db)
    // database 2
    db = bgorm.MustDbOpen(bgorm.DriverTypeMysql, dsn2)
    dbSign = fmt.Sprintf("%s", md5.Sum([]byte(dsn2)))
    mdb.Register("test_data_2", dbSign, db)
    
    db, err := mdb.Get("test_data_1")
    if err != nil {
    panic(err)
    }
    tools := []map[string]interface{}{}
    db.Raw("select tools_id,tenant_id from tools limit 2").Find(&tools)
    jsonBytes, err := json.Marshal(tools)
    if err != nil {
    panic(err)
    }
    fmt.Println(string(jsonBytes))
    // Output: [{"tenant_id":1,"tools_id":1},{"tenant_id":1,"tools_id":3}]
}
```