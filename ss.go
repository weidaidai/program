package main

import "fmt"

func main() {
	var nums []int

	// fmt.Printf("%T\n", nums)
	// fmt.Println(nums == nil)
	// //字面量
	// nums = []int{1, 2, 3, 4}
	// fmt.Println(nums)
	// nums = []int{1, 2, 3}
	// fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	// //数组qiepian赋值
	// var ss [10]int = [10]int{1, 2, 3, 3, 45}
	// nums = ss[1:10]
	// fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	// //make
	// nums = make([]int, 3, 5)
	// fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	// nums[1] = 1000            //修改
	// nums = append(nums, 9999) //增加
	// fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(i, nums[i])

	// }
	// fmt.Println(nums[1:5])
	// nums = make([]int, 3, 5)
	// n := nums[3:3:6]
	// fmt.Printf("%#v %d %d\n", n, len(n), cap(n))
	nums = make([]int, 6, 6)
	nums02 := nums[1:4:5]
	fmt.Println(nums, nums02)
	nums02[2] = 1000
	fmt.Println(nums, nums02)
	nums02 = append(nums02, 3)
	fmt.Println(nums, nums02)
	nums = append(nums, 5) //对nums02无影响，已经超出nunm02的容量范围
	fmt.Println(nums, nums02)
	//删除
	num0 := []int{1, 3, 4, 4}
	num1 := []int{10, 20, 30, 40, 50}
	copy(num0, num1)
	fmt.Println(num0, num1)
	num0 = num0[:len(num0)-1] //删除最后元素
	num0 = num0[1:]           //删除前面元素
	fmt.Println(num0)
	//删除中间元素
	copy(num1[2:], num1[3:])
	fmt.Println(num1)
}
