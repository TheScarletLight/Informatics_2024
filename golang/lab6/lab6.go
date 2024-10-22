package main

import "fmt"

type PC struct {
	diskSize int
}

func (c *PC) GetDiskSize() int {
	return c.diskSize
}
func (c *PC) SetDiskSize(size int) {
	if size >= 0 {
		c.diskSize = size
	}
}
func main() {
	var diskSize int = 500
	fmt.Printf("Объем диска", diskSize)
}
