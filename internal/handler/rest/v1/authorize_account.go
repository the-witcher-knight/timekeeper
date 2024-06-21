package v1

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio/respond"
)

func (hdl Handler) AuthorizeAccount() http.HandlerFunc {
	return httpio.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		req, err := httpio.DecodeValidate[authorizeAccountRequest](w, r)
		if err != nil {
			return toAPIError(err)
		}

		ctrlInput := common.HexToAddress(req.Account)
		if err := hdl.bcAuthCtrl.AuthorizeAccount(ctx, ctrlInput); err != nil {
			return toCtrlError(err)
		}

		respond.OK(httpio.Message{
			Code: "authorized",
			Desc: "Account authorized",
		}).Write(w, r)
		return nil
	})
}

type authorizeAccountRequest struct {
	Account string `json:"account"`
}

func (req authorizeAccountRequest) Valid(ctx context.Context) error {
	if req.Account == "" {
		return errAccountIsRequired
	}

	if !common.IsHexAddress(req.Account) {
		return errAccountAddressIsInvalid
	}

	return nil
}
