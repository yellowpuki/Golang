package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeNode struct {
	value      int
	leftChild  *TreeNode
	rightChild *TreeNode
}

func (t *TreeNode) Insert(value int) bool {
	if t == nil {
		return false
	}

	if t.value == value {
		return false
	}

	if t.value > value {
		if t.leftChild == nil {
			t.leftChild = &TreeNode{value: value}
			return true
		}
		return t.leftChild.Insert(value)
	}

	if t.value < value {
		if t.rightChild == nil {
			t.rightChild = &TreeNode{value: value}
			return true
		}
		return t.rightChild.Insert(value)
	}
	return true
}

func (t *TreeNode) Exist(value int) bool {
	tn := t.FindNode(value)
	return tn != nil
}

func (t *TreeNode) FindNode(value int) *TreeNode {
	if t.value == value {
		return t
	} else {
		if t.value > value {
			if t.leftChild != nil {
				return t.leftChild.FindNode(value)
			} else {
				return nil
			}
		} else {
			if t.value < value {
				if t.rightChild != nil {
					return t.rightChild.FindNode(value)
				} else {
					return nil
				}
			}
		}
	}
	return t
}

func main() {
	tree := TreeNode{value: 8}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите элемент или 'q' чтобы завершить ввод")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "q" {
			break
		}
		num, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Некорректный ввод! Повторите попытку.")
			continue
		}
		tree.Insert(int(num))
		fmt.Printf("%v\n", &tree)
	}
	b := tree.Exist(5)
	fmt.Println(b)
}
