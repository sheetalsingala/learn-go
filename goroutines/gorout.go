//Race condtion can be detected when multiple func() access the map. Run - go test --race to check. 
// Use channels

package concurrency

type WebsiteChecker func(string) bool
type result struct{
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool{
	results:=make(map[string]bool)
	resultC:=make(chan result)

	for _, url:=range urls{
		go func(u string){
			resultC<-result{u,wc(u)}
		}(url)
	}

	for i:=0;i<len(urls);i++{
		result:=<-resultC
		results[result.string]=result.bool
	}

	return results
}
