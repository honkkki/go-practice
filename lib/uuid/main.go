package main

import (
	"fmt"

	"github.com/sony/sonyflake"
)

func main() {
	st := sonyflake.Settings{}
	st.MachineID = func() (uint16, error) {
		return 1, nil
	}
	sk := sonyflake.NewSonyflake(st)
	uid, _ := sk.NextID()
	fmt.Println(uid)
	fmt.Println(uid % 10)

}
