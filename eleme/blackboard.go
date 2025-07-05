package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// TrieNode 表示字典树的节点
type TrieNode struct {
    children [2]*TrieNode
    count    [2]int
}

// Trie 表示字典树
type Trie struct {
    root  *TrieNode
    index int
}

// NewTrie 创建一个新的字典树
func NewTrie() *Trie {
    return &Trie{
        root:  &TrieNode{},
        index: 0,
    }
}

// insert 向字典树中插入一个数字，并根据 mod 增加或减少计数
func (t *Trie) insert(number, mod int) {
    node := t.root
    for i := 31; i >= 0; i-- {
        bit := (number >> i) & 1
        if node.children[bit] == nil {
            node.children[bit] = &TrieNode{}
            t.index++
        }
        node.count[bit] += mod
        node = node.children[bit]
    }
}

// getMaxXor 计算给定数字的最大异或值
func (t *Trie) getMaxXor(number int) int {
    node := t.root
    result := 0
    for i := 31; i >= 0; i-- {
        bit := (number >> i) & 1
        oppositeBit := 1 - bit
        if node.count[oppositeBit] > 0 {
            result += 1 << i
            node = node.children[oppositeBit]
        } else {
            node = node.children[bit]
        }
    }
    return result
}

func main() {
    trie := NewTrie()
    count := 0

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("请输入操作数量: ")
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    for i := 0; i < n; i++ {
        fmt.Print("请输入操作类型和数字: ")
        scanner.Scan()
        input := strings.Split(scanner.Text(), " ")
        a, _ := strconv.Atoi(input[0])
        b, _ := strconv.Atoi(input[1])

        if a == 1 {
            count++
            trie.insert(b, 1)
        } else if a == 2 {
            count--
            trie.insert(b, -1)
        } else {
            if count == 0 {
                fmt.Println(-1)
            } else {
                fmt.Println(trie.getMaxXor(b))
            }
        }
    }
}
