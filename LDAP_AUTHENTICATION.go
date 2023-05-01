package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strings"
	"os"
)

type Env struct {
	l *ldap.Conn
	err *ldap.Error
}

func main() {

	ldapurl := os.Getenv("HOST")
	serviceusername := os.Getenv("SERVICEUSERNAME")
	servicepassword := os.Getenv("SERVICEPASSWORD")
	bindstring := os.Getenv("BINDSTRING")
	loginusername := os.Getenv("LOGINUSERNAME")
	loginpassword := os.Getenv("LOGINPASSWORD")





	ldapresult, err := ldapauthentication(ldapurl, serviceusername, servicepassword, bindstring, loginusername, loginpassword)
	if err != nil {
		fmt.Println("Invalid credentials")
	} else {
		for _, entry := range ldapresult.Entries {
			fmt.Println(entry.DN)
			fmt.Println("Sucessfuly Logged in")
		}
	}

}

func ldapauthentication(ldapurl string, serviceusername string, servicepassword string, bindstring string, loginusername string, loginpassword string) (*ldap.SearchResult, error) {


	if strings.Contains(ldapurl, "636") {
		l, err := ldap.DialURL(fmt.Sprintf("ldaps://%s", ldapurl), ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
		if err != nil {
			return nil, err
		}
		l.Bind(serviceusername, servicepassword)

		searchReq := ldap.NewSearchRequest(
			bindstring,
			ldap.ScopeWholeSubtree, // you can also use ldap.ScopeWholeSubtree
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(loginusername)),
			[]string{},
			nil,
		)
		result, err := l.Search(searchReq)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Search Error: %s", err)
		}

		if len(result.Entries) > 0 {
			//fmt.Println(result)

			for _, entry := range result.Entries {
				fmt.Println(entry.DN)
				fmt.Println(loginpassword)

				err = l.Bind(entry.DN, loginpassword)
				if err != nil {
					fmt.Println("usererror ", err)
				}
			}

			return result, nil

		} else {
			return nil, fmt.Errorf("Couldn't fetch search entries")
		}

	} else {

		l, err := ldap.DialURL(fmt.Sprintf("ldap://%s", ldapurl))
		if err != nil {
			return nil, err
		}
		l.Bind(serviceusername, servicepassword)

		searchReq := ldap.NewSearchRequest(
			bindstring,
			ldap.ScopeWholeSubtree, // you can also use ldap.ScopeWholeSubtree
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(loginusername)),
			[]string{},
			nil,
		)
		result, err := l.Search(searchReq)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Search Error: %s", err)
		}

		if len(result.Entries) > 0 {
			//fmt.Println(result)

			for _, entry := range result.Entries {
				fmt.Println(entry.DN)
				err = l.Bind(entry.DN, loginpassword)
				if err != nil {
					fmt.Println("usererror ", err)
				}
			}

			return result, nil

		} else {
			return nil, fmt.Errorf("Couldn't fetch search entries")
		}

	}


}
