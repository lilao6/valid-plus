package rule

import "testing"

func TestCharacterRule_Valid(t *testing.T) {
	type fields struct {
		character string
		characterType int
	}
	tests := []struct{
		name string
		field fields
		wantErr bool
	}{
		{
			name:"character:标准测试成功",
			field:fields{
				character:"success",
				characterType:lowercase,
			},
			wantErr:false,
		}, {
			name:"character:标准测试失败",
			field:fields{
				character:"success",
				characterType:uppercase,
			},
			wantErr:true,
		},{
			name:"character:标准测试成功",
			field:fields{
				character:"Success",
				characterType:arbitrarily,
			},
			wantErr:false,
		},{
			name:"character:标准测试成功",
			field:fields{
				character:"SUCCESS",
				characterType:uppercase,
			},
			wantErr:false,
		},{
			name:"character:标准测试失败1",
			field:fields{
				character:"SUCCESS",
				characterType:lowercase,
			},
			wantErr:true,
		},
		{
			name:"character:标准测试失败2",
			field:fields{
				character:"success1",
				characterType:arbitrarily,
			},
			wantErr:true,
		},
	}
	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			r := CharacterRule{
				character:tt.field.character,
				characterType:tt.field.characterType,
			}
			if err := r.Valid();(nil != err)!=tt.wantErr{
				t.Errorf("Character.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
