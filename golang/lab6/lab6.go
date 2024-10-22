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
	PC := PC{}
	PC.SetDiskSize(512)
	fmt.Printf("Объем жесткого диска:", PC.GetDiskSize())
}
