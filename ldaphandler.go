package fmhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"os/exec"
	"bytes"
)

type LDAP struct {

	Url     string
	Username string
	Commonname string
	Domaincomponentfirst string
	Domaincomponentsecond string
	Activated string



}

func LDAPHandler(w http.ResponseWriter, r *http.Request) {

	var res map[string]interface{}
	json.NewDecoder(r.Body).Decode(&res)
	urlpath := "ldap://"+fmt.Sprintf("%v", res["url"])+":389"
	ldapdata:= LDAP{
		Url:urlpath,
		Username:fmt.Sprintf("%v", res["username"]),
		Commonname:fmt.Sprintf("%v", res["cm"]),
		Domaincomponentfirst:fmt.Sprintf("%v", res["dc1"]),
		Domaincomponentsecond:fmt.Sprintf("%v", res["dc2"]),
		Activated:"Sucssefuly",

	}
	data, _ := json.Marshal(ldapdata)
	path := "./configuration/authentication/ldap_config.json"
	query := r.URL.Query()
	os.Setenv("SEARCHNAME",query.Get("SearchName"))
	cmd := exec.Command("./.internal/create_objects_zip.sh")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		fmt.Fprintf(w, "%s", "{\"Message\":\"The file does not exist\"}");
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
		_ = ioutil.WriteFile(path, string(data), 0644)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", "{\"Message\":\"Successfully Activated\"}");
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		fmt.Fprintf(w, "%s", "Not Saved")
	}

}

