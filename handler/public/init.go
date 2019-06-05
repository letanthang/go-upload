package handler_public

import (
	"g.ghn.vn/logistic/crm/types"
	"github.com/liip/sheriff"
)

var notFoundErrorMessage = types.PayloadResponse("404", "Không tìm thấy thông tin người dùng")
var optionPublic = &sheriff.Options{
	Groups: []string{"public"},
}

func init() {

}
