/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import "testing"

func TestPhoneRule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "phone:标准校验成功",
			fields: fields{
				Value: "19941678167",
			},
			wantErr: false,
		},
		{
			name: "phone:标准校验失败",
			fields: fields{
				Value: "1994167816",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PhoneRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("PhoneRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
