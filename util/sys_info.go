package util

import "github.com/shirou/gopsutil/disk"

func DiskPartitions() []disk.PartitionStat {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil
	}
	return partitions
}
