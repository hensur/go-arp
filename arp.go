package arp

/*
#include <string.h>
#include <net/if_arp.h>
*/
import "C"

import (
	"golang.org/x/sys/unix"
	"net"
	"syscall"
	"unsafe"
)

const (
	ATF_COM  C.int = 0x02
	ATF_PERM C.int = 0x04
	ATF_PUBL C.int = 0x08
)

func arpSyscall(call uintptr, interfaceName string, address net.IPNet) error {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, unix.IPPROTO_IP)
	if err != nil {
		return err
	}
	var flags C.int
	flags = ATF_PERM | ATF_COM
	flags |= ATF_PUBL
	arr := [16]C.char{}
	arr2 := C.CString(interfaceName)
	C.strcpy((*C.char)(&arr[0]), (*C.char)(arr2))

	var ip [4]byte
	for i := 0; i < 4; i++ {
		ip[i] = []byte(address.IP)[i]
	}

	var raw syscall.RawSockaddrInet4
	raw.Family = unix.AF_INET
	copy(raw.Addr[:], address.IP.To4())

	req := C.struct_arpreq{
		arp_pa:    *(*C.struct_sockaddr)(unsafe.Pointer(&raw)),
		arp_dev:   arr,
		arp_flags: flags,
	}

	_, _, err = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), call, uintptr(unsafe.Pointer(&req)))
	if err.(syscall.Errno) != 0 {
		return err
	}
	return nil
}

func SetARP(interfaceName string, address net.IPNet) error {
	return arpSyscall(syscall.SIOCSARP, interfaceName, address)
}

func DeleteARP(interfaceName string, address net.IPNet) error {
	return arpSyscall(syscall.SIOCDARP, interfaceName, address)
}
