package rule

import "testing"

func TestLengthRule_Valid(t *testing.T) {
	type fields struct {
		value  interface{}
		length int
	}

	test := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "string:标准检验成功",
			fields: fields{
				value:  "这是一条测试数据",
				length: 8,
			},
			wantErr: false,
		},
		{
			name: "string:标准检验失败",
			fields: fields{
				value:  "这是一条测试数据",
				length: 9,
			},
			wantErr: true,
		},
		{
			name: "slice:标准检验成功",
			fields: fields{
				value:  []int{1, 2, 3},
				length: 3,
			},
			wantErr: false,
		},
		{
			name: "slice:标准检验失败",
			fields: fields{
				value:  []int{1, 2, 3, 4},
				length: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			r := &LengthRule{
				value:  tt.fields.value,
				length: tt.fields.length,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Length.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
