// @Author: Perry
// @Date  : 2020/10/11
// @Desc  : json相关工具

package common

import "encoding/json"

func IndentJson(obj interface{}) []byte {
	ret, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return []byte(err.Error())
	}
	return ret
}
