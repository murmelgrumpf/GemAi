package configs

import (
	"fmt"

	tiny "github.com/Yiwen-Chan/tinydb"
)

type param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var botParams string = "db/botParams.json"

func SetBotParam(name string, value string) {
	fmt.Println(name + " - " + value)
	storage, err := tiny.JSONStorage(botParams)
	if err != nil {
		panic(err)
	}
	defer storage.Close()

	database, err := tiny.TinyDB(tiny.CachingMiddleware(storage, 1))
	if err != nil {
		panic(err)
	}

	table := tiny.GetTable[param](database)

	par, err := table.Select(func(p param) bool { return p.Name == name })
	if err != nil {
		panic(err)
	}

	if len(par) > 0 {
		err = table.Update(func(p param) param { p.Value = value; return p }, func(p param) bool { return p.Name == name })
		if err != nil {
			panic(err)
		}
		return
	}

	err = table.Insert(param{name, value})
	if err != nil {
		panic(err)
	}

}

func GetBotParam(name string) string {
	storage, err := tiny.JSONStorage(botParams)
	if err != nil {
		panic(err)
	}
	defer storage.Close()

	database, err := tiny.TinyDB(tiny.CachingMiddleware(storage, 1))
	if err != nil {
		panic(err)
	}

	table := tiny.GetTable[param](database)

	par, err := table.Select(func(p param) bool { return p.Name == name })
	if err != nil {
		panic(err)
	}

	return par[0].Value
}
