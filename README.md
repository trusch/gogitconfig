# gogitconfig
a simple "git config" interface for golang with set/get/unset and global support

## Example

```golang
package main

import(
	"github.com/trusch/gogitconfig"
	"log"
)

func main() {
	configValue := gogitconfig.New("test.value")
	v, err := configValue.Get()
	if err != nil {
		log.Println("error in Get:", err.Error())
		err := configValue.Set("foobar")
		if err != nil {
			log.Println("error in Set:", err.Error())
		}
	} else {
		log.Println("value: ", v)
		err = configValue.Unset()
		if err != nil {
			log.Println("error in Unset:", err.Error())
		}
	}
}
```