package configs

import (
	tiny "github.com/Yiwen-Chan/tinydb"
)

type botParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var botParams string = "db/botParams.json"

func SetBotParam(name string, value string) {
	storage, err := tiny.JSONStorage(botParams)
	if err != nil {
		panic(err)
	}
	defer storage.Close()

	database, err := tiny.TinyDB(tiny.CachingMiddleware(storage, 1))
	if err != nil {
		panic(err)
	}

	table := tiny.GetTable[botParam](database)

	par, err := table.Select(func(p botParam) bool { return p.Name == name })
	if err != nil {
		panic(err)
	}

	if len(par) > 0 {
		err = table.Update(
			func(p botParam) botParam { p.Value = value; return p },
			func(p botParam) bool { return p.Name == name },
		)
		if err != nil {
			panic(err)
		}
		return
	}

	err = table.Insert(botParam{name, value})
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

	table := tiny.GetTable[botParam](database)

	par, err := table.Select(func(p botParam) bool { return p.Name == name })
	if err != nil {
		panic(err)
	}

	if len(par) > 0 {
		return par[0].Value
	}
	return ""
}
