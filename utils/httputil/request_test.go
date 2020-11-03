// @Author : DAIPENGYUAN
// @File : request_test
// @Time : 2020/10/21 17:54 
// @Description : 

package httputil

import "testing"

func TestReq(t *testing.T) {
	ctt := `{
        "packets_recv": 1,
        "packets_sent": 1,
        "packet_loss": 0,
        "remote_ip_addr": "112.80.248.76",
        "remote_addr": "www.baidu.com",
        "rtts": [
                "24.9331ms"
        ],
        "min_rtt": "24.9331ms",
        "max_rtt": "24.9331ms",
        "avg_rtt": "24.9331ms",
        "std_dev_rtt": "0s"}`
	c := &Client{
		URI:    "httputil://localhost:8080/common",
		Method: "POST",
	}
	c.ContentType = TJSON
	if c.Headers == nil {
		c.Headers = make(map[string]string)
	}
	c.Headers["ID"] = "c03bc457-1399-11eb-9bd6-f875a41832dc"
	b, err := c.RAW([]byte(ctt))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
