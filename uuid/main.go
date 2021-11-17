package main

import (
	"fmt"
	"github.com/sony/sonyflake"
)

func main()  {
	st := sonyflake.Settings{}
	st.MachineID = func() (uint16, error) {
		return 1, nil
	}
	sk := sonyflake.NewSonyflake(st)
	fmt.Println(sk.NextID())
}
