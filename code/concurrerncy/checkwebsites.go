package code

type Websitechecker func(string) bool

type results struct {
	s string
	b bool
}

func checkwebsites(ws Websitechecker, urls []string) map[string]bool {
	website := make(map[string]bool)
	ch := make(chan results)

	for _, v := range urls {
		go func(s string) {
			ch <- results{s, ws(s)}
		}(v)
	}

	for i := 0; i < len(urls); i++ {
		c := <-ch
		website[c.s] = c.b

	}

	return website
}
