package cfg

type KubernetesPodInfoConfig struct {
	NodeName     string `env:"K8S_NODE_NAME"`
	PodName      string `env:"K8S_POD_NAME"`
	PodNamespace string `env:"K8S_POD_NAMESPACE"`
	PodIP        string `env:"K8S_POD_IP"`
}
