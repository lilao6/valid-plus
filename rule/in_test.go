/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import "testing"

func TestInRule_Valid(t *testing.T) {
	type fields struct {
		Value        interface{}
		IntValues    []int64
		UintValues   []uint64
		FloatValues  []float64
		StringValues []string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "int:标准校验成功",
			fields: fields{
				Value:     int8(2),
				IntValues: []int64{1, 2, 3, 4},
			},
			wantErr: false,
		},
		{
			name: "int:标准校验失败",
			fields: fields{
				Value:     int8(2),
				IntValues: []int64{1, 6, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "uint:标准校验成功",
			fields: fields{
				Value:      uint8(2),
				UintValues: []uint64{1, 2, 3, 4},
			},
			wantErr: false,
		},
		{
			name: "uint:标准校验失败",
			fields: fields{
				Value:      uint8(2),
				UintValues: []uint64{1, 6, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "float:标准校验成功",
			fields: fields{
				Value:       float32(1.77767777777),
				FloatValues: []float64{1.77767777777, 2, 3, 4},
			},
			wantErr: false,
		},
		{
			name: "float:标准校验失败",
			fields: fields{
				Value:       float32(1.77767777777),
				FloatValues: []float64{1.777, 2, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "string:标准校验成功",
			fields: fields{
				Value:        "bbb",
				StringValues: []string{"aaa", "bbb", "ccc", "ddd"},
			},
			wantErr: false,
		},
		{
			name: "string:标准校验失败",
			fields: fields{
				Value:        "bb",
				StringValues: []string{"aaa", "bbb", "ccc", "ddd"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InRule{
				value:        tt.fields.Value,
				intValues:    tt.fields.IntValues,
				uintValues:   tt.fields.UintValues,
				floatValues:  tt.fields.FloatValues,
				stringValues: tt.fields.StringValues,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("InRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
