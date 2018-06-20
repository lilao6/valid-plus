/*
 * Copyright 2018 The golyu Authors. All rights reserved.
 * Use of this source code is governed by a Apache License, Version 2.0 (the "License");
 * license that can be found in the LICENSE file.
 */

package rule

import "testing"

func TestIpRule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}
	type test struct {
		name    string
		fields  fields
		wantErr bool
	}
	tests := make([]test, 0)
	for i := 0; i < 100; i++ {
		ip, _ := generateIPv4()

		temp := test{
			name: "ip:标准校验成功:" + ip.String(),
			fields: fields{
				Value: ip.String(),
			},
		}
		tests = append(tests, temp)
	}

	tests2 := []test{
		{
			name: "ip:标准校验失败",
			fields: fields{
				Value: "0.103.222.333",
			},
			wantErr: true,
		},
		{
			name: "ip:标准校验失败",
			fields: fields{
				Value: "202.222.333",
			},
			wantErr: true,
		},
	}
	tests = append(tests, tests2...)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &IpRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("IpRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
