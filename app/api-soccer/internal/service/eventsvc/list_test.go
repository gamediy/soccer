package eventsvc

import (
	"context"
	"reflect"
	"testing"
)

func Test_eventsList_Exec(t *testing.T) {
	type fields struct {
		Status int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []EventsListOutput
		wantErr bool
	}{
		{
			args: args{
				ctx: context.TODO(),
			},
			fields: fields{
				Status: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &eventsList{
				Status: tt.fields.Status,
			}
			got, err := this.Exec(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() got = %v, want %v", got, tt.want)
			}
		})
	}
}
