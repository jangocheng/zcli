package i18n

const (
	// cmd
	CmdDeployDesc    = "deploy your application into zerops.io"
	CmdPushDesc      = "deploy your application into zerops.io and build it"
	CmdLogin         = "log you into zerops.io"
	CmdVpn           = "vpn commands group"
	CmdVpnStart      = "start vpn"
	CmdVpnStop       = "stop vpn"
	CmdVpnStatus     = "show vpn status"
	CmdLog           = "logs commands"
	CmdLogShow       = "show logs"
	CmdDaemon        = "daemon commands group"
	CmdDaemonRun     = "run daemon"
	CmdDaemonInstall = "install daemon"
	CmdDaemonRemove  = "remove daemon"

	// flags description
	BuildVersionName = "custom version name"
	BuildWorkingDir  = "working dir, all files path are relative to this directory"
	BuildZipFilePath = "save final zip file"

	// process
	ProcessInvalidState = "process is in wrong state"

	// zipClient
	ZipClientWorkingDirectory = "working directory"
	ZipClientMaxOneAsterix    = "only one *(asterisk) is allowed"
	ZipClientPackingDirectory = "packing directory"
	ZipClientPackingFile      = "packing file"

	// login
	LoginParamsMissing         = "either zeropsLogin + zeropsPassword or zeropsToken params must be set"
	LoginZeropsLoginMissing    = "param zeropsLogin must be set"
	LoginZeropsPasswordMissing = "param zeropsPassword must be set"
	LoginSuccess               = "you are logged in"
	LoginVpnClosed             = "vpn connection was closed"

	// deploy
	BuildDeployProjectNameMissing      = "project name must be filled"
	BuildDeployServiceStackNameMissing = "service name must be filled"
	BuildDeployProjectNotFound         = "project not found"
	BuildDeployProjectsWithSameName    = "there are multiple projects with same name"
	BuildDeployServiceStatus           = "service status"
	BuildDeployTemporaryShutdown       = "temporaryShutdown"
	BuildDeployCreatingPackageStart    = "creating package"
	BuildDeployCreatingPackageDone     = "package created"
	BuildDeployPackageSavedInto        = "package file saved into"
	BuildDeployUploadingPackageStart   = "uploading package"
	BuildDeployUploadingPackageDone    = "package uploaded"
	BuildDeployUploadPackageFailed     = "package upload failed"
	BuildDeployDeployingStart          = "deploying service"
	BuildDeployBuildConfigNotFound     = "config file zerops.yml is not found"
	BuildDeployBuildConfigEmpty        = "config file zerops.yml is empty"
	BuildDeployBuildConfigTooLarge     = "max. size of zerops.yml is 10 MB"
	BuildDeploySuccess                 = "service deployed"

	// vpn start
	VpnStartProjectNameIsEmpty         = "project name must be filled"
	VpnStartProjectNotFound            = "project not found"
	VpnStartInterfaceAssignFailed      = "interface name assign failed"
	VpnStartWireguardInterfaceNotfound = "wireguard interface not found"
	VpnStartProjectsWithSameName       = "there are multiple projects with same name"
	VpnStartDaemonIsUnavailable        = "daemon is currently unavailable, did you install it?"
	VpnStartInstallDaemonPrompt        = "is it ok if we are going to install daemon for you?"
	VpnStartTerminatedByUser           = "when you will be ready, try `./zcli daemon install`"
	VpnStartUserIsUnableToWriteYorN    = "type 'y' or 'n' please"
	VpnStartWireguardUtunError         = "we weren't able to start vpn, there is possibility that you have another vpn, if so, try to shut it down"
	VpnStartVpnNotReachable            = "zerops vpn servers aren't reachable"
	VpnStartTunnelIsNotAlive           = "we weren't able to establish zerops vpn"

	// vpn status
	VpnStatusDaemonIsUnavailable     = "daemon is currently unavailable, did you install it?"
	VpnStatusTunnelStatusActive      = "wireguard tunnel is working properly"
	VpnStatusTunnelStatusSetInactive = "wireguard tunnel is established but it isn't working properly, try `./zcli vpn start` command"
	VpnStatusTunnelStatusUnset       = "wireguard tunnel isn't established, try `./zcli vpn start` command"
	VpnStatusDnsStatusActive         = "dns is working properly"
	VpnStatusDnsStatusSetInactive    = "dns is set but it isn't working properly, try `./zcli vpn start` command"
	VpnStatusDnsStatusUnset          = "dns isn't set, try `./zcli vpn start` command"
	VpnStatusAdditionalInfo          = "additional info:"
	VpnStatusDnsCheckError           = "we weren't able to check that dns working correctly"

	// vpn stop
	VpnStopDaemonIsUnavailable   = "daemon is currently unavailable, did you install it?"
	VpnStopSuccess               = "vpn connection was closed"
	VpnStopAdditionalInfo        = "additional info:"
	VpnStopAdditionalInfoMessage = "dns could be set by yourself, if so it must be removed manually"

	// daemon
	DaemonInstallerDesc = "zerops daemon"

	// daemon install
	DaemonInstallSuccess                 = "zerops daemon has been installed"
	DaemonInstallWireguardNotFound       = "wireguard was not found"
	DaemonInstallWireguardNotFoundDarwin = "wireguard was not found, try `brew install wireguard-tools`"

	// daemon remove
	DaemonRemoveStopVpnUnavailable = "zerops daemon isn't running, vpn couldn't be removed"
	DaemonRemoveSuccess            = "zerops daemon has been removed"

	// generic
	GrpcApiTimeout    = "zerops api didn't response within assigned time, try it again later"
	GrpcVpnApiTimeout = "zerops vpn server didn't response within assigned time, try it again later"
)
