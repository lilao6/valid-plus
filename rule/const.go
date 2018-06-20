/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

const (
	ACCURACY     = 0.00001 // 浮点数比较大小的精度
	SPLIT_SEP    = ","     // 用来分隔参数
	SPLIT_SEP_OR = " "     // 用来分隔函数参数,or功能专用
	TAG_PRE      = "("     // 标记,例 Max(
	TAG_PRE_OR   = "<"     // 标记,例 Or<
)
