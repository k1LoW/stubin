package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/antonmedv/expr"
)

const (
	expectCondPrefix    = "EXPECT_COND"
	expectStatusPrefix  = "EXPECT_STATUS"
	expectStdoutPrefix  = "EXPECT_STDOUT"
	expectStderrPrefix  = "EXPECT_STDERR"
	defaultStatusPrefix = "STATUS"
	defaultStdoutPrefix = "STDOUT"
	defaultStderrPrefix = "STDERR"
)

func Run(args []string, stdout, stderr io.Writer) int {
	bin := filepath.Base(args[0])
	args = args[1:]
	prefix := fmt.Sprintf("%s_", strings.ToUpper(bin))

	var expects []string
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if !strings.HasPrefix(pair[0], prefix) {
			continue
		}
		if strings.HasPrefix(pair[0], prefix+expectCondPrefix) {
			expects = append(expects, pair[0])
		}
	}
	sort.Strings(expects)
	for _, e := range expects {
		cond := os.Getenv(e)
		v, err := expr.Eval(cond, map[string]any{"args": args})
		if err != nil {
			_, _ = fmt.Fprintln(stderr, err)
			return 1
		}
		tf, ok := v.(bool)
		if !ok {
			_, _ = fmt.Fprintf(stderr, "invalid condition: %s\n", cond)
			return 1
		}
		if !tf {
			continue
		}
		suffix := strings.TrimPrefix(e, prefix+expectCondPrefix)
		_, _ = fmt.Fprint(stdout, os.Getenv(prefix+expectStdoutPrefix+suffix))
		_, _ = fmt.Fprint(stderr, os.Getenv(prefix+expectStderrPrefix+suffix))
		s, _ := strconv.Atoi(os.Getenv(prefix + expectStatusPrefix + suffix))
		return s
	}

	_, _ = fmt.Fprint(stdout, os.Getenv(prefix+defaultStdoutPrefix))
	_, _ = fmt.Fprint(stderr, os.Getenv(prefix+defaultStderrPrefix))
	s, _ := strconv.Atoi(os.Getenv(prefix + defaultStatusPrefix))
	return s
}
