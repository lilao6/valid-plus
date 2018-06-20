/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import (
	"crypto/rand"
	"net"
	"io"
	"fmt"
	"strconv"
)

// intAll2int int8,16,32,64转int64转之前要确定类型一定是其中的一个
func intAll2int(value interface{}) int64 {
	switch value.(type) {
	case int:
		return int64(value.(int))
	case int8:
		return int64(value.(int8))
	case int16:
		return int64(value.(int16))
	case int32:
		return int64(value.(int32))
	case int64:
		return value.(int64)
	default:
		panic("The value not in (int,int8,int16,int32,int64)")
	}
}

// uintAll2uint uint8,16,32,64转uint64转之前要确定类型一定是其中的一个
func uintAll2uint(value interface{}) uint64 {
	switch value.(type) {
	case uint:
		return uint64(value.(uint))
	case uint8:
		return uint64(value.(uint8))
	case uint16:
		return uint64(value.(uint16))
	case uint32:
		return uint64(value.(uint32))
	case uint64:
		return value.(uint64)
	default:
		panic("The value not in (uint,uint8,uint16,uint32,uint64)")
	}
}

// floatAll2float float32,64转float64转之前要确定类型一定是其中的一个
func floatAll2float(value interface{}) float64 {
	switch value.(type) {
	case float32:
		return float64(value.(float32))
	case float64:
		return value.(float64)
	default:
		panic("The value not in (float32,float64)")
	}
}

// generateIPv6 随机生成一个ipv6地址
//func generateIPv6(local bool) (net.IP, error) {
//	addr := make(net.IP, 16)
//	_, err := io.ReadFull(rand.Reader, addr[:])
//	if err != nil {
//		return nil, fmt.Errorf("failed to generate IPv6: %s", err)
//	}
//
//	addr[0] = 0xfd
//	addr[1] = 0x4c
//	addr[2] = 0xbd
//	addr[3] = 0x56
//	addr[4] = 0x5c
//	addr[5] = 0xee
//
//	if local {
//		addr[6] = 0x80
//		addr[7] = 0x00
//	} else {
//		addr[6] = 0x00
//		addr[7] = 0x00
//	}
//
//	return addr, nil
//}

// generateIPv4 随机生成一个ipv4地址
func generateIPv4() (net.IP, error) {
	addr := make(net.IP, 4)
	_, err := io.ReadFull(rand.Reader, addr[:])
	if err != nil {
		return nil, fmt.Errorf("failed to generate IPv4: %s", err)
	}

	addr[0] = 172
	addr[1] = 18
	if addr[3] == 0 {
		addr[3] = addr[3] + 1
	}
	if addr[3] == 255 {
		addr[3] = addr[3] - 1
	}
	return addr, nil
}

// strings2Int64s string切片转int64切片
func strings2Int64s(src []string) ([]int64, error) {
	var dst = make([]int64, len(src))
	for i, v := range src {
		temp, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		dst[i] = temp
	}
	return dst, nil
}

// strings2Uint64s string切片转uint64切片
func strings2Uint64s(src []string) ([]uint64, error) {
	var dst = make([]uint64, len(src))
	for i, v := range src {
		temp, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, err
		}
		dst[i] = temp
	}
	return dst, nil
}

// strings2float64s string切片转float64切片
func strings2float64s(src []string) ([]float64, error) {
	var dst = make([]float64, len(src))
	for i, v := range src {
		temp, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		dst[i] = temp
	}
	return dst, nil
}
