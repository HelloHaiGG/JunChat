package utils

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"encoding/json"
)

//将src中的值拷贝给dst
//对象值拷贝

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

//map2struct
func Map2Struct(dst, src interface{}) error {
	if b, err := json.Marshal(&src); err != nil {
		return err
	} else {
		return json.Unmarshal(b, &dst)
	}
}

//将 sql 查询结果 Scan To Map

func ScanResultToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	length := len(columns)
	res := make([]map[string]interface{}, 0)
	for rows.Next() {
		current := makeResultReceiver(length)
		if err := rows.Scan(current...); err != nil {
			panic(err)
		}
		value := make(map[string]interface{})
		for i := 0; i < length; i++ {
			k := columns[i]
			v := current[i]
			value[k] = v
		}
		res = append(res, value)
	}
	return res, nil
}

func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}
