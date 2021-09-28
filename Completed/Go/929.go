import "strings"

type Email struct {
    local string
    domain string
}

func numUniqueEmails(emails []string) int {
    emailSet := map[Email]bool{}
    
    for _, email := range emails {
        key := GetEmailObj(email)
        emailSet[key] = true
    }
    
    return len(emailSet)
}

func GetEmailObj(email string) Email {
    tokenized := strings.Split(email, "@")
    local := tokenized[0]
    
    // Rule: remove anything after +
    local = strings.Split(local, "+")[0]
    // Rule: ignore .
    local = strings.ReplaceAll(local, ".", "")
    
    domain := tokenized[1]
    
    return Email{
        local,
        domain,
    }
}
