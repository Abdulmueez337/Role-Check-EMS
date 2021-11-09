package service

import (
	"fmt"
	"strings"
)

func CheckRole(designation string,role string)(int,string){
	CheckDesignation := strings.ToLower(designation)
	CheckApiName := strings.ToLower(role)
	roles := []string{
		"adddetails","deletedetails","showdetailsadmin","showdetailsemployeeself","showdetailsteamlead","updatedetails",
	}
	if CheckDesignation=="admin"{
		for _,v := range roles{
			if v==CheckApiName{
				return 200,"Roles found successfully"
			}
		}
	}else if CheckDesignation=="teamlead"{
		if roles[3]== CheckApiName || roles[4]== CheckApiName{
			fmt.Println("tealead")
			return 200,"Roles found successfully"
		}else if roles[0]== CheckApiName || roles[1]== CheckApiName || roles[2]== CheckApiName || roles[5] == CheckApiName{
			fmt.Println("tealead1")
			return 403,"Forbidden"
		}else{
			fmt.Println("tealead3")
			return 400,"Invalid apiName"
		}
	}else if CheckDesignation=="employee"{
		if roles[3]== CheckApiName {
			return 200,"Role found successfully"
		}else if roles[0]== CheckApiName || roles[1]== CheckApiName || roles[2]== CheckApiName || roles[4]== CheckApiName || roles[5]== CheckApiName{
			return 403,"Forbidden"
		}else {
			return 400,"Invalid apiName"
		}
	}else {
		return 400, "Bad Request"
	}
	return 400, "Bad Request"
}


