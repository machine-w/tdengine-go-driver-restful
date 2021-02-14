package taossqlrestful

import (
	"database/sql"
	"log"
	"testing"
)

// func TestConn(t *testing.T) {
// 	conn := taossqlrestful.MakeConn()
// 	// conn.TaosConnect("121.36.56.117", "root", "msl110918", "dianli", 6041)
// 	// fmt.Println(conn)
// }

func TestDb(t *testing.T) {
	db, err := sql.Open("taossqlrestful", "root:msl110918@/http(121.36.56.117:6041)/dianli1")
	if err != nil {
		t.Errorf("some error %s", err.Error())
	}
	// fmt.Printf("%+v", db)
	// fmt.Println(db)
	rows, err := db.Query("select ts,wdtc4,wdtc3 from node_5")
	if err != nil {
		log.Fatal("some wrong for query", err.Error())
	}
	for rows.Next() {
		var times string
		var val1 float32
		var val2 float32
		if err := rows.Scan(&times, &val1, &val2); err != nil {
			log.Println("scan value erro", err.Error())
		} else {
			log.Println(val2, val1, times)
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
