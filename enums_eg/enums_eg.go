package enums_eg

import "fmt"

//go:generate stringer -type=ProcessState
type ProcessState int 

// enums that uses iota which enable incremental values
const (
	Idle ProcessState = iota // counter, start with 0
	Running
	Halted
	Stopped
	Killed
)

// // enums that uses string values
// var StateName = map[ProcessState]string{
// 	Idle: "Idle",
// 	Halted: "Halted",
// 	Stopped: "Stopped",
// 	Killed: "Killed",
// }

// func (p ProcessState) String() string {
// 	return StateName[p]
// }

func trasition(p ProcessState) ProcessState  {
	switch p {
		case Idle:
			return Running
		case Running:
			return Halted
		case Halted:
			return Stopped
		case Stopped:
			return Killed
		case Killed:
			return Idle
		default:
			panic(fmt.Errorf("invalid process state: %d", p))
	}
}

func Enums() {
	state := Killed
	fmt.Println(state)

	processStatus := trasition(state)
	fmt.Println(processStatus)
}
