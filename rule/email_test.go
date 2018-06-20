/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import "testing"

func TestEmailRule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "email:标准校验成功1",
			fields: fields{
				Value: "xx22323243@qq.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验成功2",
			fields: fields{
				Value: "22323243@qq.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验成功3",
			fields: fields{
				Value: "43@qq.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验成功4",
			fields: fields{
				Value: "243@qq.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验成功5",
			fields: fields{
				Value: "243@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验成功6",
			fields: fields{
				Value: "243@163.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验成功7",
			fields: fields{
				Value: "243@edu.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验成功8",
			fields: fields{
				Value: "243@foxmail.com",
			},
			wantErr: false,
		},
		{
			name: "email:标准校验失败1",
			fields: fields{
				Value: "243foxmail.com",
			},
			wantErr: true,
		},
		{
			name: "email:标准校验失败2",
			fields: fields{
				Value: "@243.com",
			},
			wantErr: true,
		},
		{
			name: "email:标准校验失败3",
			fields: fields{
				Value: "243.com@",
			},
			wantErr: true,
		},
		{
			name: "email:标准校验失败4",
			fields: fields{
				Value: "@24@3.com",
			},
			wantErr: true,
		},
		{
			name: "email:标准校验失败5",
			fields: fields{
				Value: ".@243.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmailRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("EmailRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
