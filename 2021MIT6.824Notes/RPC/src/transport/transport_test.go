package transport

import (
	"net"
	"sync"
	"testing"
	"time"
)

func TestTransport_Read(t *testing.T) {

	address := "localhost:3213"
	dataToSend := "Hello TLV World"
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		l,err := net.Listen("tcp",address)
		if err != nil{
			t.Fatal(err)
		}
		defer l.Close()
		conn, _ := l.Accept()
		time.Sleep(100 * time.Millisecond)
		s:= NewTransport(conn)

		err = s.Send([]byte(dataToSend))
		t.Log("Listen and accept")
		if err != nil{
			t.Fatal(err)
		}
	}()

	go func() {

		defer wg.Done()
		conn,err := net.Dial("tcp",address)
		if err != nil{
			t.Fatal(err)
		}

		tp := NewTransport(conn)
		data,err := tp.Read()
		if err != nil{
			t.Fatal(err)
		}
		if string(data) != dataToSend{
			t.FailNow()
		}
	}()

	wg.Wait()

}
