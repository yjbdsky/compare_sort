// gettime project main.go
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println("err:", err)
	}
}

type ST struct {
	name string
	age  int
}
type STsl []ST

func (a STsl) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a STsl) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a STsl) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].age < a[i].age
}

func main() {
	var distinct, col1, col2 int
	var file string
	var max bool
	flag.StringVar(&file, "f", "", "file,file2,file3...")
	flag.BoolVar(&max, "max", true, "max or min 取最大值或最小值")
	flag.IntVar(&distinct, "d", 0, "distinct line")
	flag.IntVar(&col1, "c1", 1, "max line")
	flag.IntVar(&col2, "c2", 2, "min line")
	flag.Parse()
	fmt.Println(file, col1, col2)
	if file == "" {
		fmt.Println("input file")
		return
	}
	comp := make(map[string][][]byte)
	filearr := strings.Split(file, ",")

	for _, v := range filearr {
		all, err := ioutil.ReadFile(v)
		check(err)
		all = bytes.Replace(all, []byte("\r"), []byte{}, -1)
		allarr := bytes.Split(all, []byte("\n"))
		for _, line := range allarr {
			linearr := bytes.Split(line, []byte(" "))
			if len(linearr) < distinct || len(linearr) < col1 || len(linearr) < col2 {
				fmt.Println("line err:", line)
				continue
			}
			key := string(linearr[distinct])
			//fmt.Println("debug:", comp[key])
			if comp[key] == nil {
				comp[key] = append(linearr, []byte("0"))
			}
			c1, err := strconv.Atoi(string(linearr[col1]))
			check(err)
			c2, err := strconv.Atoi(string(linearr[col2]))
			check(err)
			m1, err := strconv.Atoi(string(comp[key][col1]))
			check(err)
			m2, err := strconv.Atoi(string(comp[key][col2]))
			check(err)
			if max {
				if c1 > m1 {
					comp[key][col1] = linearr[col1]
				}

				if c2 > m2 {
					comp[key][col2] = linearr[col2]
				}
				delta := strconv.Itoa(c1 - c2)
				comp[key][len(linearr)] = []byte(delta)
			} else {
				if c1 < m1 {
					comp[key][col1] = linearr[col1]
				}

				if c2 < m2 {
					comp[key][col2] = linearr[col2]
				}
				comp[key][len(linearr)] = []byte(strconv.Itoa(c2 - c1))
			}

		}
	}
	//排序
	sortarr := []ST{}
	for k, v := range comp {
		d, err := strconv.Atoi(string(v[len(v)-1]))
		check(err)
		sortarr = append(sortarr, ST{k, d})
	}
	//fmt.Println("sort befrom:", sortarr)
	sort.Sort(STsl(sortarr))
	//	fmt.Println("after befrom:", sortarr)
	//	return

	var dst []byte
	for _, so := range sortarr {
		v := comp[so.name]
		var tmp []byte
		for _, b := range v {
			tmp = bytes.Join([][]byte{tmp, b}, []byte(" "))
		}
		tmp = bytes.TrimPrefix(tmp, []byte(" "))
		dst = bytes.Join([][]byte{dst, tmp}, []byte("\n"))
	}
	fmt.Println(string(dst))
	//fmt.Println(distinct, line1, line2)

}
