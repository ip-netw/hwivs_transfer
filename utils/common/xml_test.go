// @Author : DAIPENGYUAN
// @File : xml_test
// @Time : 2020/10/28 14:46 
// @Description : 

package common

import (
	"fmt"
	"testing"
)

func TestIsValidXML(t *testing.T) {
	testv:=`<?xml version="1.0" encoding="UTF-8"?>
<rpc message-id="1024" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <get>
    <filter type="subtree">
      <devm xmlns="httputil://www.huawei.com/netconf/vrp/huawei-devm">
        <globalPara>
        </globalPara>
      </devm>
    </filter>
  </get>
</rpc>`
	fmt.Println(IsValidXML([]byte(testv)))
}
