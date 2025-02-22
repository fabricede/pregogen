package mix3types

import (
	"encoding/json"
	"testing"
	"time"
)

func TestYourType2_UnmarshalJSON_pathfinder(t *testing.T) {
	type fields struct {
		RuneField  rune
		RunesField []rune
		DateField  time.Time
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test0",
			fields: fields{
				RuneField:  'b',
				RunesField: []rune("123"),
				DateField:  testTime,
			},
			args: args{
				data: []byte(`{"runefield":"b","runesfield":"123","datefield":"2021-08-01T00:00:00.000000000Z"}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &YourType2{
				RuneField:  tt.fields.RuneField,
				RunesField: tt.fields.RunesField,
				DateField:  tt.fields.DateField,
			}
			if err := y.UnmarshalJSON_pathfinder(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("YourType2.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkYourType2_UnmarshalJSON_pathfinder(t *testing.B) {
	type fields struct {
		RuneField  rune
		RunesField []rune
		DateField  time.Time
	}
	type args struct {
		data []byte
	}
	benchmarks := []struct {
		name   string
		fields fields
		args   args
	}{

		// TODO: Add test cases.
		{
			name: "Test0",
			fields: fields{
				RuneField:  'b',
				RunesField: []rune("123"),
				DateField:  testTime,
			},
			args: args{
				data: []byte(`{"runefield":98,"runesfield":[94,95,96],"datefield":"2021-08-01T00:00:00.000000000Z"}`),
			},
		},
	}
	for _, bb := range benchmarks {
		t.Run(bb.name+"UnmarshalJSON_pathfinder", func(t *testing.B) {
			y := &YourType2{}
			for i := 0; i < t.N; i++ {
				if err := y.UnmarshalJSON_pathfinder(bb.args.data); err != nil {
					t.Errorf("YourType2.UnmarshalJSON() error = %v", err)
				}
			}
		})
		t.Run(bb.name+"json.Unmarshal", func(t *testing.B) {
			yRef := &YourType2Ref{}
			for i := 0; i < t.N; i++ {
				if err := json.Unmarshal(bb.args.data, yRef); err != nil {
					t.Errorf("json.Unmarshal() error = %v", err)
				}
			}
		})
	}
}
