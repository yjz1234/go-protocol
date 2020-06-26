package protocol

import (
	"strconv"
	"strings"
)

func ParserNetMask(maskip string) (string, string) {
	ip_mask := strings.Split(maskip, "/")
	ip, mask := ip_mask[0], ip_mask[1]
	maskInt, _ := strconv.Atoi(mask)
	ipNum := ipstringToNum(ip)
	lowerMask := (32 - maskInt)
	minIp := ipNum &^ ((1 << lowerMask) - 1) //设置低位为0
	maxIp := minIp + (1<<lowerMask - 1)      //直接计算出最大ip的数字值
	ipMin := numToIpstring(minIp)
	ipMax := numToIpstring(maxIp)
	return ipMin, ipMax
}

func numToIpstring(num int) string {
	ipsegs := make([]string, 4)
	for i := 0; i < 4; i++ {
		ipsegs[i] = strconv.Itoa(num >> (24 - (i * 8)))
		num = num &^ (0xff << (24 - (i * 8))) //每轮设置对应的高八位为0
	}
	ip := strings.Join(ipsegs, ".")
	return ip
}

func ipstringToNum(ip string) int {
	ipsegs := strings.Split(ip, ".")
	ipNum := 0
	for i := 0; i < 4; i++ {
		ipseg, _ := strconv.Atoi(ipsegs[i])
		ipNum += (ipseg << (24 - i*8)) //每轮加上八位ip地址的数字值
	}
	return ipNum
}
