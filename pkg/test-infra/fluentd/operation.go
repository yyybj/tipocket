package fluentd
import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	//"github.com/pingcap/tipocket/pkg/test-infra/fixture"
	"github.com/pingcap/tipocket/pkg/test-infra/tests"
	"github.com/pingcap/tipocket/pkg/cluster"
	//"github.com/pingcap/tipocket/pkg/test-infra/util"
)


type Ops struct {
	cli   client.Client
	name  string
}

func New(name string) *Ops {
	return &Ops{
		cli:   tests.TestClient.Cli,
		name:  name,
	}
}

func (o *Ops) Apply() error {
	return nil
}

// Delete MySQL instance.
func (o *Ops) Delete() error {
	return nil
}

func (o *Ops) GetNodes() ([]cluster.Node, error) {
	//pod := &corev1.Pod{} // only 1 replica
	//err := o.cli.Get(context.Background(), client.ObjectKey{
	//	Namespace: o.mysql.Sts.ObjectMeta.Namespace,
	//	Name:      fmt.Sprintf("%s-0", o.mysql.Sts.ObjectMeta.Name),
	//}, pod)
	//if err != nil {
	//	return []cluster.Node{}, err
	//}
	//
	//return []cluster.Node{{
	//	Namespace: pod.ObjectMeta.Namespace,
	//	PodName:   pod.ObjectMeta.Name,
	//	IP:        pod.Status.PodIP,
	//	Component: cluster.MySQL,
	//	Port:      util.FindPort(pod.ObjectMeta.Name, string(cluster.MySQL), pod.Spec.Containers),
	//}}, nil

	return []cluster.Node{{
		Namespace: "default",
		PodName:   "web-show-86c9f66c67-jd2ts",
		IP:        "192.168.64.51",
		Component: cluster.MySQL,
		Port:      8081,
	}}, nil
}

func (o *Ops) GetClientNodes() ([]cluster.ClientNode, error) {
	var clientNodes []cluster.ClientNode
	//ips, err := util.GetNodeIPs(o.cli, o.mysql.Sts.Namespace, o.mysql.Sts.ObjectMeta.Labels)
	//if err != nil {
	//	return clientNodes, err
	//} else if len(ips) == 0 {
	//	return clientNodes, errors.New("k8s node not found")
	//}
	//
	//svc, err := util.GetServiceByMeta(o.cli, o.mysql.Svc)
	//if err != nil {
	//	return clientNodes, err
	//}
	//port := getMySQLNodePort(svc)
	//
	var ips = []string{"1.2.3.6","1.2.3.7"}
	for _, ip := range ips {
		clientNodes = append(clientNodes, cluster.ClientNode{
			Namespace:   "default",
			ClusterName: "clustertest",
			Component:   cluster.MySQL,
			IP:          ip,
			Port:        8081,
		})
	}
	return clientNodes, nil
}
