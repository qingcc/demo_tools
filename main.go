package main

import (
	"fmt"
	"github.com/qingcc/demo_tools/util"
	"github.com/tealeg/xlsx"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	f()
	//read("country_code.xlsx")
	time.Sleep(time.Hour)
	newRandStream := func() <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			// 死循环：不断向channel中放数据，直到阻塞
			for {
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")

	// 只消耗3个数据，然后去做其他的事情，此时生产者阻塞，
	// 若主goroutine不处理生产者goroutine，则就产生了泄露
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())
	time.Sleep(10e9)
	fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())
}

func f() {
	for i := 0; i < 10; i++ {
		//num := fi(i)
		//num := lib(i)
		//num := libTail(1, 1, i)
		num := qpow(3, i)
		//if i %10 == 0 || i%10==1{
		//	m[i] = num
		//}

		println("i:", i, "num:", num)
	}
	log.Println("num len:", len(m))
}

var m = make(map[int]int)

//递归
func fi(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		if num, ok := m[n]; ok {
			return num
		}
		return fi(n-1) + fi(n-2)
	}
}

//优化1： 迭代
func lib(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		n1, n2, n3 := 1, 1, 0
		for i := 2; i <= n; i++ {
			n3 = n1 + n2
			n1 = n2
			n2 = n3
		}
		return n3
	}
}

//尾递归
func libTail(first, second, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return first + second
	}
	return libTail(second, first+second, n-1)
}

//base 底数 ，exp 指数 未完成（矩阵乘法，快速幂）
func qpow(base, exp int) int {
	if 0 == exp {
		return 1
	}

	ret := 1
	for exp != 0 {
		if exp&1 != 0 { //exp最右边一位 按位与&
			ret = ret * base
		}
		base = base * base
		exp >>= 1 //右移一位
	}

	return ret
}

func read(name string) {
	f, err := xlsx.OpenFile(name)
	if err != nil {
		log.Println("open xlsx file failed:", err)
	}
	cols := []string{"s_region_id", "country", "city", "s_hotel_id_count", "s_country_code", "d_country_code", "region_name_cn", "region_name_long_cn", "region_type"}
	data := make([][]string, len(cols))
	rows := f.Sheets[0].Rows
	for i, row := range rows[1:] {
		dump := false
		if len(row.Cells) < 6 {
			dump = true
		} else {
			s_country_code := row.Cells[4].String()
			d_country_code := row.Cells[5].String()
			if s_country_code != d_country_code || dump {
				log.Println("i:", i)
				itemRow := make([]string, 0, len(cols))
				for _, cell := range row.Cells {
					itemRow = append(itemRow, cell.String())
				}
				data = append(data, itemRow)
			}
		}
	}
	util.NewXlsx("newmap.xlsx", cols, data)
	log.Println("over")
}

func init() {
	a := 123456789
	log.Println(fmt.Sprintf("%.f", a))
	log.Println(fmt.Sprintf("%.f", a) == strconv.FormatInt(int64(a), 10))
	layout := "2006-01-02 15:04:05"
	startSec := time.Unix(time.Now().Unix()-24*60*60, 0)
	startTime := startSec.Format(layout)
	log.Println("time:", startTime)
}
