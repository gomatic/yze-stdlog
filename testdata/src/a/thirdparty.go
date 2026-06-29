package a

// A third-party package whose import path ends in "log" must NOT be flagged: the
// analyzer matches the exact path "log", not a suffix or substring.
import tplog "thirdparty/log"

var _ = tplog.Println
