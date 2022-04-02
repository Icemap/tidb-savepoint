# tidb-savepoint

Replicating the example of TiDB not supporting `savepoint` feature

## Run

Just change to your `MySQL` and `TiDB` dsn at main function

```go
func main() {
	fmt.Printf("\n\nMySQL:\n")
	dbSavepoint("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

	fmt.Printf("\n\nTiDB:\n")
	dbSavepoint("root:@tcp(127.0.0.1:4000)/test")
}
```

And then, run it:

```bash
go build -o bin/tidb-savepoint
./bin/tidb-savepoint
```

## Output

You will get output log like this:

```
MySQL:
id: 1, coins: 1, goods: 1
id: 3, coins: 1, goods: 1


TiDB:

2022/04/02 13:59:48 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.2/mysql.go:397 Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your TiDB version for the right syntax to use line 1 column 9 near "SAVEPOINT sp0x102cf8960" 
[1.119ms] [rows:0] SAVEPOINT sp0x102cf8960

2022/04/02 13:59:48 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.2/mysql.go:397 Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your TiDB version for the right syntax to use line 1 column 9 near "SAVEPOINT sp0x102cf8960" 
[0.001ms] [rows:0] SAVEPOINT sp0x102cf8a00
id: 1, coins: 1, goods: 1
```