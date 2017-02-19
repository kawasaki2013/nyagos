package completion

import (
	"errors"
	"fmt"
	"strings"

	"../lua"
	"../readline"
)

var Hook lua.Pushable = lua.TNil{}

func luaHook(this *readline.Buffer, rv *CompletionList) (*CompletionList, error) {
	L, L_ok := this.Context.Value("lua").(lua.Lua)
	if !L_ok {
		return rv, errors.New("listUpComplete: could not get lua instance")
	}

	L.Push(Hook)
	if L.IsFunction(-1) {
		L.NewTable()
		L.PushString(rv.RawWord)
		L.SetField(-2, "rawword")
		L.Push(rv.Pos + 1)
		L.SetField(-2, "pos")
		L.PushString(rv.AllLine)
		L.SetField(-2, "text")
		L.PushString(rv.Word)
		L.SetField(-2, "word")
		L.NewTable()
		for key, val := range rv.List {
			L.Push(1 + key)
			L.PushString(val)
			L.SetTable(-3)
		}
		L.SetField(-2, "list")
		if err := L.Call(1, 1); err != nil {
			fmt.Println(err)
		}
		if L.IsTable(-1) {
			list := make([]string, 0, len(rv.List)+32)
			wordUpr := strings.ToUpper(rv.Word)
			for i := 1; true; i++ {
				L.Push(i)
				L.GetTable(-2)
				str, strErr := L.ToString(-1)
				L.Pop(1)
				if strErr != nil || str == "" {
					break
				}
				strUpr := strings.ToUpper(str)
				if strings.HasPrefix(strUpr, wordUpr) {
					list = append(list, str)
				}
			}
			if len(list) > 0 {
				rv.List = list
			}
		}
	}
	L.Pop(1) // remove something not function or result-table
	return rv, nil
}
