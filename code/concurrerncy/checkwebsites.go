package code

import "time"

type Websitechecker func(string) bool

func checkwebsites(ws Websitechecker, urls []string) map[string]bool {
	website := make(map[string]bool)

	for _, v := range urls {
		go func(s string) {
			website[s] = ws(s)
		}(v)
	}

	time.Sleep(3 * time.Second)

	return website
}
