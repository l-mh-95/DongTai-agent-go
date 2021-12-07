package fmtSprintf

import (
	"go-agent/global"
	"go-agent/model/request"
	"go-agent/utils"
)

func Sprintf(format string, a ...interface{}) string {
	var sourceHash global.HashKeys = []string{}
	var sourceBytes = []byte("")
	for _, v := range a {
		switch v.(type) {
		case string:
			sourceHash = append(sourceHash, utils.GetSource(v))
			space := []byte(" ")
			str1 := append([]byte(v.(string)), space...)
			sourceBytes = append(sourceBytes, str1...)
		}
	}
	s := SprintfT(format, a...)
	var targetHash global.HashKeys = []string{utils.GetSource(s)}
	var pool = request.Pool{
		SourceValues: string(sourceBytes),
		SourceHash:   sourceHash,
		TargetHash:   targetHash,
	}
	poolTree := request.PoolTree{
		Begin:       false,
		Pool:        &pool,
		Children:    []*request.PoolTree{},
		GoroutineID: utils.CatGoroutineID(),
	}
	for k, _ := range global.PoolTreeMap {
		if k.Some(sourceHash) {
			global.PoolTreeMap[k].Children = append(global.PoolTreeMap[k].Children, &poolTree)
			break
		}
	}
	global.PoolTreeMap[&targetHash] = &poolTree
	return s
}

func SprintfT(format string, a ...interface{}) string {
	return ""
}
