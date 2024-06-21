package v1

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) DeauthorizeAccount() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		req, err := httpio.DecodeValidate[authorizeAccountRequest](w, r)
		if err != nil {
			return toAPIError(err)
		}

		ctrlInput := common.HexToAddress(req.Account)
		if err := hdl.bcAuthCtrl.DeauthorizeAccount(ctx, ctrlInput); err != nil {
			return toCtrlError(err)
		}

		respond.OK(httpio.Message{
			Code: "deauthorized",
			Desc: "Account deauthorized",
		}).Write(w, r)
		return nil
	})
}
