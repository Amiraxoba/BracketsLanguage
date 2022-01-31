package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var arr [30000]byte
var curr = 0

func main() {
	args := os.Args[1:]
	scriptContent, _ := ioutil.ReadFile(args[1])
	script := strings.ReplaceAll(string(scriptContent), "\n", "")
	if len(script)%3 != 0 {
		panic("Wrong amount of characters in file")
	}
	if args[0] == "run" {
		a := make([]string, len(script)/3)
		for i := 0; i < len(script)/3; i++ {
			a[i] = script[3*i : i*3+3]
		}
		for _, cmd := range a {
			command := Command{
				[]rune(cmd)[0],
				[]rune(cmd)[1],
				[]rune(cmd)[2],
			}
			command.Execute()
		}
	} else if args[0] == "build" {
		Compile(script)
	}
}

type Command struct {
	Command  rune
	Argument rune
	Details  rune
}

func (cmd Command) Execute() {
	switch cmd.Command {
	case '>':
		if cmd.Argument == '(' && cmd.Details == ')' {
			arr[curr] += 1
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			arr[curr] += 10
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			arr[curr] += 50
		}
		break
	case '<':
		if cmd.Argument == '(' && cmd.Details == ')' {
			arr[curr] -= 1
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			arr[curr] -= 10
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			arr[curr] -= 50
		}
		break
	case '[':
		if cmd.Argument == '(' && cmd.Details == ')' {
			curr += 1
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			curr += 10
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			curr += 50
		}
		break
	case ']':
		if cmd.Argument == '(' && cmd.Details == ')' {
			curr -= 1
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			curr -= 10
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			curr -= 50
		}
		break
	case '(':
		if cmd.Argument == '(' && cmd.Details == '(' {
			fmt.Print(string(arr[curr]))
		} else if cmd.Argument == ')' && cmd.Details == ')' {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadByte()
			arr[curr] = input
		}
		break
	}
}
func Compile(script string) {
	ret := ""
	s := script
	a := make([]string, len(s)/3)
	for i := 0; i < len(s)/3; i++ {
		a[i] = s[3*i : i*3+3]
	}
	for _, cmd := range a {
		command := Command{
			[]rune(cmd)[0],
			[]rune(cmd)[1],
			[]rune(cmd)[2],
		}
		ret += command.Compile()
	}
	fmt.Println(ret)
}

func (cmd Command) Compile() string {
	switch cmd.Command {
	case '>':
		if cmd.Argument == '(' && cmd.Details == ')' {
			return "+"
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			return "++++++++++"
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			return "++++++++++++++++++++++++++++++++++++++++++++++++++"
		}
		break
	case '<':
		if cmd.Argument == '(' && cmd.Details == ')' {
			return "-"
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			return "----------"
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			return "--------------------------------------------------"
		}
		break
	case '[':
		if cmd.Argument == '(' && cmd.Details == ')' {
			return ">"
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			return ">>>>>>>>>>"
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			return ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
		}
		break
	case ']':
		if cmd.Argument == '(' && cmd.Details == ')' {
			return "<"
		} else if cmd.Argument == '(' && cmd.Details == '(' {
			return "<<<<<<<<<<"
		} else if cmd.Argument == ')' && cmd.Details == '(' {
			return "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
		}
		break
	case '*':
		if cmd.Argument == '*' && cmd.Details == '*' {
			return "."
		}
		break

	case '(':
		if cmd.Argument == '(' && cmd.Details == '(' {
			return "."
		} else if cmd.Argument == ')' && cmd.Details == ')' {
			return ","
		}
		break
	}
	return ""
}
