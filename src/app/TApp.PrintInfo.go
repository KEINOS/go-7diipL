package app

import (
	"github.com/mkideal/cli"
	"golang.org/x/xerrors"
)

// PrintInfo は ctx に API 情報を書き込みます.
func (a *TApp) PrintInfo(ctx *cli.Context) error {
	info, err := a.GetUniformedInfo()
	if err != nil {
		return xerrors.Errorf("API 情報の取得に失敗しました: %v", err)
	}

	ctx.String("%s\n", info)

	return nil
}
