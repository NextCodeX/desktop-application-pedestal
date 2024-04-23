package main

import (
	"fmt"
	"github.com/changlie/go-common/a"
	"github.com/shirou/gopsutil/disk"
	"testing"
)

func Test_os(t *testing.T) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return
	}
	fmt.Println(a.JsonOf(partitions).EncodePretty())
}
