package event

type Signal string

type SignalCallback func(data any)

const (
	SignalCellsCreated  Signal = "CellsCreated"
	SignalCellHovered   Signal = "CellHovered"
	SignalCellUnhovered Signal = "CellUnhovered"
	SignalCellPressed   Signal = "CellPressed"
	SignalTeamsSorted   Signal = "TeamsSorted"
)

type SignalBus struct {
	listeners map[Signal][]SignalCallback
}

func NewSignalBus() *SignalBus {
	return &SignalBus{
		listeners: make(map[Signal][]SignalCallback),
	}
}

func (bus *SignalBus) Connect(signal Signal, callback SignalCallback) {
	bus.listeners[signal] = append(bus.listeners[signal], callback)
}

func (bus *SignalBus) Emit(signal Signal, data any) {
	for _, cb := range bus.listeners[signal] {
		cb(data)
	}
}
