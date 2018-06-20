/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import "testing"

func TestMinRule_Valid(t *testing.T) {
	type fields struct {
		Value    interface{}
		Min      int64
		UMin     uint64
		MinFloat float64
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
				Min:   10,
			},
			wantErr: false,
		},
		{
			name: "int:标准校验成功",
			fields: fields{
				Value: int(-10),
				Min:   -11,
			},
			wantErr: false,
		},
		{
			name: "int:标准校验失败",
			fields: fields{
				Value: int(10),
				Min:   11,
			},
			wantErr: true,
		},
		{
			name: "int:因为条件值赋错了变量导致的校验条件为零值,造成非预期的成功",
			fields: fields{
				Value: int(10),
				UMin:  11,
			},
			wantErr: false,
		},
		{
			name: "uint:标准校验成功",
			fields: fields{
				Value: uint(10),
				UMin:  10,
			},
			wantErr: false,
		},
		{
			name: "uint:标准校验失败",
			fields: fields{
				Value: uint(10),
				UMin:  11,
			},
			wantErr: true,
		},
		{
			name: "uint:因为条件值赋错了变量导致的校验条件为零值,造成非预期的成功",
			fields: fields{
				Value: uint(10),
				Min:   11,
			},
			wantErr: false,
		},
		{
			name: "float:标准校验成功",
			fields: fields{
				Value: float32(10.5),
				MinFloat:  10.5,
			},
			wantErr: false,
		},
		{
			name: "float:标准校验成功2",
			fields: fields{
				Value: float32(10.9),
				MinFloat:  10.5,
			},
			wantErr: false,
		},
		{
			name: "float:标准校验失败",
			fields: fields{
				Value: float32(10.5),
				MinFloat:  10.6,
			},
			wantErr: true,
		},
		{
			name: "string:标准校验成功",
			fields: fields{
				Value: "aaa",
				Min:   10,
			},
			wantErr: true,
		},
		{
			name: "string:标准校验失败",
			fields: fields{
				Value: "aaa",
				Min:   3,
			},
			wantErr: false,
		},
		{
			name: "string:因为条件值赋错了变量导致的校验条件为零值,造成非预期的成功",
			fields: fields{
				Value: "aaa",
				UMin:  4,
			},
			wantErr: false,
		},
		{
			name: "[]int slice:标准校验成功",
			fields: fields{
				Value: []int{1, 2, 3},
				Min:   3,
			},
			wantErr: false,
		},
		{
			name: "[]string slice:标准校验成功",
			fields: fields{
				Value: []string{"1", "2", "3"},
				Min:   3,
			},
			wantErr: false,
		},
		{
			name: "[]int slice:标准校验失败",
			fields: fields{
				Value: []int{1, 2, 3},
				Min:   4,
			},
			wantErr: true,
		},
		{
			name: "[]int slice:因为条件值赋错了变量导致的校验条件为零值,造成非预期的成功",
			fields: fields{
				Value: []int{1, 2, 3},
				UMin:  4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MinRule{
				value:    tt.fields.Value,
				min:      tt.fields.Min,
				uMin:     tt.fields.UMin,
				minFloat: tt.fields.MinFloat,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("MinRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
