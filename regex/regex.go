package main

import(
	"fmt"
	"regexp"
	"strings"
)

func main(){

	str1 := "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.4.3/openshift-client-linux-4.4.3.tar.gz"
	re := regexp.MustCompile(`\d+\.?\d+\.\d*`)
//	fmt.Printf("String contains any match: %v\n", re.MatchString(str1)) // True
	submatchall := re.FindStringSubmatch(str1)
	versions := strings.Split(submatchall[0], ".")	
	fmt.Println(versions[0]+"."+ versions[1])
}
