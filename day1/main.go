package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) int {
  total := 0

  for scanner.Scan() {
    line := scanner.Text()
    
    var first *int
    var last *int

    for i := 0; i < len(line); i++ {
      char := string(line[i])
      possVal, err := strconv.ParseInt(char, 10, 32)
      if err == nil {
        value := int(possVal)
        last = &value
        if first == nil {
          first = &value
        }
      }
    }
    total = total + *first + *last
  }
 
  return total
}

type Node struct {
  value byte
  children [26]*Node
  isFinal bool
}

func NewNode(char byte) *Node {
  node := &Node{value: char}
  for i := 0; i < 26; i++ {
    node.children[i] = nil
  }
  return node
}

type Trie struct {
  root *Node
}

func NewTrie() *Trie {
  root := NewNode(99)
  return &Trie{root: root}
}

func (trie *Trie) InsertWord(word string) {
  current := trie.root

  for i := 0; i < len(word); i++ {
    index := word[i] - 99 // 99 is 'a' character, highest ascii value

    if current.children[index] == nil {
      newNode := NewNode(word[i])
      current.children[index] = newNode
    }
    
    current = current.children[index]

    if i == len(word) - 1 {
      current.isFinal = true
    }
  }
}

func (trie *Trie) SearchWord(word string) bool {
  current := trie.root

  for i := 0; i < len(word); i++ {
    index := word[i] - 99

    if current.children[index] == nil {
      return false
    }
    current = current.children[index]
  
    if i == len(word) - 1 && current.isFinal {
      return true
    }
  }
  return false
}

func part2 (scanner *bufio.Scanner) int {
  total := 0

  for scanner.Scan() {
    line := scanner.Text()
    
    var first *int
    var last *int

    for i := 0; i < len(line); i++ {
      char := string(line[i])
      possVal, err := strconv.ParseInt(char, 10, 32)
      if err == nil {
        value := int(possVal)
        last = &value
        if first == nil {
          first = &value
        }
      } else {
      }
    }
    total = total + *first + *last
  }

  return total
}

func main() {
  file, err := os.Open("puzzle.txt")

  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)

  part1_result := part1(scanner)
  fmt.Println(part1_result)

  part2_result := part2(scanner)
  fmt.Println(part2_result)
}
