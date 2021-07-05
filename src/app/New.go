package app

func New() *TApp {
	appTmp := new(TApp)

	appTmp.Name = NameDefault
	appTmp.Version = VersionDefault
	appTmp.Argv = new(TFlagOptions)

	appTmp.Argv.SetHelpMsg()

	return appTmp
}
