package es

import "testing"

func TestEs(t *testing.T) {
	InitEs()
	//Read()

	//err := Del()
	//if err != nil {
	//	t.Log(err)
	//}

	//InsertData(Tweet{
	//	Username: "karina",
	//	Message:  "hello",
	//})

	ReadData("karina")
}
