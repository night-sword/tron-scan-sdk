package tronscan

import (
	"encoding/json"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/structs"
	"github.com/night-sword/kratos-kit/errors"
)

func toStringMap(in any) (out map[string]string) {
	out = make(map[string]string)

	sm, _ := structs.ToMap(in)
	for s, v := range sm {
		out[s] = convertor.ToString(v)
	}
	return
}

func jsonDecode[T any](in []byte) (out *T, err error) {
	err = json.Unmarshal(in, &out)
	if err != nil {
		err = errors.BadRequest("decode rsp fail", "").WithCause(err)
	}

	return
}
