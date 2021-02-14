package taossqlrestful

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// TaosResq is
type TaosResq struct {
	Status string          `json:"status"`
	Head   []string        `json:"head"`
	Data   [][]interface{} `json:"data"`
	Rows   int64           `json:"rows"`
	Code   int             `json:"code"`
	Desc   string          `json:"desc"`
}

// TaosConnect is
func (mc *Conn) TaosConnect(ip, user, pass, db string, port int) (taos string, err error) {
	b := []byte(fmt.Sprintf("%s:%s", user, pass))
	body := strings.NewReader("SELECT CLIENT_VERSION()")
	sEnc := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("enc=[%s]\n", sEnc)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/rest/sql", ip, port), body)
	req.Header.Add("Authorization", "Basic "+sEnc)
	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		return "", errors.New("taos_connect() fail! ")
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	taosResq := new(TaosResq)
	jsonErr := json.Unmarshal(content, &taosResq)
	if jsonErr != nil {
		fmt.Printf("%s", jsonErr)
	}
	fmt.Println(taosResq.Data)

	return db, nil
}

// func (mc *Conn) taosQuery(sqlstr string) (int, error) {

// 	csqlstr := C.CString(sqlstr)
// 	defer C.free(unsafe.Pointer(csqlstr))
// 	if mc.result != nil {
// 		C.taos_free_result(mc.result)
// 		mc.result = nil
// 	}
// 	mc.result = unsafe.Pointer(C.taos_query(mc.taos, csqlstr))
// 	//mc.result = unsafe.Pointer(C.taos_query_c(mc.taos, csqlstr.Str, C.uint32_t(csqlstr.Len)))
// 	code := C.taos_errno(mc.result)
// 	if 0 != code {
// 		errStr := C.GoString(C.taos_errstr(mc.result))
// 		mc.taos_error()
// 		return 0, errors.New(errStr)
// 	}

// 	// read result and save into mc struct
// 	num_fields := int(C.taos_field_count(mc.result))
// 	if 0 == num_fields { // there are no select and show kinds of commands
// 		mc.affectedRows = int(C.taos_affected_rows(mc.result))
// 		mc.insertId = 0
// 	}

// 	return num_fields, nil
// }

// func (mc *Conn) taos_close() {
// 	mc.taos = nil
// }
