type Logger struct {
	m map[string]map[int]bool
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{
		make(map[string]map[int]bool),
	}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
If this method returns false, the message will not be printed.
The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	for i := timestamp; i > timestamp-10; i-- {
		if mapping, ok := this.m[message]; ok && mapping[i] {
			return false
		}
	}

	if _, ok := this.m[message]; !ok {
		this.m[message] = make(map[int]bool)
	}
	this.m[message][timestamp] = true

	return true
}

/**
* Your Logger object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.ShouldPrintMessage(timestamp,message);
 */