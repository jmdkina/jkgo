package jksys

import (
	"bufio"
	"errors"
	"io"
	"jk/jklog"
	"os"
	"strings"
)

func IfDomainValid(domain string) bool {
	if strings.Index(domain, ".") <= 0 {
		return false
	}
	return true
}

func FilterDomainOfIptables(domain string) error {
	if !IfDomainValid(domain) {
		jklog.Lfile().Warnf("domain: %s is invalid\n", domain)
		return errors.New("invalid domain")
	}
	strargs := "-I INPUT -m string --string " + domain + " --algo bm -j DROP"
	jklog.Lfile().Debugf("Insert: %s\n", strargs)
	systemargs("iptables", strargs)
	return nil
}

func FilterDomainRemove(domain string) error {
	if !IfDomainValid(domain) {
		jklog.Lfile().Warnf("domain: %s is invalid\n", domain)
		return errors.New("invalid domain")
	}
	strargs := "-D INPUT -m string --string " + domain + " --algo bm -j DROP"
	jklog.Lfile().Debugf("Remove: %s\n", strargs)
	systemargs("iptables", strargs)
	return nil
}

func SysFilterDomainFromFiles(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return errors.New("filename not exist.")
	}
	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		FilterDomainOfIptables(line)
	}
	return nil
}

var domainLists []string

func SysFilterDomainWriteToFile(filename string, data []string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	for _, v := range data {
		f.WriteString(v+"\n")
	}
	f.Close()
	return nil
}

func SysFilterDomainDiff(filename string, data []string) bool {
	if len(domainLists) == 0 {
		f, err := os.Open(filename)
		if err == nil {
			buf := bufio.NewReader(f)
			for {
				line, err := buf.ReadString('\n')
				line = strings.TrimSpace(line)
				if err != nil {
					if err == io.EOF {
						break
					}
					break
				}
				domainLists = append(domainLists, line)
			}
		}
	}

	dls := strings.Join(domainLists, "")
	ndls := strings.Join(data, "")

	if strings.Compare(dls, ndls) == 0 {
		jklog.Lfile().Infoln("domain lists same, needn't do")
		return false
	}

	for _, v := range domainLists {
		FilterDomainRemove(v)
	}
	domainLists = data
	SysFilterDomainWriteToFile(filename, data)
	return true
}
