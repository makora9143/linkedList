package main

import "fmt"

type Cell struct {
    value int
    next *Cell
}

type LinkedList struct {
    top *Cell
}

func newCell(x int, nextcell *Cell) *Cell {
    var cell = new(Cell)
    cell.value = x
    cell.next = nextcell
    return cell
}

func newLinkedList() *LinkedList {
    var llst = new(LinkedList)
    llst.top = new(Cell)
    return llst
}

func (cell *Cell) search(index int) *Cell {
    i := -1
    for cell != nil {
        if i == index {
            return cell
        }
        i++
        cell = cell.next
    }
    return nil
}

func (llst *LinkedList) get(index int) int {
    targetcell := llst.top.search(index)
    if targetcell == nil {
        return 0
    }
    return targetcell.value
}

func (llst *LinkedList) insert(index int, x int) bool {
    prevcell := llst.top.search(index-1)
    nextcell := prevcell.next
    if prevcell == nil {
        return false
    }
    newcell := newCell(x, nextcell)
    prevcell.next = newcell
    return true
}

func (llst *LinkedList) delete(index int) bool {
    prevcell := llst.top.search(index-1)
    if prevcell == nil || prevcell.next == nil{
        return false
    }
    prevcell.next = prevcell.next.next
    return true
}

func (llst *LinkedList) size() int {
    cell := llst.top.search(1)
    var i = 0
    for cell != nil {
        i++
        cell = cell.next
    }
    return i
}

func (llst *LinkedList) showAll() {
    cell := llst.top.search(1)
    for cell != nil {
        fmt.Printf("%d, ", cell.value)
        cell = cell.next
    }
    fmt.Println("")
}

func main() {
    llst := newLinkedList()
    fmt.Println("Start insert")
    for i := 0; i < 10; i++ {
        fmt.Println(llst.insert(i, i))
    }
    fmt.Println("Finish insert")
    llst.showAll()
    fmt.Println(llst.size())
    fmt.Println(llst.get(4))
    for llst.size() != 0 {
        llst.delete(0)
        llst.showAll()
    }
    fmt.Println("Delete all")
}