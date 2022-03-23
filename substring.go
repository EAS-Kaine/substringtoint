package substring

import (
    "fmt"
	s "strings"
	"strconv"
)

func Substring(sub string, str string) int {

	sArr := s.Split(sub, str)

	i, err := strconv.Atoi(s.Join(sArr, ""))
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
	
    return i
}