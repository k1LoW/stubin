package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		args       []string
		envs       map[string]string
		wantStatus int
		wantStdout string
		wantStderr string
	}{
		{
			[]string{"stubin", "test"},
			map[string]string{},
			0, "", "",
		},
		{
			[]string{"stubin", "test"},
			map[string]string{
				"STUBIN_STATUS": "1",
				"STUBIN_STDOUT": "out\n",
				"STUBIN_STDERR": "err",
			},
			1, "out\n", "err",
		},
		{
			[]string{"stubin", "-h"},
			map[string]string{
				"STUBIN_STATUS":        "1",
				"STUBIN_STDOUT":        "out\n",
				"STUBIN_STDERR":        "err",
				"STUBIN_EXPECT_COND":   "\"-h\" in args",
				"STUBIN_EXPECT_STATUS": "0",
				"STUBIN_EXPECT_STDOUT": "help\n",
				"STUBIN_EXPECT_STDERR": "",
			},
			0, "help\n", "",
		},
		{
			[]string{"stubin", "test"},
			map[string]string{
				"STUBIN_STATUS":        "1",
				"STUBIN_STDOUT":        "out\n",
				"STUBIN_STDERR":        "err",
				"STUBIN_EXPECT_COND":   "\"-h\" in args",
				"STUBIN_EXPECT_STATUS": "0",
				"STUBIN_EXPECT_STDOUT": "help\n",
				"STUBIN_EXPECT_STDERR": "",
			},
			1, "out\n", "err",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)
			for k, v := range tt.envs {
				t.Setenv(k, v)
			}
			got := Run(tt.args, stdout, stderr)
			if got != tt.wantStatus {
				t.Errorf("got %v, want %v", got, tt.wantStatus)
			}
			if stdout.String() != tt.wantStdout {
				t.Errorf("got STDOUT %v, want %v", stdout.String(), tt.wantStdout)
			}
			if stderr.String() != tt.wantStderr {
				t.Errorf("got STDERR %v, want %v", stderr.String(), tt.wantStderr)
			}
		})
	}
}
