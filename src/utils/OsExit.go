package utils

import "os"

// OsExit は os.Exit のエイリアスです. テストで os.Exit の終了ステータスをモックする場合に使われます.
var OsExit = os.Exit
