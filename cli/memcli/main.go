package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tkuchiki/memcli"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("memcli", "memcached cli tool")
	servers := app.Flag("servers", "servers").Default("127.0.0.1:11211").Strings()

	get := app.Command("get", "get")
	getKey := get.Flag("key", "key").Required().String()

	set := app.Command("set", "set")
	setKey := set.Flag("key", "key").Required().String()
	setFlags := set.Flag("flags", "flags").Uint32()
	setExpiration := set.Flag("expire", "expiration").Int32()

	delete := app.Command("delete", "delete")
	deleteKey := delete.Flag("key", "key").Required().String()

	deleteAll := app.Command("delete-all", "delete-all")
	flushAll := app.Command("flush-all", "flush-all")

	app.Version("0.1.0")
	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	client := memcli.NewClient(*servers...)
	defer client.Close()

	switch command {
	case set.FullCommand():
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		client.Set(*setKey, b, *setFlags, *setExpiration)
	case get.FullCommand():
		item, err := client.Get(*getKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(item.Value))
	case delete.FullCommand():
		err := client.Delete(*deleteKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf(`Deleted %s`, *deleteKey))
	case deleteAll.FullCommand():
		err := client.DeleteAll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(`Deleted all items`)
	case flushAll.FullCommand():
		err := client.FlushAll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(`Flushed all items`)
	}
}
