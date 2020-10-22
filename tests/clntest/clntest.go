package clntest

import (
	"context"
//"database/sql"
	"fmt"
//"math/rand"
//"sync"
	"time"

//"github.com/codahale/hdrhistogram"
	"github.com/ngaut/log"

	"github.com/pingcap/tipocket/pkg/cluster"
	"github.com/pingcap/tipocket/pkg/core"
//"github.com/pingcap/tipocket/util"
//"github.com/gruntwork-io/terratest/modules/http-helper"
//"bytes"
)

// CaseCreator is a creator of a read-stress test
//type CaseCreator struct {
//	NumRows          int
//	LargeConcurrency int
//	LargeTimeout     time.Duration
//	SmallConcurrency int
//	SmallTimeout     time.Duration
//	ReplicaRead      string
//}

type httpCreator struct {
	testTime time.Duration
}

func NewHttpCreator(testTime time.Duration) core.ClientCreator {
	return &httpCreator{
		testTime: testTime,
	}
}

func (c *httpCreator) Create(_ cluster.ClientNode) core.Client {
	return &httpClient{
		testTime: 		c.testTime,
	}
}

type httpClient struct {
	testTime time.Duration
}

// Create creates a read-stress test client
//func (c CaseCreator) Create(node cluster.ClientNode) core.Client {
//	return &stressClient{
//		numRows:          c.NumRows,
//		largeConcurrency: c.LargeConcurrency,
//		largeTimeout:     c.LargeTimeout,
//		smallConcurrency: c.SmallConcurrency,
//		smallTimeout:     c.SmallTimeout,
//		replicaRead:      c.ReplicaRead,
//	}
//}
//
//type stressClient struct {
//	numRows          int
//	largeConcurrency int
//	largeTimeout     time.Duration
//	smallConcurrency int
//	smallTimeout     time.Duration
//	db               *sql.DB
//	replicaRead      string
//}

func (c *httpClient) SetUp (ctx context.Context, _ []cluster.Node, clientNodes []cluster.ClientNode, idx int) error {
	print("clntest setup")
	return nil
}

func (c *httpClient) TearDown(ctx context.Context, nodes []cluster.ClientNode, idx int) error {
	print("clntest teardown")
	return nil
}

//func (c *stressClient) SetUp(ctx context.Context, _ []cluster.Node, clientNodes []cluster.ClientNode, idx int) error {
//	// only prepare data through the first TiDB
//	if idx != 0 {
//		return nil
//	}
//
//	node := clientNodes[idx]
//	dsn := fmt.Sprintf("root@tcp(%s:%d)/test", node.IP, node.Port)
//
//	log.Info("[stressClient] Initializing...")
//	var err error
//	concurrency := c.largeConcurrency + c.smallConcurrency
//	if concurrency < 128 {
//		concurrency = 128
//	}
//	c.db, err = util.OpenDB(dsn, concurrency)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	util.RandomlyChangeReplicaRead("read-stress", c.replicaRead, c.db)
//
//	if _, err := c.db.Exec("DROP TABLE IF EXISTS t"); err != nil {
//		log.Fatal(err)
//	}
//	if _, err := c.db.Exec("CREATE TABLE t(id INT PRIMARY KEY, k INT, v varchar(255), KEY (v))"); err != nil {
//		log.Fatal(err)
//	}
//
//	var wg sync.WaitGroup
//	start := 0
//	step := c.numRows / 128
//	for i := 0; i < 128; i++ {
//		end := start + step
//		if i+1 == 128 {
//			end = c.numRows
//		}
//		wg.Add(1)
//		go func(start, end int) {
//			defer wg.Done()
//			if err != nil {
//				log.Fatal(err)
//			}
//			stmt, err := c.db.Prepare("INSERT INTO t VALUES (?, ?, ?)")
//			if err != nil {
//				log.Fatal(err)
//			}
//			for ; start < end; start += 256 {
//				txn, err := c.db.Begin()
//				if err != nil {
//					log.Fatal(err)
//				}
//				txnStmt := txn.Stmt(stmt)
//				for id := start; id < start+256 && id < end; id++ {
//					if _, err := txnStmt.Exec(id, rand.Intn(64), string(rand.Intn(26)+'a')); err != nil {
//						log.Fatal(err)
//					}
//				}
//				if err := txn.Commit(); err != nil {
//					log.Fatal(err)
//				}
//			}
//
//		}(start, end)
//		start = end
//	}
//	if _, err := c.db.Exec("ANALYZE TABLE t"); err != nil {
//		log.Fatal(err)
//	}
//	wg.Wait()
//	log.Info("Data prepared")
//	return nil
//}


//func (c *stressClient) TearDown(ctx context.Context, nodes []cluster.ClientNode, idx int) error {
//	// only tear down through the first TiDB
//	if idx != 0 {
//		return nil
//	}
//	_, err := c.db.Exec("DROP TABLE IF EXISTS t")
//	return err
//}

func (c *httpClient) Invoke(ctx context.Context, node cluster.ClientNode, r interface{}) core.UnknownResponse {
	log.Info("httpclient invoke")
	return nil
}

