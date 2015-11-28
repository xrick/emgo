package main

import (
	"os"
	"strconv"
	"strings"
	"text/template"
)

type mmioCtx struct {
	Pkg   string
	Base  string
	Group string
	Len   uint
	Regs  []mmioReg
}

type mmioReg struct {
	Reg   string
	Bits  string
	Field string
	N     string
}

func mmio(pkg, f, txt string) {
	lines := strings.Split(txt, "\n")
	_, base := nameval(lines[0], ':')
	if base == "" {
		fdie(f, "BaseAddr not set")
	}
	ctx := &mmioCtx{
		Pkg:   pkg,
		Base:  base,
		Group: strings.TrimSuffix(f, ".go"),
	}
	for _, line := range lines[1:] {
		num, name := nameval(line, ':')
		if num == "" || name == "" {
			continue
		}
		n, err := strconv.ParseUint(num, 0, 0)
		if err != nil {
			fdie(f, "bad index %s: %v", num, err)
		}
		ctx.Regs = append(ctx.Regs, mmioReg{
			Reg:   name,
			Bits:  name + "_Bits",
			Field: name + "_Field",
			N:     num,
		})
		if ctx.Len < uint(n) {
			ctx.Len = uint(n)
		}
	}
	ctx.Len++
	w, err := os.Create("xgen_" + f)
	checkErr(err)
	checkErr(mmioTmpl.Execute(w, ctx))
	checkErr(w.Close())
}

const mmioText = `package {{.Pkg}}

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"
)

{{$rf := .Group}}
func {{$rf}}(n uint) *mmio.U32 {
	return &(*[{{.Len}}]mmio.U32)(unsafe.Pointer(uintptr({{.Base}})))[n]
}
{{range .Regs}}
{{$rfn := print $rf "(" .N ")"}}
type {{.Bits}} uint32

func (m {{.Bits}}) Load() {{.Bits}}   { return {{.Bits}}({{$rfn}}.Bits(uint32(m))) }
func (m {{.Bits}}) Store(b {{.Bits}}) { {{$rfn}}.StoreBits(uint32(m), uint32(b)) }
func (m {{.Bits}}) Set()    { {{$rfn}}.SetBits(uint32(m)) }
func (m {{.Bits}}) Clear() { {{$rfn}}.ClearBits(uint32(m)) }

type {{.Field}} uint16

func (f {{.Field}}) Load() int   { return {{$rfn}}.Field(uint16(f)) }
func (f {{.Field}}) Store(v int) { {{$rfn}}.SetField(uint16(f), v) }

func {{.Reg}}_Load() {{.Bits}}   { return {{.Bits}}({{$rfn}}.Load()) }
func {{.Reg}}_Store(b {{.Bits}}) { {{$rfn}}.Store(uint32(b)) }
{{end}}
`

var mmioTmpl = template.Must(template.New("mmio").Parse(mmioText))