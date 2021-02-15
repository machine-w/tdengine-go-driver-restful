package taossqlrestful

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/machine-w/tdengine-go-driver-restful/taosSqlRestful"
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

func TestExec(t *testing.T) {
	db, err := sql.Open("taossqlrestful", "root:msl110918@/http(121.36.56.117:6041)/dianli1")
	if err != nil {
		t.Errorf("some error %s", err.Error())
	}
	stmt, err := db.Prepare("insert into node_5(ts,wdtc3,wdtc4) values(?,?)")
	if err != nil {
		log.Println(err)
	}
	rs, err := stmt.Exec("2020-02-10 00:00:00", 11, 12)
	if err != nil {
		log.Println(err)
	}
	log.Println(rs)
	//我们可以获得插入的id
	// id, err := rs.LastInsertId()
	//可以获得影响行数
	// _, err = rs.RowsAffected()
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(affect)
}

// func Example() {
// 	db, err := sql.Open("taossqlrestful", "mydb://dalong@127.0.0.1/demoapp")
// 	if err != nil {
// 		log.Fatalf("some error %s", err.Error())
// 	}
// 	rows, err := db.Query("select * from demoapp")
// 	if err != nil {
// 		log.Println("some wrong for query", err.Error())
// 	}
// 	for rows.Next() {
// 		rows.Scan()
// 	}
// }
