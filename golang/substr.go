package main

import "fmt"
import "sort"

func makeLookUp(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		k := string(v)
		if e, ok := m[k]; ok {
			e++
			m[k] = e
		} else {
			m[k] = 1
		}
	}
	return m
}

type subStr []string

func (s subStr) Len() int           { return len(s) }
func (s subStr) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s subStr) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

func MinWindowSubstring(strArr []string) string {
	if len(strArr) < 2 {
		return ""
	}
	first := strArr[0]
	second := strArr[1]
	substr := make([]string, 0)
	for i := 0; i < len(first); i++ {
		count := 0
		m := makeLookUp(second)
		for j := i; j < len(first); j++ {
			k := string(first[j])
			if v, ok := m[k]; ok && v > 0 {
				count++
				v--
				m[k] = v
			}
			if count == len(second) {
				substr = append(substr, first[i:j+1])
				break
			}
		}
	}
	sort.Sort(subStr(substr))
	//for _, v := range substr {
	//	fmt.Println("#############", v)
	//}
	return substr[0]

}

func main() {
	v := [...]string{"ahffaksfajeeubsne", "jefaa"}
	fmt.Println(MinWindowSubstring(v[:]))
}
