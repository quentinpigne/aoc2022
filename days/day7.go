package days

import (
	"fmt"
	"quentinpigne/aoc2022/utils"
	"strings"
)

var pointer *utils.TreeNode
var root = new(utils.TreeNode)

// Parts Computation

func ComputeDay7Part1(fileLines []string) int {
	for _, line := range fileLines {
		if isCommand(line) {
			executeCommand(line)
		} else {
			addDirOrFile(line)
		}
	}
	sum := 0
	computeSum(root, &sum)
	return sum
}

func ComputeDay7Part2(fileLines []string) int {
	for _, line := range fileLines {
		if isCommand(line) {
			executeCommand(line)
		} else {
			addDirOrFile(line)
		}
	}
	neededSpace := 30000000 - computeFreeSpace()
	var dir *utils.TreeNode
	findSmallestDirWithMinSize(root, neededSpace, dir)
	return dir.GetSize()
}

// Utilities Functions

func isCommand(line string) bool {
	return strings.HasPrefix(line, "$")
}

func executeCommand(line string) {
	parts := strings.Split(line, " ")
	if parts[1] == "cd" {
		changeDirectory(parts[2])
	}
}

func changeDirectory(directory string) {
	switch directory {
	case "/":
		pointer = root
		break
	case "..":
		pointer = pointer.GetParent()
		break
	default:
		pointer = pointer.FindChild(directory)
		break
	}
}

func addDirOrFile(line string) {
	parts := strings.Split(line, " ")
	if parts[0] == "dir" {
		pointer.AddDirectory(parts[1])
	} else {
		pointer.AddFile(parts[1], utils.ToInt(parts[0]))
	}
}

func computeFreeSpace() int {
	return 70000000 - root.GetSize()
}

func printNode(node *utils.TreeNode, tab int) {
	for i := 0; i < tab; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("- %s (%s, size=%d)\n", node.GetName(), node.GetType(), node.GetSize())
	for _, childNode := range node.GetChildren() {
		printNode(childNode, tab+2)
	}
}

func computeSum(node *utils.TreeNode, sum *int) {
	if node.GetType() == "dir" && node.GetSize() <= 100000 {
		*sum += node.GetSize()
	}
	for _, childNode := range node.GetChildren() {
		computeSum(childNode, sum)
	}
}

func findSmallestDirWithMinSize(node *utils.TreeNode, size int, current *utils.TreeNode) {
	if current == nil || (node.GetType() == "dir" && node.GetSize() < current.GetSize() && node.GetSize() > size) {
		current = node
	}
	for _, childNode := range node.GetChildren() {
		findSmallestDirWithMinSize(childNode, size, current)
	}
}
