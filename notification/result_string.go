// Code generated by "stringer -type Result"; DO NOT EDIT.

package notification

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ResultAcknowledge-0]
	_ = x[ResultResolve-1]
}

const _Result_name = "ResultAcknowledgeResultResolve"

var _Result_index = [...]uint8{0, 17, 30}

func (i Result) String() string {
	if i < 0 || i >= Result(len(_Result_index)-1) {
		return "Result(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Result_name[_Result_index[i]:_Result_index[i+1]]
}
