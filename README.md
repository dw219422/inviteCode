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
	var uid int32 = 123456789
	inviteCode := icd.Create(uid)
	fmt.Println("uid:", uid, "--->", inviteCode)
	fmt.Println("inviteCode:", inviteCode, "--->", icd.Decode(inviteCode))
}

output:
uid: 123456789 ---> JG4RDNB
inviteCode: JG4RDNB ---> 123456789
```
