/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import "testing"

func TestMaxRule_Valid(t *testing.T) {
	type fields struct {
		Value    interface{}
		Max      int64
		UMax     uint64
		MaxFloat float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "int:标准校验成功",
			fields: fields{
				Value: int(10),
				Max:   10,
			},
			wantErr: false,
		},
		{
			name: "int:标准校验成功",
			fields: fields{
				Value: int(-10),
				Max:   -9,
			},
			wantErr: false,
		},
		{
			name: "int:标准校验失败",
			fields: fields{
				Value: int(10),
				Max:   9,
			},
			wantErr: true,
		},
		{
			name: "int:因为条件值赋错了变量导致的校验条件为零值,造成非预期的失败",
			fields: fields{
				Value: int(10),
				UMax:  11,
			},
			wantErr: true,
		},
		{
			name: "uint:标准校验成功",
			fields: fields{
				Value: uint(10),
				UMax:  10,
			},
			wantErr: false,
		},
		{
			name: "uint:标准校验失败",
			fields: fields{
				Value: uint(10),
				UMax:  9,
			},
			wantErr: true,
		},
		{
			name: "uint:因为条件值赋错了变量导致的校验条件为零值,造成非预期的失败",
			fields: fields{
				Value: uint(10),
				Max:   11,
			},
			wantErr: true,
		},
		{
			name: "float:标准校验成功",
			fields: fields{
				Value:    float32(10.7),
				MaxFloat: 10.9,
			},
			wantErr: false,
		},
		{
			name: "float:标准校验成功2",
			fields: fields{
				Value:    float32(10.7),
				MaxFloat: 10.7,
			},
			wantErr: false,
		},
		{
			name: "float:标准校验失败",
			fields: fields{
				Value:    float32(10.6),
				MaxFloat: 10.5,
			},
			wantErr: true,
		},
		{
			name: "string:标准校验成功",
			fields: fields{
				Value: "aaa",
				Max:   3,
			},
			wantErr: false,
		},
		{
			name: "string:标准校验失败",
			fields: fields{
				Value: "aaa",
				Max:   2,
			},
			wantErr: true,
		},
		{
			name: "string:因为条件值赋错了变量导致的校验条件为零值,造成非预期的失败",
			fields: fields{
				Value: "aaa",
				UMax:  4,
			},
			wantErr: true,
		},
		{
			name: "[]int slice:标准校验成功",
			fields: fields{
				Value: []int{1, 2, 3},
				Max:   3,
			},
			wantErr: false,
		},
		{
			name: "[]string slice:标准校验成功",
			fields: fields{
				Value: []string{"1", "2", "3"},
				Max:   3,
			},
			wantErr: false,
		},
		{
			name: "[]int slice:标准校验失败",
			fields: fields{
				Value: []int{1, 2, 3},
				Max:   2,
			},
			wantErr: true,
		},
		{
			name: "[]int slice:因为条件值赋错了变量导致的校验条件为零值,造成非预期的失败",
			fields: fields{
				Value: []int{1, 2, 3},
				UMax:  4,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MaxRule{
				value:    tt.fields.Value,
				max:      tt.fields.Max,
				uMax:     tt.fields.UMax,
				maxFloat: tt.fields.MaxFloat,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("MaxRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
