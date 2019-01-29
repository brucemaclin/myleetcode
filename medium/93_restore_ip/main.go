package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {
		return nil
	}
	var result []string
	doRestoreIP(0, "", s, &result)
	return result
}

func doRestoreIP(count int, tempAddress string, src string, result *[]string) {
	if count == 4 || len(src) == 0 {
		if count == 4 && len(src) == 0 {
			*result = append(*result, tempAddress)
		}
		return
	}
	runes := []rune(src)
	for i := 0; i < len(runes) && i <= 2; i++ {
		if i != 0 && runes[0] == '0' {
			break
		}
		part := string(runes[0 : i+1])
		num, err := strconv.Atoi(part)
		if err != nil {
			return
		}
		if num <= 255 {
			if tempAddress != "" {
				part = "." + part
			}
			tempAddress += part
			doRestoreIP(count+1, tempAddress, string(runes[i+1:]), result)
			tempAddress = string(tempAddress[:len(tempAddress)-len(part)])
		}
	}
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
