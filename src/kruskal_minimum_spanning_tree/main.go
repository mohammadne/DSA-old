package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

type Edge struct {
	u int
	v int
	w int
}

var cnt int
var fa [200005]int

type shuzu []Edge

func (a shuzu) Len() int {
	return len(a)
}
func (a shuzu) Less(i, j int) bool {
	return a[i].w < a[j].w
}
func (a shuzu) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	n := ReadInt()
	m := ReadInt()
	edge := make([]Edge, m+1)
	for i := 1; i <= m; i++ {
		edge[i].u = ReadInt()
		edge[i].v = ReadInt()
		edge[i].w = ReadInt()
	}
	for i := 1; i <= n+1; i++ {
		fa[i] = i
	}
	sort.Sort(shuzu(edge))
	var tot int = 0
	var ans int64 = 0
	for i := 1; i <= m; i++ {
		u := find(edge[i].u)
		v := find(edge[i].v)
		if u != v {
			ans += int64(edge[i].w)
			fa[u] = v
			tot++
		}
		if tot == n-1 {
			break
		}
	}
	if tot == n-1 {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}

}

////////////////////////////////////
func find(x int) int {
	p := x
	var tmp int = 0
	for p != fa[p] {
		p = fa[p]
	}
	for x != p {
		tmp = fa[x]
		fa[x] = p
		x = tmp
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var out chan string
var in *bufio.Scanner
var outWg *sync.WaitGroup

func ReadString() string {
	in.Scan()
	return in.Text()
}

func ReadStringSlice(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = ReadString()
	}
	return s
}

func ReadInt() int {
	intStr := ReadString()
	i, _ := strconv.ParseInt(intStr, 10, 32)
	return int(i)
}

func ReadIntSlice(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = ReadInt()
	}
	return arr
}

func init() {
	//set input
	in = bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 1024), int(2e+5))
	in.Split(bufio.ScanWords)

	//set output
	out = make(chan string, 16)
	outWg = &sync.WaitGroup{}
	outWg.Add(1)

	writer := bufio.NewWriterSize(os.Stdout, int(2e+5))
	go func(write *bufio.Writer) {
		defer outWg.Done()
		defer write.Flush()

		for line := range out {
			write.WriteString(line + "\n")
		}
	}(writer)
}
