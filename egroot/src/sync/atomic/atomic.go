// Package atomic provides low-level atomic memory primitives.
//
// These functions works in relaxed memory model so they don't provide any implicit
// memory barriers nor "happens-before" edges for different gorutines.
package atomic

import "unsafe"

func compareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func compareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func compareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
func compareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)

func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool) {
	return compareAndSwapInt32(addr, old, new)
}
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool) {
	return compareAndSwapUint32(addr, old, new)
}
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool) {
	return compareAndSwapUintptr(addr, old, new)
}
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) {
	return compareAndSwapPointer(addr, old, new)
}

func addInt32(addr *int32, delta int32) (new int32)
func addUint32(addr *uint32, delta uint32) (new uint32)
func addUintptr(addr *uintptr, delta uintptr) (new uintptr)

func AddInt32(addr *int32, delta int32) (new int32) {
	return addInt32(addr, delta)
}
func AddUint32(addr *uint32, delta uint32) (new uint32) {
	return addUint32(addr, delta)
}
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) {
	return addUintptr(addr, delta)
}

func orInt32(addr *int32, mask int32) (new int32)
func orUint32(addr *uint32, mask uint32) (new uint32)
func orUintptr(addr *uintptr, mask uintptr) (new uintptr)

func OrInt32(addr *int32, mask int32) (new int32) {
	return orInt32(addr, mask)
}
func OrUint32(addr *uint32, mask uint32) (new uint32) {
	return orUint32(addr, mask)
}
func OrUintptr(addr *uintptr, mask uintptr) (new uintptr) {
	return orUintptr(addr, mask)
}

func andInt32(addr *int32, mask int32) (new int32)
func andUint32(addr *uint32, mask uint32) (new uint32)
func andUintptr(addr *uintptr, mask uintptr) (new uintptr)

func AndInt32(addr *int32, mask int32) (new int32) {
	return andInt32(addr, mask)
}
func AndUint32(addr *uint32, mask uint32) (new uint32) {
	return andUint32(addr, mask)
}
func AndUintptr(addr *uintptr, mask uintptr) (new uintptr) {
	return andUintptr(addr, mask)
}

func xorInt32(addr *int32, mask int32) (new int32)
func xorUint32(addr *uint32, mask uint32) (new uint32)
func xorUintptr(addr *uintptr, mask uintptr) (new uintptr)

func XorInt32(addr *int32, mask int32) (new int32) {
	return xorInt32(addr, mask)
}
func XorUint32(addr *uint32, mask uint32) (new uint32) {
	return xorUint32(addr, mask)
}
func XorUintptr(addr *uintptr, mask uintptr) (new uintptr) {
	return xorUintptr(addr, mask)
}

func swapInt32(addr *int32, new int32) (old int32)
func swapUint32(addr *uint32, new uint32) (old uint32)
func swapUintptr(addr *uintptr, new uintptr) (old uintptr)
func swapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)

func SwapInt32(addr *int32, new int32) (old int32) {
	return swapInt32(addr, new)
}
func SwapUint32(addr *uint32, new uint32) (old uint32) {
	return swapUint32(addr, new)
}
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr) {
	return swapUintptr(addr, new)
}
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) {
	return swapPointer(addr, new)
}

func loadInt32(addr *int32) (val int32)
func loadUint32(addr *uint32) (val uint32)
func loadUintptr(addr *uintptr) (val uintptr)
func loadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)

func LoadInt32(addr *int32) (val int32) {
	return loadInt32(addr)
}
func LoadUint32(addr *uint32) (val uint32) {
	return loadUint32(addr)
}
func LoadUintptr(addr *uintptr) (val uintptr) {
	return loadUintptr(addr)
}
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) {
	return loadPointer(addr)
}

func storeInt32(addr *int32, val int32)
func storeUint32(addr *uint32, val uint32)
func storeUintptr(addr *uintptr, val uintptr)
func storePointer(addr *unsafe.Pointer, val unsafe.Pointer)

func StoreInt32(addr *int32, val int32) {
	storeInt32(addr, val)
}
func StoreUint32(addr *uint32, val uint32) {
	storeUint32(addr, val)
}
func StoreUintptr(addr *uintptr, val uintptr) {
	storeUintptr(addr, val)
}
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	storePointer(addr, val)
}