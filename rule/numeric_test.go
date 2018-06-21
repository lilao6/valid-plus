package rule

import "testing"

func TestNumericRule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "nimeric:标准测试成功",
			fields: fields{
				Value: "111",
			},
			wantErr: false,
		}, {
			name: "nimeric:标准测试失败",
			fields: fields{
				Value: "111a",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NumericRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Numeric.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
