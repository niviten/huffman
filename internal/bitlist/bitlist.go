package bitlist

import "fmt"

type BitList struct {
    buf byte
    arr []byte
    bufsize int
}

func New() *BitList {
    return &BitList{
        buf: 0,
        arr: make([]byte, 0),
        bufsize: 0,
    }
}

func (list *BitList) Add(bit uint8) {
    list.buf = list.buf << 1
    if bit != 0 {
        list.buf = list.buf | 1
    }
    list.bufsize = list.bufsize + 1
    if list.bufsize == 8 {
        list.arr = append(list.arr, list.buf)
        list.bufsize = 0
        list.buf = 0
    }
}

func (list *BitList) Print() {
    fmt.Println("{")
    for _, a := range list.arr {
        fmt.Printf("\t")
        printByte(a, 8)
        fmt.Println()
    }
    fmt.Printf("\t")
    printByte(list.buf, list.bufsize)
    fmt.Println()
    fmt.Println("}")
}

func printByte(b byte, c int) {
    for i := c - 1; i >= 0; i-- {
        fmt.Printf("%d ", (b >> i) & 1)
    }
}
