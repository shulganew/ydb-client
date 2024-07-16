package storage

import (
	"context"
	"log"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func DescribeTable(ctx context.Context, c table.Client, path string) error {
	return c.Do(ctx,
		func(ctx context.Context, s table.Session) error {
			desc, err := s.DescribeTable(ctx, path)
			if err != nil {
				return err
			}
			log.Printf("> describe table: %s, %d", path, len(desc.Columns))
			for i := range desc.Columns {
				log.Printf("column, name: %s, %s", desc.Columns[i].Type, desc.Columns[i].Name)
			}

			return nil
		},
	)
}

func DescribeTableOptions(ctx context.Context, c table.Client) error {
	return c.Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
			desc, err := s.DescribeTableOptions(ctx)
			if err != nil {
				return err
			}

			log.Println("> describe_table_options:")

			for i := range desc.TableProfilePresets {
				log.Printf("TableProfilePresets: %d/%d: %+v", i+1,
					len(desc.TableProfilePresets), desc.TableProfilePresets[i],
				)
			}
			for i := range desc.StoragePolicyPresets {
				log.Printf("StoragePolicyPresets: %d/%d: %+v", i+1,
					len(desc.StoragePolicyPresets), desc.StoragePolicyPresets[i],
				)
			}
			for i := range desc.CompactionPolicyPresets {
				log.Printf("CompactionPolicyPresets: %d/%d: %+v", i+1,
					len(desc.CompactionPolicyPresets), desc.CompactionPolicyPresets[i],
				)
			}
			for i := range desc.PartitioningPolicyPresets {
				log.Printf("PartitioningPolicyPresets: %d/%d: %+v", i+1,
					len(desc.PartitioningPolicyPresets), desc.PartitioningPolicyPresets[i],
				)
			}
			for i := range desc.ExecutionPolicyPresets {
				log.Printf("ExecutionPolicyPresets: %d/%d: %+v", i+1,
					len(desc.ExecutionPolicyPresets), desc.ExecutionPolicyPresets[i],
				)
			}
			for i := range desc.ReplicationPolicyPresets {
				log.Printf("ReplicationPolicyPresets: %d/%d: %+v", i+1,
					len(desc.ReplicationPolicyPresets), desc.ReplicationPolicyPresets[i],
				)
			}
			for i := range desc.CachingPolicyPresets {
				log.Printf("CachingPolicyPresets: %d/%d: %+v", i+1,
					len(desc.CachingPolicyPresets), desc.CachingPolicyPresets[i],
				)
			}

			return nil
		},
	)
}
