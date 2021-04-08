package dialect

import (
	"testing"
)

func Test_logEntry_String(t *testing.T) {
	tests := []struct {
		entry LogEntry
		want  string
	}{
		{
			entry: LogEntry{
				Action: DriverActionExec,
				Query:  "SELECT * FROM nothing WHERE id = ?",
				Args:   []interface{}{1},
			},
			want: "driver.Exec: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: LogEntry{
				Action: DriverActionQuery,
				Query:  "SELECT * FROM nothing WHERE id = ?",
				Args:   []interface{}{1},
			},
			want: "driver.Query: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: LogEntry{
				Action: DriverActionTx,
				TxID:   "tx-id",
			},
			want: "driver.Tx(tx-id): started",
		},
		{
			entry: LogEntry{
				Action: DriverActionBeginTx,
				TxID:   "tx-id",
			},
			want: "driver.BeginTx(tx-id): started",
		},
		{
			entry: LogEntry{
				Action:   DriverActionTx,
				TxAction: TxActionExec,
				TxID:     "tx-id",
				Query:    "SELECT * FROM nothing WHERE id = ?",
				Args:     []interface{}{1},
			},
			want: "Tx(tx-id).Exec: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: LogEntry{
				Action:   DriverActionTx,
				TxAction: TxActionQuery,
				TxID:     "tx-id",
				Query:    "SELECT * FROM nothing WHERE id = ?",
				Args:     []interface{}{1},
			},
			want: "Tx(tx-id).Query: query=SELECT * FROM nothing WHERE id = ? args=[1]",
		},
		{
			entry: LogEntry{
				Action:   DriverActionTx,
				TxAction: TxActionCommit,
				TxID:     "tx-id",
			},
			want: "Tx(tx-id): committed",
		},
		{
			entry: LogEntry{
				Action:   DriverActionTx,
				TxAction: TxActionRollback,
				TxID:     "tx-id",
			},
			want: "Tx(tx-id): rollbacked",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.entry.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
