package rule

import "testing"

func TestAlphaDashRule_Valid(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "alphaDash:标准测试成功",
			fields: fields{
				value: "1234aB-_",
			},
			wantErr: true,
		}, {
			name: "alphaDash:标准测试失败",
			fields: fields{
				value: "a,1-@ $",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := AlphaDashRule{
				value: tt.fields.value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("AlphaDash.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
