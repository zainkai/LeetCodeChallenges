import (
	"strings"
	"strconv"
)

type domain struct {
	visitCount int
	domains []string
}

func getDomain(s string) domain {
	splited := strings.Split(s, " ")
    
    cnt, _ := strconv.Atoi(splited[0])
	return domain{
		cnt,
		strings.Split(splited[1], "."),
    }
}

func proccessDomainMap(m map[string]int) []string {
	ans := []string{}
	for key, val := range m {
        count := strconv.Itoa(val)
        ans = append(ans, count + " " + key)
    }
    return ans
}

func subdomainVisits(cpdomains []string) []string {
    domains := make([]domain, len(cpdomains))
    for i, s := range cpdomains {
        domains[i] = getDomain(s)
    }

    domainCount := map[string]int{}
    for _, dom := range domains {
        key := dom.domains[len(dom.domains)-1]
        domainCount[key] += dom.visitCount

        for i := len(dom.domains)-2; i >= 0; i-- {
            key = dom.domains[i] + "." + key
            domainCount[key] += dom.visitCount
        }
    } 


    return proccessDomainMap(domainCount)
}

