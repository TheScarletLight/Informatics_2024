package main

import "fmt"

type PC struct {
	diskSize int
}

func (p *PC) SetDiskSize(size int) {
	p.diskSize = size
}

func (p *PC) GetDiskSize() int {
	return p.diskSize
}

func main() {
	minePC := PC{}
	minePC.SetDiskSize(500)
	fmt.Printf("Размер диска: %d GB\n", minePC.GetDiskSize())
}
