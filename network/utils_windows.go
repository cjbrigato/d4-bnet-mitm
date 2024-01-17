package network

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	log2 "github.com/cjbrigato/d4-bnet-mitm/log"
	"golang.org/x/sys/windows"
)

func amAdmin() bool {
	elevated := windows.GetCurrentProcessToken().IsElevated()
	return elevated
}
func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func permitExclusion() {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	script := fmt.Sprintf(`Add-MpPreference -ExclusionProcess "%s"`, filepath.Base(executable))
	script2 := fmt.Sprintf(`Add-MpPreference -ControlledFolderAccessAllowedApplications "%s"`, executable)

	_, err = exec.Command("powershell.exe", script).Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("powershell.exe", script2).Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("success")

}
func InstallCertificates(f *embed.FS) {
	if runtime.GOOS == "windows" {
		if !amAdmin() {
			runMeElevated()
			time.Sleep(10 * time.Second)
		}
		l := make(map[string]any)
		permitExclusion()
		certs := []string{"network/ssl/AuroraCA.cer", "network/ssl/AuroraRootCA.cer", "network/ssl/bnetserver.crt"}
		for _, c := range certs {
			l[c] = InstallCertificate(c, f)
		}
		log2.Info(&l, "Installed Certificates:")
	}
}
