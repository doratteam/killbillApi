package killbill_go

import "fmt"

type transaction struct {
	uuid      string
	ownerUuid string
	amount    float32
}
