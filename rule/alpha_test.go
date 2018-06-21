package rule

import "testing"

func TestAlphaRule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "alpha:标准测试成功",
			fields: fields{
				Value: "aaabbb",
			},
			wantErr: false,
		}, {
			name: "alpha:标准测试失败",
			fields: fields{
				Value: "aaabbb1",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := AlphaRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Alpha.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
