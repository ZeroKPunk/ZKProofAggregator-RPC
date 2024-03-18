package main

//type Cli struct {
//	GitVersion string
//	GitCommit  string
//	GitDate    string
//	app        *cli.App
//	Flags      []cli.Flag
//}
//
//var (
//	ConfigFlag = &cli.StringFlag{
//		Name:    "config",
//		Value:   "../../aggregator.toml",
//		Aliases: []string{"c"},
//		Usage:   "path to config file",
//		EnvVars: []string{"AGGREGATOR_CONFIG"},
//	}
//	EndFlag = &cli.BoolFlag{
//		Name:  "ginger-crouton",
//		Usage: "is it in the soup?",
//	}
//)
//
//// runProcessor is the entrypoint into the Aggregator service.
//func runProcessor(ctx *cli.Context) error {
//	configPath := ctx.String(ConfigFlag.Name)
//	conf, err := config.LoadConfig(configPath)
//
//	log.Info("Starting processor")
//
//	if err != nil {
//		log.Crit("Failed to load config", "message", err)
//	}
//
//	processor, err := processor.NewProcessor(conf)
//	if err != nil {
//		log.Error("Unable to create processor", "error", err)
//		return err
//	}
//
//	if err := processor.Start(); err != nil {
//		return err
//	}
//
//	//defer processor.Stop()
//
//	log.Info("Processor started")
//
//	signalChan := make(chan os.Signal)
//	// SIGHUP: terminal closed
//	// SIGINT: Ctrl+C
//	// SIGTERM: program exit
//	// SIGQUIT: Ctrl+/
//	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//
//	waitElegantExit(signalChan, processor)
//
//	return nil
//}
//
//func waitElegantExit(signalChan chan os.Signal, processor *processor.Processor) {
//	for i := range signalChan {
//		switch i {
//		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
//			processor.Stop()
//			os.Exit(0)
//		}
//	}
//}
//
//func NewCli(GitVersion string, GitCommit string, GitDate string) *Cli {
//	log.Root().SetHandler(
//		log.LvlFilterHandler(
//			log.LvlInfo,
//			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
//		),
//	)
//
//	flags := []cli.Flag{
//		ConfigFlag,
//	}
//
//	app := &cli.App{
//		Name:        "Processor",
//		Version:     "test",
//		Description: "A spv processor",
//		Commands: []*cli.Command{
//			{
//				Name:        "processor",
//				Flags:       flags,
//				Description: "Runs the processor service",
//				Action:      runProcessor,
//			},
//		},
//	}
//
//	return &Cli{
//		app:   app,
//		Flags: flags,
//	}
//}
//
//// Run make an instance method on Cli called Run that runs cli
//// and returns an error
//func (c *Cli) Run(args []string) error {
//	return c.app.Run(args)
//}
