package corvus

import (
	"fmt"
	"github.com/fiftyonefifty/corvus/version"
)

func Version() string {
	v := fmt.Sprintf("Coven v%s", version.String())
	return v
}
