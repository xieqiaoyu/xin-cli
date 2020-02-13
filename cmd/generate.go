package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xieqiaoyu/xin-cli/generate/reqjsonschema"
	"os"
)

func GenApiReqSchemaCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "genApiReqSchema",
		Short: "generate api request schema",
		Long:  "generate api request schema by loading comment",
		Run: func(cmd *cobra.Command, args []string) {
			schemas, err := reqjsonschema.LoadAndParse()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = reqjsonschema.GenerateFile(schemas)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("generate success")
		},
	}
}
