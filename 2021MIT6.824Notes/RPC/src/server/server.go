package server

import (
	"reflect"
	"simpleRPC/src/RPCformat"

/**
	RPC Server will receive the RPCdata that will have an function Name.
	So we need to maintain and map that contains an function name to actual
	function mapping
 */

type RPCServer struct {
	addres string
	funcs map[string]reflect.Value
}

func NewServer(addres string) *RPCServer{
	return &RPCServer{addres: addres, funcs: make(map[string]reflect.Value)}
}

func (s *RPCServer) Register(fnName string,fFunc interface{}){
	if _, ok := s.funcs[fnName]; ok{
		return
	}
	s.funcs[fnName] = reflect.ValueOf(fFunc)
}

// Execute the given function if  present
func (s *RPCServer) Execute( RPCdata){

}