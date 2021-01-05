package main

type Logger struct {
	dict map[string]int
	rate int
}

func New(rate int) Logger {
	return Logger{
		dict: make(map[string]int),
		rate: rate,
	}
}

func (l *Logger) ShouldPrintMessage(t int, m string) bool {
	if ot, ok := l.dict[m]; ok && t-ot < l.rate {
		return false
	}

	l.dict[m] = t
	return true
}
