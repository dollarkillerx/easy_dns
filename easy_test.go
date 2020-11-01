package easy_dns

import (
	"log"
	"testing"
)

func TestEasyDnsIP(t *testing.T) {
	ip, err := LookupIP("baidu.com", "223.5.5.5:53")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ip.Answers[0].Body.GoString())
}

func TestEasyDnsIP2(t *testing.T) {
	simple, err := LookupIPSimple("baidu.com", "223.5.5.5:53")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(simple)
}

func TestEasyCNAME(t *testing.T) {
	cname, err := LookupCNAME("www.baidu.com", "223.5.5.5:53")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(cname)
	if len(cname.Answers) != 0 {
		log.Println(cname.Answers[0].Body.GoString())
	}
}
