/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

type FullTag struct {
	fullTag string
}

func (ft *FullTag) SetFullTag(tag string) {
	ft.fullTag = tag
}
func (ft *FullTag) GetFullTag() string {
	return ft.fullTag
}
