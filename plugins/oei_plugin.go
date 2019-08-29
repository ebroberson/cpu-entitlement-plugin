package plugins

import (
	"os"

	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/cf/trace"
	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cpu-entitlement-plugin/cf"
	"code.cloudfoundry.org/cpu-entitlement-plugin/fetchers"
	"code.cloudfoundry.org/cpu-entitlement-plugin/output"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter"
)

type CPUEntitlementAdminPlugin struct{}

func NewOverEntitlementInstancesPlugin() CPUEntitlementAdminPlugin {
	return CPUEntitlementAdminPlugin{}
}

func (p CPUEntitlementAdminPlugin) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "CLI-MESSAGE-UNINSTALL" {
		os.Exit(0)
	}

	traceLogger := trace.NewLogger(os.Stdout, true, os.Getenv("CF_TRACE"), "")
	ui := terminal.NewUI(os.Stdin, os.Stdout, terminal.NewTeePrinter(os.Stdout), traceLogger)

	logCacheURL, err := getLogCacheURL(cli)
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}

	fetcher := fetchers.NewCumulativeUsage(createLogClient(logCacheURL, cli.AccessToken))
	cfClient := cf.NewClient(cli)
	reporter := reporter.NewOverEntitlementInstances(cfClient, fetcher)
	renderer := output.NewOverEntitlementInstancesRenderer(output.NewTerminalDisplay(ui))
	runner := NewOverEntitlementInstancesRunner(reporter, renderer)

	err = runner.Run()
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}
}

func (p CPUEntitlementAdminPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "CPUEntitlementAdminPlugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "over-entitlement-instances",
				Alias:    "oei",
				HelpText: "See which instances are over entitlement",
				UsageDetails: plugin.Usage{
					Usage: "cf over-entitlement-instances",
				},
			},
		},
	}
}