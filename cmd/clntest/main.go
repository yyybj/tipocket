package main

import (
	"context"
	"flag"
	"time"
	"github.com/pingcap/tipocket/cmd/util"
	"github.com/pingcap/tipocket/pkg/cluster"
	"github.com/pingcap/tipocket/pkg/control"
	test_infra "github.com/yyybj/tipocket/pkg/test-infra"
	"github.com/pingcap/tipocket/pkg/test-infra/fixture"
	cll "github.com/yyybj/tipocket/tests/clntest"

	// use mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	//numRows          = flag.Int("rows", 10000000, "the number of rows")
	//smallConcurrency = flag.Int("small-concurrency", 32, "concurrency of small queries")
	//smallTimeout     = flag.Duration("small-timeout", 100*time.Millisecond, "maximum latency of small queries")
	//largeConcurrency = flag.Int("big-concurrency", 1, "concurrency of large queries")
	//largeTimeout     = flag.Duration("large-timeout", 10*time.Second, "maximum latency of big queries")
	//replicaRead      = flag.String("tidb-replica-read", "leader", "tidb_replica_read mode, support values: leader / follower / leader-and-follower, default value: leader.")
	testTime	= flag.Duration("httptest-time", 100*time.Second,"http test duration")
)

func main() {
	flag.Parse()
	cfg := control.Config{
		Mode:        control.ModeSelfScheduled,
		ClientCount: 1,
		RunTime:     fixture.Context.RunTime,
		RunRound:    1,
	}
	c := fixture.Context
	suit := util.Suit{
		Config:   &cfg,
		Provider: cluster.NewDefaultClusterProvider(),
		ClientCreator: cll.NewHttpCreator(*testTime),
 		NemesisGens: util.ParseNemesisGenerators(fixture.Context.Nemesis),
		ClusterDefs: test_infra.NewEmptyCluster(c.Namespace),
	}
	suit.Run(context.Background())
}
