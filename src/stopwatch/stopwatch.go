package stopwatch

import "time"

// Stopwatch は処理時間を計測します
type stopwatch struct {
	start   *time.Time
	elapsed float64
}

// Stopwatch のインターフェイスです
type Stopwatch interface {
	Start() Stopwatch
	Stop() Stopwatch
	Reset() Stopwatch
	Restart() Stopwatch
	GetElapsedSeconds() float64
}

// New はインスタンスを生成します
func New() Stopwatch {
	return &stopwatch{elapsed: 0}
}

// StartNew はStopwatchをスタートさせてインスタンスを生成します
func StartNew() Stopwatch {
	var sw = New()
	return sw.Start()
}

// Start はStopwatchをスタートします
func (sw *stopwatch) Start() Stopwatch {
	var start = time.Now()
	sw.start = &start
	return sw
}

// Stop はStopwatchを一時停止します
func (sw *stopwatch) Stop() Stopwatch {
	sw.GetElapsedSeconds()
	sw.start = nil
	return sw
}

// Reset はStopwatchを初期化します
func (sw *stopwatch) Reset() Stopwatch {
	sw.elapsed = 0
	sw.start = nil
	return sw
}

// Restart はStopwatchを初期化して再スタートします
func (sw *stopwatch) Restart() Stopwatch {
	sw.Reset()
	return sw.Start()
}

// GetElapsedSeconds は経過時間を取得します
func (sw stopwatch) GetElapsedSeconds() float64 {
	if sw.start == nil {
		return sw.elapsed
	}
	sw.elapsed += (time.Now().Sub(*sw.start)).Seconds()
	return sw.elapsed
}
