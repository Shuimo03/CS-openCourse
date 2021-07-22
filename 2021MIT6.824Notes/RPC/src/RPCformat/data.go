package RPCformat

import (
	"bytes"
	"encoding/gob"

//define client send data structure

type RPCdata struct {
	Name string
	Age int
	Sex string
	Err error
}

/*
Encode The RPCdata in binary format which can be sent over the network
*/

func Encode(data RPCdata) ([]byte,error){
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil{
		return nil,err
	}

	return buf.Bytes(), nil
}

func Decode(b []byte)(RPCdata, error){
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	var data RPCdata
	if err := decoder.Decode(&data); err != nil{
		return RPCdata{},err
	}
	return data,nil
}

