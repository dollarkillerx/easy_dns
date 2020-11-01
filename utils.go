package easy_dns

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func LookupTXT(domain string, dnsIP string) (*Message, error) {
	domain = domain + "."
	var m Message
	m.Header.ID = uint16(Random(1000, 65534))
	m.Questions = append(m.Questions, Question{
		Name:  MustNewName(domain),
		Type:  TypeTXT,
		Class: ClassINET,
	})

	pack, err := m.Pack()
	if err != nil {
		log.Fatalln(err)
	}

	return dial(pack, dnsIP)
}

func LookupNS(domain string, dnsIP string) (*Message, error) {
	domain = domain + "."
	var m Message
	m.Header.ID = uint16(Random(1000, 65534))
	m.Questions = append(m.Questions, Question{
		Name:  MustNewName(domain),
		Type:  TypeNS,
		Class: ClassINET,
	})

	pack, err := m.Pack()
	if err != nil {
		log.Fatalln(err)
	}

	return dial(pack, dnsIP)
}

func LookupMX(domain string, dnsIP string) (*Message, error) {
	domain = domain + "."
	var m Message
	m.Header.ID = uint16(Random(1000, 65534))
	m.Questions = append(m.Questions, Question{
		Name:  MustNewName(domain),
		Type:  TypeMX,
		Class: ClassINET,
	})

	pack, err := m.Pack()
	if err != nil {
		log.Fatalln(err)
	}

	return dial(pack, dnsIP)
}

func LookupCNAME(domain string, dnsIP string) (*Message, error) {
	domain = domain + "."
	var m Message
	m.Header.ID = uint16(Random(1000, 65534))
	m.Questions = append(m.Questions, Question{
		Name:  MustNewName(domain),
		Type:  TypeCNAME,
		Class: ClassINET,
	})

	pack, err := m.Pack()
	if err != nil {
		log.Fatalln(err)
	}

	return dial(pack, dnsIP)
}

func LookupIP(domain string, dnsIP string) (*Message, error) {
	domain = domain + "."
	var m Message
	m.Header.ID = uint16(Random(1000, 65534))
	m.Questions = append(m.Questions, Question{
		Name:  MustNewName(domain),
		Type:  TypeA,
		Class: ClassINET,
	})

	pack, err := m.Pack()
	if err != nil {
		log.Fatalln(err)
	}

	return dial(pack, dnsIP)
}

func LookupIPSimple(domain string, dnsIP string) (string, error) {
	ip, err := LookupIP(domain, dnsIP)
	if err != nil {
		return "", err
	}
	var result string
	for _, v := range ip.Answers {
		if v.Header.Type == TypeA {
			return v.Body.GoString(), nil
		}
	}

	return result, fmt.Errorf("not fund")
}

func dial(msg []byte, dns string) (*Message, error) {
	conn, err := net.DialTimeout("udp", dns, time.Second)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(time.Second))
	conn.SetReadDeadline(time.Now().Add(time.Second))
	conn.SetDeadline(time.Now().Add(time.Second))

	if _, err := conn.Write(msg); err != nil {
		return nil, err
	}

	buf := make([]byte, 512)
	var m Message
	read, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	if err := m.Unpack(buf[:read]); err != nil {
		return nil, err
	}

	return &m, nil
}
