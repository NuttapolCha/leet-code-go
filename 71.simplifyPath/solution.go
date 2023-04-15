package solution

import (
	"strings"
)

type stack struct {
	a []string
}

func (st *stack) push(x string) {
	st.a = append(st.a, x)
}

func (st *stack) peek() string {
	return st.a[len(st.a)-1]
}

func (st *stack) pop() string {
	ret := st.peek()
	st.a = st.a[:len(st.a)-1]
	return ret
}

func (st *stack) isEmpty() bool {
	return len(st.a) == 0
}

func (st *stack) constructAbsolutePath() string {
	paths := strings.Join(st.a, "/")
	return "/" + paths
}

func simplifyPath(path string) string {
	// ensure slash at ending
	path += "/"

	st := stack{}

	chars := strings.Split(path, "")
	dir := ""
	for _, char := range chars {
		switch {
		case char == "/":
			switch dir {
			case ".", "":
				// do nothing
			case "..":
				if !st.isEmpty() {
					st.pop()
				}
			default:
				st.push(dir)
			}
			dir = ""

		default:
			dir += char
		}
	}

	return st.constructAbsolutePath()
}
