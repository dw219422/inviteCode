# inviteCode
Generating Invitation Code Based on UID,require UID must greater than 0

# example
```
package main

import (
	"fmt"

	icd "github.com/dw219422/inviteCode"
)

func main() {
	var uid uint64 = 123456789
	inviteCode, err := icd.Create(uid)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("uid:", uid, "--->", inviteCode, err)
	fmt.Println("inviteCode:", inviteCode, "--->", icd.Decode(inviteCode))
}

output:
uid: 123456789 ---> JG4RDNB
inviteCode: JG4RDNB ---> 123456789
```
