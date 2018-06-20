/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import "testing"

func TestNumberRule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Number:标准校验成功",
			fields: fields{
				Value: "0123",
			},
			wantErr: false,
		},
		{
			name: "Number:标准校验成功",
			fields: fields{
				Value: "68843",
			},
			wantErr: false,
		},
		{
			name: "Number:标准校验失败",
			fields: fields{
				Value: "pp",
			},
			wantErr: true,
		},
		{
			name: "Number:标准校验失败",
			fields: fields{
				Value: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NumberRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("NumberRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
