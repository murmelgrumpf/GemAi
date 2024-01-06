package configs

import (
	tiny "github.com/Yiwen-Chan/tinydb"
)

type guildParam struct {
	GuildId  string          `json:"guildId"`
	Features map[string]bool `json:"features"`
}

var guildParams string = "db/guildParams.json"

func SetGuildParam(guildId string, features map[string]bool) {
	storage, err := tiny.JSONStorage(guildParams)
	if err != nil {
		panic(err)
	}
	defer storage.Close()

	database, err := tiny.TinyDB(tiny.CachingMiddleware(storage, 1))
	if err != nil {
		panic(err)
	}

	table := tiny.GetTable[guildParam](database)

	par, err := table.Select(func(p guildParam) bool { return p.GuildId == guildId })
	if err != nil {
		panic(err)
	}

	if len(par) > 0 {
		err = table.Update(
			func(p guildParam) guildParam { p.Features = features; return p },
			func(p guildParam) bool { return p.GuildId == guildId },
		)
		if err != nil {
			panic(err)
		}
		return
	}

	err = table.Insert(guildParam{guildId, features})
	if err != nil {
		panic(err)
	}

}

func GetGuildParam(guildId string) map[string]bool {
	storage, err := tiny.JSONStorage(guildParams)
	if err != nil {
		panic(err)
	}
	defer storage.Close()

	database, err := tiny.TinyDB(tiny.CachingMiddleware(storage, 1))
	if err != nil {
		panic(err)
	}

	table := tiny.GetTable[guildParam](database)

	par, err := table.Select(func(p guildParam) bool { return p.GuildId == guildId })
	if err != nil {
		panic(err)
	}

	if len(par) > 0 {
		return par[0].Features
	}
	return nil
}
