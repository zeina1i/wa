// Code generated by "stringer -linecomment -type ClauseType"; DO NOT EDIT.

package llenum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ClauseTypeCatch-1]
	_ = x[ClauseTypeFilter-2]
}

const _ClauseType_name = "catchfilter"

var _ClauseType_index = [...]uint8{0, 5, 11}

func (i ClauseType) String() string {
	i -= 1
	if i >= ClauseType(len(_ClauseType_index)-1) {
		return "ClauseType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ClauseType_name[_ClauseType_index[i]:_ClauseType_index[i+1]]
}
