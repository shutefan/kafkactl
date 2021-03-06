package create

import (
	"github.com/deviceinsight/kafkactl/operations"
	"github.com/spf13/cobra"
)

var flags operations.CreateTopicFlags

var cmdCreateTopic = &cobra.Command{
	Use:   "topic",
	Short: "create a topic",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		(&operations.TopicOperation{}).CreateTopics(args, flags)
	},
}

func init() {
	cmdCreateTopic.Flags().Int32VarP(&flags.Partitions, "partitions", "p", 1, "number of partitions")
	cmdCreateTopic.Flags().Int16VarP(&flags.ReplicationFactor, "replication-factor", "r", 1, "replication factor")
	cmdCreateTopic.Flags().BoolVarP(&flags.ValidateOnly, "validate-only", "V", false, "validate only")
	cmdCreateTopic.Flags().StringArrayVarP(&flags.Configs, "config", "C", flags.Configs, "configs in format `key=value`")
}
