package main

import (
	"fmt"
	"go-practice/protobuf/phone"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/golang/protobuf/proto"
)

func writeProto(filename string) (err error) {
	var contactBook phone.ContactBook
	for i := 0; i < 10; i++ {
		person := &phone.Person{
			Id:   int32(i + 1),
			Name: fmt.Sprintf("name%d", i+1),
		}

		phoneInfo1 := &phone.Phone{
			Type:   phone.PhoneType_HOME,
			Number: strconv.Itoa(rand.Intn(99999999)),
		}

		phoneInfo2 := &phone.Phone{
			Type:   phone.PhoneType_WORK,
			Number: strconv.Itoa(rand.Intn(99999999)),
		}
		person.Phone = append(person.Phone, phoneInfo1, phoneInfo2)
		contactBook.Person = append(contactBook.Person, person)
	}

	data, err := proto.Marshal(&contactBook)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		return err
	}

	return nil
}

func readProto(filename string) (err error) {
	var contactBook phone.ContactBook
	data, _ := ioutil.ReadFile(filename)
	err = proto.Unmarshal(data, &contactBook)
	if err != nil {
		return err
	}

	fmt.Println(contactBook)
	fmt.Println(contactBook.Person[0].Phone[0].Type)
	return nil
}

func main() {
	filename := "./protobuf.txt"
	err := writeProto(filename)
	if err != nil {
		fmt.Println(err)
	}

	err = readProto(filename)
	if err != nil {
		fmt.Println(err)
	}
}
