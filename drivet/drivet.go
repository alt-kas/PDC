package drivet

import (
    "fmt"
    "syscall"
    "unsafe"
)

var (
    kernel32              = syscall.NewLazyDLL("kernel32.dll")
    getVolumeInformationW = kernel32.NewProc("GetVolumeInformationW")
)

func GetDriveLabel(driveLetter string) string {
    driveLetter = driveLetter + "\\"

    // Convert drive letter to a UTF16 string
    drivePath, err := syscall.UTF16PtrFromString(driveLetter)
    if err != nil {
        fmt.Println("Error converting drive path to UTF16:", err)
        return "ERR-CVT-DP-UTF16"
    }

    // Prepare buffers for receiving volume information
    var volumeNameBuffer [256]uint16
    var volumeSerialNumber uint32
    var fileSystemFlags uint32
    var fileSystemNameBuffer [256]uint16

    // Call GetVolumeInformationW function
    _, _, err = getVolumeInformationW.Call(
        uintptr(unsafe.Pointer(drivePath)),
        uintptr(unsafe.Pointer(&volumeNameBuffer[0])),
        uintptr(len(volumeNameBuffer)),
        uintptr(unsafe.Pointer(&volumeSerialNumber)),
        0,
        uintptr(unsafe.Pointer(&fileSystemFlags)),
        uintptr(unsafe.Pointer(&fileSystemNameBuffer[0])),
        uintptr(len(fileSystemNameBuffer)),
    )

    if err != nil && err.Error() != "The operation completed successfully." {
        fmt.Println("Error getting volume information:", err)
        return "ERR-GET-VOL-INFO"
    }

    // Convert volume name from UTF16 to string
    volumeName := syscall.UTF16ToString(volumeNameBuffer[:])

    return volumeName
}

