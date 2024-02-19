读取部分字段，使用函数 QueryRowPartialCtx 。

假设有如下一张表：

```sql
CREATE TABLE test (id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, ctime DATETIME);
```

要读取字段 ctime 值。

定义一结构体：

```go
type X struct {
	state int `db:"-"`
	Ctime time.Time `db:"ctime"`
}
```


查询语句：

```
var x X
query := fmt.Sprintf("SELECT ctime FROM %s WHERE id=?", "test")
err := sqlConn.QueryRowPartialCtx(context.Background(), &x, query, id)
if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
				fmt.Printf("no data")
		} else {
				fmt.Printf("query error: %s", err.Error())
		}
} else {
		fmt.Println(x)
}
```

这里需要注意的点：

1、结构 X 中不参与的成员使用 `db:"-"` 修饰，否则报错：

```
sql: Scan error on column index 0, name "ctime": converting driver.Value type time.Time ("2024-02-16 12:05:31 +0800 CST") to a int: invalid syntax
```

2、Ctime 一定要用 `db:"ctime"` 修饰看否则报错：

```
ErrNotMatchDestination = errors.New("not matching destination to scan")
```
