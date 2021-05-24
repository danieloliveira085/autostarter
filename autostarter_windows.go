// +build windows

package autostarter

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

const shortcutExt = ".lnk"

type oleObject struct {
	oleShellObject *ole.IUnknown
	wShell         *ole.IDispatch
}

func newShellObject() *oleObject {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		oleShellObject.Release()
		panic(err)
	}
	wShell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		oleShellObject.Release()
		wShell.Release()
		panic(err)
	}
	return &oleObject{
		oleShellObject: oleShellObject,
		wShell:         wShell,
	}
}

func (o *oleObject) release() {
	o.oleShellObject.Release()
	o.wShell.Release()
}

func getStartupDir() (dir string) {
	dir = filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	return
}

func checkIcon(path string) error {
	switch filepath.Ext(path) {
	case ".ico":
	default:
		return errors.New("format invalid: " + filepath.Ext(path))
	}
	return nil
}

func createShortcut(sc Shortcut, ic icon) (path string, err error) {
	object := newShellObject()
	defer object.release()
	path = filepath.Join(getStartupDir(), sc.Name+shortcutExt)
	cs, err := oleutil.CallMethod(object.wShell, "CreateShortcut", path)
	if err != nil {
		return
	}
	iDispatch := cs.ToIDispatch()
	oleutil.PutProperty(iDispatch, "TargetPath", sc.Exec)
	if len(sc.Args) > 0 {
		oleutil.PutProperty(iDispatch, "Arguments", sc.getArgsString())
	}
	oleutil.PutProperty(iDispatch, "WorkingDirectory", sc.StartIn)
	if ic != "" {
		oleutil.PutProperty(iDispatch, "IconLocation", string(ic))
	}
	_, err = oleutil.CallMethod(iDispatch, "Save")
	if err != nil {
		return
	}
	return
}
