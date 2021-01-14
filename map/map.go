package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["name"] = "jisoo"
	map1["a"] = "jisoo"
	map1["s"] = "jisoo"
	map1["d"] = "jisoo"
	map1["c"] = "jisoo"
	map1["x"] = "jisoo"
	map1["z"] = "jisoo"
	fmt.Println(map1)
	fmt.Printf("%#v \n", map1)

	map2 := map[string]string{
		"name": "jisoo",
		"age":  "20",
	}

	fmt.Println(map2)
	_, ok := map1["age"] // 判断key是否存在
	fmt.Println(ok)
	res := map1["name"]
	fmt.Println(res)
	delete(map2, "age") // 删除键值对
	fmt.Println(map2)

	s1 := make([]map[string]string, 5) // 只初始化了切片 里面的map也需要初始化 如下
	fmt.Println(s1)
	s1[0] = make(map[string]string, 1) // 初始化内部map
	s1[0]["name"] = "jisoo"
	fmt.Println(s1)

	map3 := make(map[int][]int, 5) // 值为切片的map
	map3[0] = []int{1}
	fmt.Println(map3)

}
