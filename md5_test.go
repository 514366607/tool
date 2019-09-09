package tools

import "testing"

func Test_Md5(t *testing.T) {
	if MD5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
		t.Fatal("md5失败", MD5("123456"))
	}
}
