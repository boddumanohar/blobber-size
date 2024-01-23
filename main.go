package main
import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	var volStat unix.Statfs_t
	dir := "/var"
	fmt.Println("dir: ", dir+"/0chain/blobber/hdd")
	err := unix.Statfs(dir+"/0chain/blobber/hdd", &volStat)
	if err != nil {
		panic(err)
	}

	diskCapacity := volStat.Bavail * uint64(volStat.Bsize)
	fmt.Println("capacity: ", diskCapacity)
}
