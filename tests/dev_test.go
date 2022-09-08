package dev

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
)

func TestIPV4RegMatch(t *testing.T) {
	str := "Aug 27 09:05:59 pppd[26882]: local  IP address 110.165.101.102"
	regStr := `local  IP address (?P<ipv4>((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?))`
	reg := regexp.MustCompile(regStr)
	match := reg.FindStringSubmatch(str)
	groupNames := reg.SubexpNames()
	result := make(map[string]string)
	// convert to map
	for i, name := range groupNames {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	prettyResult, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("%s\n", prettyResult)
	fmt.Printf("%v,%v", match[0], groupNames[0])
}

func TestRegNameGroup(t *testing.T) {

	str := `Alice 20 alice@gmail.com`
	// 使用命名分组，显得更清晰
	re := regexp.MustCompile(`(?P<name>[a-zA-Z]+)\s+(?P<age>\d+)\s+(?P<email>\w+@\w+(?:\.\w+)+)`)
	match := re.FindStringSubmatch(str) //匹配的值的数组
	groupNames := re.SubexpNames()      //匹配的名称的数组
	fmt.Printf("%v, %v, %d, %d\n", match, groupNames, len(match), len(groupNames))
	result := make(map[string]string)
	// 转换为map
	for i, name := range groupNames {
		if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
			result[name] = match[i]
		}
	}
	prettyResult, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("%s\n", prettyResult)

}
