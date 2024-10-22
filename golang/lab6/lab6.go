package main

import "fmt"

type PC struct {
	diskSize int
}

func (c *PC) SetDiskSize(size int) {
	c.diskSize = size
}

func (c *PC) GetDiskSize() int {
	return c.diskSize
}

func main() {
	minePC := PC{}
	minePC.SetDiskSize(500)
	fmt.Printf("Размер диска:", minePC.GetDiskSize())
}
