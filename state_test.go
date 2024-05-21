package overflow

import (
	"reflect"
	"testing"

	"github.com/onflow/cadence"
	"github.com/stretchr/testify/require"
)

func TestOverflowState_parseArguments(t *testing.T) {
	o, err := OverflowTesting()
	require.NoError(t, err)
	require.NotNil(t, o)

	type args struct {
		fileName  string
		code      []byte
		inputArgs map[string]interface{}
	}

	tests := []struct {
		name    string
		args    args
		want    []cadence.Value
		want1   CadenceArguments
		wantErr bool
	}{
		{
			name: "transaction(u:UInt64, i:Int64) valid args",
			args: args{
				fileName:  "somefile",
				code:      []byte("transaction(u:UInt64, i:Int64){}"),
				inputArgs: map[string]interface{}{"u": 42, "i": 42},
			},
			want:    []cadence.Value{cadence.NewUInt64(42), cadence.NewInt64(42)},
			want1:   map[string]cadence.Value{"u": cadence.NewUInt64(42), "i": cadence.NewInt64(42)},
			wantErr: false,
		},
		// {
		// 	name: "transaction(u:UInt64, c:Int64){} invalid args",
		// 	args: args{
		// 		fileName:  "somefile",
		// 		code:      []byte("transaction(u:UInt64, i:Int64){}"),
		// 		inputArgs: map[string]interface{}{"u": 42, "i": "something else"},
		// 	},
		// 	want:    nil,
		// 	want1:   nil,
		// 	wantErr: true,
		// },
		{
			name: "transaction(adr:Address)",
			args: args{
				fileName:  "somefile",
				code:      []byte("transaction(adr:Address){}"),
				inputArgs: map[string]interface{}{"adr": "first"},
			},
			want:    []cadence.Value{cadence.BytesToAddress(o.FlowAddress("first").Bytes())},
			want1:   map[string]cadence.Value{"adr": cadence.BytesToAddress(o.FlowAddress("first").Bytes())},
			wantErr: false,
		},
		// {
		// 	name: "transaction(arr:[Address])",
		// 	args: args{
		// 		fileName:  "somefile",
		// 		code:      []byte("transaction(arr:[Address]){}"),
		// 		inputArgs: map[string]interface{}{"arr": []string{"first"}},
		// 	},
		// 	want:    []cadence.Value{cadence.BytesToAddress(o.FlowAddress("first").Bytes())},
		// 	want1:   map[string]cadence.Value{"adr": cadence.BytesToAddress(o.FlowAddress("first").Bytes())},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := o.parseArguments(tt.args.fileName, tt.args.code, tt.args.inputArgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("OverflowState.parseArguments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OverflowState.parseArguments() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OverflowState.parseArguments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
