package nvic

type Vector func()

type Table struct {
	_          Vector `C:"volatile"`
	Reset      Vector `C:"volatile"` // prio -3 (fixed)
	NMI        Vector `C:"volatile"` // prio -2 (fixed)
	HardFault  Vector `C:"volatile"` // prio -1 (fixed)
	MemManage  Vector `C:"volatile"` // prio 0
	BusFault   Vector `C:"volatile"` // prio 1
	UsageFault Vector `C:"volatile"` // prio 2
	_          Vector `C:"volatile"`
	_          Vector `C:"volatile"`
	_          Vector `C:"volatile"`
	_          Vector `C:"volatile"`
	SVCall     Vector `C:"volatile"` // prio 3
	DebugMon   Vector `C:"volatile"` // prio 4
	_          Vector `C:"volatile"`
	PendSV     Vector `C:"volatile"` // prio 5
	SysTick    Vector `C:"volatile"` // prio 6
}