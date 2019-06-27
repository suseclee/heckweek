package main

import "github.com/spf13/cobra"
import "k8s.io/klog"
import "github.com/spf13/pflag"
import "fmt"
import "os"
import "flag"
//import "strings"


var (
	Version   string = "Heckweek"
	BuildDate string = "July 2019"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(os.Stderr, "deploy version: %s  %s \n", Version, BuildDate)
		},
	}
}

type nodeOptions struct {
	count string
	image string
	cpu   string
	mem   string
}

func NewNodeCmd() *cobra.Command {
	nodeOptions := &nodeOptions{}
	cmd := &cobra.Command{
		Use:   "deploy <num of nodes> --image <image> --cpu <num> --mem <mem>",
		Short: "Deploy nodes in aws",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		//	err := cluster.Init(cluster.InitConfiguration{
		//		ClusterName:  args[0],
		//		ControlPlane: initOptions.ControlPlane,
		//	})
		//	if err != nil {
		//		klog.Fatalf("init failed due to error: %s", err)
		//	}
		image, err := cmd.Flags().GetString("image")
		cpu, err := cmd.Flags().GetString("cpu")
		mem, err := cmd.Flags().GetString("mem")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		fmt.Fprintf(os.Stderr, "deploy %s  --image %s --cpu %s --mem %s\n", args[0], image, cpu, mem)
		},
	}
	//cmd.Flags().StringVar("A","B","C","D")
	cmd.Flags().StringVarP(&nodeOptions.image, "image", "i", "", "image-name to use")
	cmd.Flags().StringVarP(&nodeOptions.cpu, "cpu", "c", "", "number of cpus per node")
	cmd.Flags().StringVarP(&nodeOptions.mem, "mem", "m", "", "memory size per node 4096")
	cmd.MarkFlagRequired("image")
	cmd.MarkFlagRequired("cpu")
	cmd.MarkFlagRequired("mem")
//	flag.Parse()
	return cmd
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		// grab the base filename if the binary file is link
		//Use: filepath.Base(os.Args[0]),
	}

	cmd.AddCommand(
		NewVersionCmd(),
		NewNodeCmd(),
		//skuba.NewNodeCmd(),
	)

	register(cmd.PersistentFlags(), "v")

	return cmd

}
func register(local *pflag.FlagSet, globalName string) {
	if f := flag.CommandLine.Lookup(globalName); f != nil {
		pflagFlag := pflag.PFlagFromGoFlag(f)
		local.AddFlag(pflagFlag)
	} else {
		klog.Fatalf("failed to find flag in global flagset (flag): %s", globalName)
	}
}


func main() {
	fmt.Println("** This is a Heckweek project and NOT intended for production usage. **")
	klog.InitFlags(nil)
	cmd := newRootCmd()
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}