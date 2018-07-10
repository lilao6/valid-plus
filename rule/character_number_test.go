package rule

import "testing"

type fields struct {
	character     string
	characterType int
}

func TestCharacterNumberRule_Valid(t *testing.T) {
	tests := []struct {
		name    string
		field   fields
		wantErr bool
	}{
		{
			name: "character:标准测试成功",
			field: fields{
				character:     "success1",
				characterType: lowercase,
			},
			wantErr: false,
		}, {
			name: "character:标准测试失败",
			field: fields{
				character:     "success",
				characterType: uppercase,
			},
			wantErr: true,
		}, {
			name: "character:标准测试成功",
			field: fields{
				character:     "Success1",
				characterType: arbitrarily,
			},
			wantErr: false,
		}, {
			name: "character:标准测试成功",
			field: fields{
				character:     "SUCCESS1",
				characterType: uppercase,
			},
			wantErr: false,
		}, {
			name: "character:标准测试失败1",
			field: fields{
				character:     "SUCCESS",
				characterType: lowercase,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CharacterNumberRule{
				value:     tt.field.character,
				valueType: tt.field.characterType,
			}
			if err := r.Valid(); (nil != err) != tt.wantErr {
				t.Errorf("CharacterNumber.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}