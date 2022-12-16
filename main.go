package main

import (
	"fmt"
	"gonote/note"
	"gonote/note/factory"

	// noteutil 是别名
	noteutil "gonote/util"
)

var A = noteutil.F("main.A")

func init() {
	noteutil.F("main.init1")
}

func init() {
	noteutil.F("main.init2")
}

func main() {
	// grammar.go
	note.EscapedCharacters()
	note.VariablesAndConstants()
	fmt.Println(note.Version)
	note.BasicDataTypes()
	note.Pointer()
	note.FmtVerbs()
	note.Operator()
	// note.IfElse()
	// note.SwitchCase()
	note.ForLoop()
	note.LabelAndGoto()
	note.Function()
	note.Defer()
	note.DeferRecover()
	note.Array()
	note.Slice()
	note.Map()
	note.TypeDefinationAndTypeAlias()
	note.Struct()
	note.Method()
	note.Interface()
	note.Goroutine()
	note.Channel()

	// std.go
	note.Random()
	note.StrConv()
	note.PackageStrings()
	note.PackageUtf8()
	note.PackageTime()
	note.FileOperation()
	note.FileReadAndWrite()
	note.Errors()
	note.Log()
	note.CmdArgs()
	note.PackageBuildIn()
	note.PackageSync()
	note.PackageSort()
	note.PackageBinarySearch()

	// algorithm.go
	note.Recursion()
	note.Closure()
	note.Sort()

	//
	m := factory.NewMes()
	m.SetPwd("")
	note.PackageJson()
	// note.TcpServer()
	// note.TcpCli()

	// db.go
	// note.LeveldbBasic()
	// note.LeveldbIterate()
	note.LeveldbTransactionAndSnapshot()
	note.RedisBasic()
	note.RedisPipeline()
	note.RedisTranscation()
	note.RedisIterate()
	note.RedisHashToStruct()
}
