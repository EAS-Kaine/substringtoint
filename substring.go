package substring

import (
    "fmt"
    "os"
	s "strings"
	"strconv"
)

func Substring() int {

    sub := os.Args[1]
	str := os.Args[2]

	sArr := s.Split(sub, str)

	i, err := strconv.Atoi(s.Join(sArr, ""))
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
	
    return i
}