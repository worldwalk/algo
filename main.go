package main

import (
	"errors"
	"fmt"
	"time"
)

// type FeedbackVO struct {
// 	Content string `json:"comment"`
// 	UserID  string `json:"uin"`
// 	Time    string `json:"time"` // e.g. "2023-08-01 10:12:27"
// }
// type getFeedbackRsp struct {
// 	Msg  string `json:"msg"`
// 	Code int    `json:"code"`
// 	Data struct {
// 		Total int           `json:"total"`
// 		List  []*FeedbackVO `json:"list"`
// 	} `json:"data"`
// }

func testSlice(path []int) {
	path = append(path, 6)
	fmt.Println(path)
}

func main() {

	path := []int{1, 2, 3, 4, 5}
	testSlice(path)
	fmt.Println(path)
	// // 调用方法读取并反序列化文件内容
	// fbs, err := readAndDeserialize("output.json1")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// // 打印反序列化后的结果
	// for _, fb := range fbs {
	// 	fmt.Printf("ID: %s, Name: %s\n", fb.Message, fb.AppId)
	// 	//fmt.Printf("%+v\n", fb)
	// }
	//fmt.Println(len(fbs))
	//cache.LRUTest()
	//maze.MazeDFSTest()
	//maze.MazeBFSTest()
	//maze.MazeBacktrackingTest()
	//sort.QuickSortTest()
	//sort.HeapSortTest()
	//sort.MergeSortTest()
	//sort.BubbleSortTest()

	//list.ReverseTest()

	//search.BinarySearchTest()

	//n := 10000
	//n = 10
	//for i := 1; i <= n; i++ {
	//	fact := factorial(i)
	//	fmt.Printf("%d! = %s\n", i, fact.String())
	//}

	// test4()
	// p1 := Person{Name: "zhansan", Age: 20}
	// p2 := Person{Name: "zhansan", Age: 20}
	// fmt.Println(p1 == p2)
}

type Person struct {
	Name string
	Age  int
}
type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func test4() {

	//err := &CustomError{Message: "custom error"}
	err := fmt.Errorf("testadd %w", &CustomError{Message: "custom error"})

	fmt.Println(err)
	// 使用 errors.Is 判断错误类型
	if errors.Is(err, err) {
		fmt.Println("This is a CustomError")
	}

	// 使用 errors.As 提取原始错误
	var targetErr *CustomError
	if errors.As(err, &targetErr) {
		fmt.Println("Extracted CustomError:", targetErr.Message)
	}

}

func test3(b bool) {
	fmt.Println("test3 in")
	if b {
		defer func() {
			fmt.Println("defer xxx")
		}()
	}
	fmt.Println("test3 done")
}

func test2() {
	fmt.Println(time.Month(12 + 1))
	fmt.Println(time.Date(2023, time.Month(11), 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second).Format(time.DateTime))

}

var m map[string]interface{}
var mstr map[string]string

func test() {
	m = make(map[string]interface{})
	mstr = make(map[string]string)
	f := float64(1.1223344556677e+13)
	m["a"] = f
	//str := fmt.Sprintf("%g", m["a"])
	//str := strconv.FormatFloat(f, 'f', -1, 64)
	str := fmt.Sprintf("由深度学习模型预测，置信度：%.4f", 0.982923)
	fmt.Println(str)
	fmt.Println(get())

	//var rsp getFeedbackRsp
	//if err := json.Unmarshal([]byte("{\"msg\": \"success\",\"code\": 0}"), &rsp); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("nil:%v", rsp.Data.List)
}

func get() string {
	return mstr["ni"]
}
