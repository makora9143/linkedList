/**
*   @author  Makoto Kawano
*   @version 0.1
*/

package main

import "fmt"

/**
*   Cell構造体
*   valueで値を保持
*   nextで次のCellポインタを示す．
*/
type Cell struct {
    value int
    next *Cell
}

/**
*   LinkedList構造体
*   最初のCellを保持
*/
type LinkedList struct {
    top *Cell
}

/**
*   Cell構造体を生成する
*   @param x 値
*   @param nextcell 次のセル
*   @return cell 生成したCell構造体
*/
func newCell(x int, nextcell *Cell) *Cell {
    var cell = new(Cell)
    cell.value = x
    cell.next = nextcell
    return cell
}

/**
*   LinkedList構造体を生成する
*   @param
*   @return llst 生成したLinkedList構造体
*/
func newLinkedList() *LinkedList {
    var llst = new(LinkedList)
    llst.top = new(Cell)
    return llst
}

/**
*   ベースとなるメソッド
*   @param index index番目
*   @return index番目のCell
*/
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

/**
*   getメソッド
*   @param index index番目
*   @return target.value if 値があれば or nil if Cellがなければ
*/
func (llst *LinkedList) get(index int) int {
    targetcell := llst.top.search(index)
    if targetcell == nil {
        return 0
    }
    return targetcell.value
}

/**
*   insertメソッド
*   @param index index番目
*   @param x 挿入したい値
*   @return true if 成功したら or false 失敗したら
*/
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

/**
*   deleteメソッド
*   @param index index番目
*   @return true if 成功したら or false 失敗したら
*/
func (llst *LinkedList) delete(index int) bool {
    prevcell := llst.top.search(index-1)
    if prevcell == nil || prevcell.next == nil{
        return false
    }
    prevcell.next = prevcell.next.next
    return true
}

/**
*   要素数を取得する
*   @param
*   @return 要素数
*/
func (llst *LinkedList) size() int {
    cell := llst.top.search(1)
    var i = 0
    for cell != nil {
        i++
        cell = cell.next
    }
    return i
}

/**
*   LinkedList内の要素を全て表示する
*   @param
*   @return
*/
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