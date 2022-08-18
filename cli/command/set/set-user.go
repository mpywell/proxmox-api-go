package set

import (
	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

var set_userCmd = &cobra.Command{
	Use:   "user USERID PASSWORD",
	Short: "Sets the current state of a user",
	Long: `Sets the current state of a user.
Depending on the current state of the user, the user will be created or updated.
The config can be set with the --file flag or piped from stdin.
For config examples see "example user"`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		id := cli.ValidateIDset(args, 0, "UserID")
		config, err := proxmox.NewConfigUserFromJson(cli.NewConfig())
		if err != nil {
			return
		}
		var password string
		if len(args) > 1 {
			password = args[1]
		}
		c := cli.NewClient()
		err = config.SetUser(id, password, c)
		if err != nil {
			return
		}
		cli.PrintItemSet(setCmd.OutOrStdout() ,id ,"User")
		return
	},
}

func init() {
	setCmd.AddCommand(set_userCmd)
}