func (c *httpClient) NextRequest() interface{} {
	log.Info("httpclient nextrequest")
	return nil
}

func (c *httpClient) DumpState(ctx context.Context) (interface{}, error) {
	log.Info("httpclient dumpstate")
	return nil, nil
}

func (c *httpClient) Start(ctx context.Context, cfg interface{}, clientNodes []cluster.ClientNode) error {
	//url := "http://gaia-dev.tpaas.domain/devops/v1/projects"
	//headers := map[string]string{"x-jdcloud-oauth-token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6ImNoZW4iLCJhY2NvdW50IjoiamRpZGFhcyIsImlhdCI6MTU5MDExNDQ0NCwiZXhwIjoxNjIxNjUwNDQ0fQ.XPFk_OCFa8x6bWVnsrZ3SN8WfHhjeoQ8-4nL24W9jHPiLgjfmBe4nqwWRzqrtnSNkzLdPfOkH5TD09wlLoy7yFK_2CH2gDcuZLMuNfqgJUtM8Kwkfkj-Tqs5QfjApmXPmade1yIVsEY9JLwtGz5_D6LaVKsJ_77rVkM6sppO_Olr2qlxQsXDc00EmbIY2vgJjw6wKsJ4sDgnAZISIGwm8KCZrJXlzmq6GsoyECyKM2ayP3zzBVaKwQkT_wYUIFjUmVJ1kY2FYEPylLW4l_P1yLobwUC_3CSQ52GmhnIA4KC4Vj5qjUVidJAcu3hFklwvioqB4_WQOdoCM1LnJxBXZQ"}
	//expectedBody := ""
	//body := bytes.NewReader([]byte(expectedBody))
	//statusCode, _ := http_helper.HTTPDo(t, "GET", url, body, headers, nil)

	//select {
	//case <-time.After(20 * time.Second):
	//	fmt.Println("sleep 1s")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	//}

	deadline, ok := ctx.Deadline()
	log.Info(deadline, ok)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	for !ok || time.Now().Before(deadline) {

		time.Sleep(5 * time.Second)
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	log.Info("httpclient start finished")
	return nil
}

//func (c *stressClient) Start(ctx context.Context, cfg interface{}, clientNodes []cluster.ClientNode) error {
//	var wg sync.WaitGroup
//	for i := 0; i < c.largeConcurrency; i++ {
//		wg.Add(1)
//		go func() {
//			finished := false
//			for !finished {
//				ch := make(chan struct{})
//				go func() {
//					if _, err := c.db.Exec("SELECT v, COUNT(*) FROM t WHERE k % 10 = 1 GROUP BY v HAVING v > 'b'"); err != nil {
//						log.Fatal(err)
//					}
//					ch <- struct{}{}
//				}()
//				select {
//				case <-ch:
//					continue
//				case <-time.After(c.largeTimeout):
//					log.Fatal("[stressClient] Large query timed out")
//				case <-ctx.Done():
//					finished = true
//				}
//			}
//			wg.Done()
//		}()
//	}
//	deadline, ok := ctx.Deadline()
//	log.Info(deadline, ok)
//	for !ok || time.Now().Before(deadline) {
//		hist := c.runSmall(ctx, time.Second*10)
//		mean, quantile99 := hist.Mean(), hist.ValueAtQuantile(99)
//		log.Infof("[stressClient] Small queries in the last minute, mean: %d(us), 99th: %d(us)", int64(mean), quantile99)
//		if mean > float64(c.smallTimeout.Microseconds()) {
//			log.Fatal("[stressClient] Small query timed out")
//		}
//	}
//	wg.Wait()
//	log.Info("[stressClient] Test finished.")
//	return nil
//}

//func (c *stressClient) runSmall(ctx context.Context, duration time.Duration) *hdrhistogram.Histogram {
//	var wg sync.WaitGroup
//	durCh := make(chan time.Duration)
//	deadline := time.Now().Add(duration)
//	for i := 0; i < c.smallConcurrency; i++ {
//		wg.Add(1)
//		go func() {
//			for time.Now().Before(deadline) {
//				beginInst := time.Now()
//				txn, err := c.db.Begin()
//				if err != nil {
//					log.Fatal(err)
//				}
//				if _, err := txn.Exec("SELECT v FROM t WHERE id IN (?, ?, ?)", rand.Intn(c.numRows), rand.Intn(c.numRows), rand.Intn(c.numRows)); err != nil {
//					log.Fatal(err)
//				}
//				start := rand.Intn(c.numRows)
//				end := start + 50
//				if _, err := txn.Exec("SELECT k FROM t WHERE id BETWEEN ? AND ? AND v = 'b'", start, end); err != nil {
//					log.Fatal(err)
//				}
//				if err := txn.Commit(); err != nil {
//					log.Fatal(err)
//				}
//				dur := time.Now().Sub(beginInst)
//				durCh <- dur
//			}
//			wg.Done()
//		}()
//	}
//	go func() {
//		wg.Wait()
//		close(durCh)
//	}()
//	hist := hdrhistogram.New(0, 60000000, 3)
//	for dur := range durCh {
//		hist.RecordValue(dur.Microseconds())
//	}
//	return hist
//}
