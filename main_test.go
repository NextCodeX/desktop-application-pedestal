package main

import (
	"fmt"
	"github.com/changlie/go-common/a"
	"github.com/shirou/gopsutil/disk"
	"sort"
	"testing"
)

func Test_os(t *testing.T) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return
	}
	fmt.Println(a.JsonOf(partitions).EncodePretty())
}

func Test_sort(t *testing.T) {
	arr := []string{"a", "F", "DFE", "WW", "C", "c", "1", "99", "3"}
	sort.Strings(arr)
	fmt.Println(arr)
}
func Test_arr(t *testing.T) {
	arr := make([]int, 9)
	arr[2] = 998
	fmt.Println(arr)
}
