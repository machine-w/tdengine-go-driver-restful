package taossqlrestful

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

type TaosResp struct {
	Status string          `json:"status"`
	Head   []string        `json:"head"`
	Data   [][]interface{} `json:"data"`
	Rows   int             `json:"rows"`
	Code   int             `json:"code"`
	Desc   string          `json:"desc"`
}

// TaosConnect is
func (mc *taosConn) taosConnect(ip, user, pass, db string, port int) (taos string, err error) {
	mc.db = db
	b := []byte(fmt.Sprintf("%s:%s", user, pass))
	sqlStr := strings.NewReader("SELECT CLIENT_VERSION()")
	mc.token = base64.StdEncoding.EncodeToString(b)
	mc.reqUrl = fmt.Sprintf("http://%s:%d/rest/sqlt", ip, port)
	wrapper := NewWrapper(mc.token)
	taosResp, err := wrapper.Post(mc.reqUrl, sqlStr)
	if err != nil {
		return "", err
	}
	mc.dbVersion = taosResp.Data[0][0].(string)
	return taosResp.Status, nil
}

func (mc *taosConn) taosQuery(sqlstr string) (int, error) {
	mc.mu.Lock()
	wrapper := NewWrapper(mc.token)
	useDbResp, err := wrapper.Post(mc.reqUrl, strings.NewReader("use "+mc.db))
	if err != nil {
		return 0, err
	}
	if useDbResp == nil {
		return 0, errors.New("bad return ")
	}
	selectResp, err := wrapper.Post(mc.reqUrl, strings.NewReader(sqlstr))
	if err != nil {
		return 0, err
	}

	mc.result = selectResp
	numFields := len(selectResp.Head)
	if numFields == 1 && selectResp.Head[0] == "affected_rows" { // there are no select and show kinds of commands
		mc.affectedRows = int(selectResp.Data[0][0].(float64))
		mc.insertId = 0
		numFields = 0
	}
	mc.mu.Unlock()
	return numFields, nil
}

func (mc *taosConn) taos_close() {
	mc.taos = ""
}
