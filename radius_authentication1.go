package main

import (
	"context"
	"fmt"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"

)
// radtest {username} {password} {hostname} 10 {radius_secret}


func main() {
	radiusurl, radisusername, radispassword,radiussecret :="10.91.170.171:1812","bob","hello","12345"

	//packet := radius.New(radius.CodeAccessRequest, []byte(radiussecret))
	//rfc2865.UserName_SetString(packet, radisusername)
	//rfc2865.UserPassword_SetString(packet, radispassword)
	//response, err := radius.Exchange(context.Background(), packet, radiusurl)
	//if err != nil {
	//	fmt.Println(err)
	//}
	////fmt.Println(response)
	//fmt.Println(response.Code)
	result,err := RadiusAuthActivation(radiusurl, radisusername, radispassword,radiussecret)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(result)

}

func RadiusAuthActivation(radiusurl string, radisusername string, radispassword string,radiussecret string)(string,error){

	packet := radius.New(radius.CodeAccessRequest, []byte(radiussecret))
	rfc2865.UserName_SetString(packet, radisusername)
	rfc2865.UserPassword_SetString(packet, radispassword)
	response, err := radius.Exchange(context.Background(), packet, radiusurl)
	if err != nil {
		return "RadiusError", err
	}
	return fmt.Sprintf("%v",response.Code),nil
}