// Package itm provides interface to Instrumentation Trace Macrocell
package itm

import (
	"mmio"
	"unsafe"
)

type regs struct {
	stim [256]mmio.Reg32
	_    [640]uint32
	te   [8]mmio.Reg32
	_    [8]uint32
	tp   mmio.Reg32
	_    [15]uint32
	tc   mmio.Reg32
}

var irs = (*regs)(unsafe.Pointer(uintptr(0xe0000000)))

type Control uint32

// Flags
const (
	ITMEna Control = 1 << iota
	TSEna
	SyncEna
	TxEna
	SWOEna

	Busy Control = 1 << 23
)

func (c Control) TSPresc() int {
	return int(c>>8) & 3
}

func (c *Control) SetTSPresc(f int) {
	*c = *c&^(3<<8) | Control(f&3)<<8
}

func (c Control) GTSFreq() int {
	return int(c>>10) & 3
}

func (c *Control) SetGTSFreq(f int) {
	*c = *c&^(3<<10) | Control(f&3)<<10
}

// Ctrl returns value of ITM Trace Control Register.
func Ctrl() Control {
	return Control(irs.tc.Load())
}

func SetCtrl(c Control) {
	irs.tc.Store(uint32(c))
}

// PrivMask returns conten of Trace Privilege Register. Every bit in returned
// value corresponds to eight stimulus ports. If bit is set then the
// corresponding ports can be accessed by privileged code only.
//
// Bits for unimplemented ports are always returned as 0. To determine the
// number of implemnted ports call SetPrivMask(0xffffffff) and next call
// PrivMask().
func PrivMask() uint32 {
	return irs.tp.Load()
}

// SetPrivMask writes mask to Trace Privilege Register.
func SetPrivMask(mask uint32) {
	irs.tp.Store(mask)
}

type Port int

// Enabled returns true if port is enabled.
func (p Port) Enabled() bool {
	bit := int(p & 31)
	p >>= 5
	return irs.te[p].Bit(bit)
}

// Enable enables stimulus por.
func (p Port) Enable() {
	bit := int(p & 31)
	p >>= 5
	irs.te[p].SetBit(bit)
}

// Disable enables stimulus por.
func (p Port) Disable() {
	bit := int(p & 31)
	p >>= 5
	irs.te[p].ClearBit(bit)
}

// Ready returns true if port can accept data.
func (p Port) Ready() bool {
	return irs.stim[p].Bit(0)
}

// Store8 stores byte into p.
func (p Port) Store8(b byte) {
	mmio.PtrReg8(uintptr(unsafe.Pointer(&irs.stim[p]))).Store(b)
}

// Store16 stores half-word into p.
func (p Port) Store16(h uint16) {
	mmio.PtrReg16(uintptr(unsafe.Pointer(&irs.stim[p]))).Store(h)
}

// Store32 stores word into p.
func (p Port) Store32(w uint32) {
	irs.stim[p].Store(w)
}

// Write implements io.Writer interface.
func (p Port) Write(data []byte) (int, error) {
loop:
	for _, b := range data {
		for !p.Ready() {
			if !p.Enabled() || Ctrl()&ITMEna == 0 {
				// Silently discard data.
				break loop
			}
		}
		p.Store8(b)
	}
	return len(data), nil
}

// WriteString implements io.StringWriter interface.
func (p Port) WriteString(s string) (int, error) {
loop:
	for i := 0; i < len(s); i++ {
		for !p.Ready() {
			if !p.Enabled() || Ctrl()&ITMEna == 0 {
				// Silently discard data.
				break loop
			}
		}
		p.Store8(s[i])
	}
	return len(s), nil
}