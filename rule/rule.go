/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

type Rule interface {
	Tag() string
	Generate(value interface{}, tagValue string) error
	Valid() error
	SetFullTag(tag string)
	GetFullTag() string
	Clone() Rule
}
