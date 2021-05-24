package autostarter

import (
	"os"
	"testing"
)

func TestAutostarter(t *testing.T) {
	exec, _ := os.Executable()
	type args struct {
		sc Shortcut
		ic icon
	}
	test := struct {
		name string
		args args
	}{
		name: "Test",
		args: args{
			sc: Shortcut{
				Name:    "test",
				Exec:    exec,
				Args:    []string{"a", "b", "c"},
				StartIn: os.TempDir(),
			},
			ic: DefaultIcon,
		},
	}
	t.Run(test.name, func(t *testing.T) {
		a := NewAutostart(test.args.sc, test.args.ic)
		if err := a.Enable(); err != nil || !a.IsEnabled() {
			t.Errorf("Error: got: false, want: true, error info: %s", err.Error())
		}
		if err := a.Disable(); err != nil || a.IsEnabled() {
			t.Errorf("Error: got: true, want: false, error info: %s", err.Error())
		}
		if !a.Trigger() {
			t.Error("Error: got: false, want: true")
		}
		if a.Trigger() {
			t.Error("Error: got: true, want: false")
		}
		if got := a.sc.getArgsString(); got != `"a" "b" "c"` {
			t.Errorf("Error: got: %s, want: %s", got, `"a" "b" "c"`)
		}
	})
}
