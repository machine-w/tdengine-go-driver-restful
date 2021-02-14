package taossqlrestful

import (
	"database/sql"
	"fmt"
	"log"
	taossqlrestful "taossqlrestful/taosSqlRestful"
	"testing"
)

func TestConn(t *testing.T) {
	conn := new(taossqlrestful.Conn)
	conn.TaosConnect("121.36.56.117", "root", "msl110918", "dianli", 6041)
	fmt.Printf(conn.Taos)
}

func TestDb(t *testing.T) {
	db, err := sql.Open("taossqlrestful", "mydb://dalong@127.0.0.1/demoapp")
	if err != nil {
		t.Errorf("some error %s", err.Error())
	}
	rows, err := db.Query("select name,age,version from demoapp")
	if err != nil {
		log.Fatal("some wrong for query", err.Error())
	}
	for rows.Next() {
		var user taossqlrestful.MyUser
		if err := rows.Scan(&user.Name, &user.Age, &user.Version); err != nil {
			log.Println("scan value erro", err.Error())
		} else {
			log.Println(user)
		}
	}
}

func Example() {
	db, err := sql.Open("taossqlrestful", "mydb://dalong@127.0.0.1/demoapp")
	if err != nil {
		log.Fatalf("some error %s", err.Error())
	}
	rows, err := db.Query("select * from demoapp")
	if err != nil {
		log.Println("some wrong for query", err.Error())
	}
	for rows.Next() {
		rows.Scan()
	}
}
