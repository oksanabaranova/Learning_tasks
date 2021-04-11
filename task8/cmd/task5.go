package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func deployManifest(filename string, username string) {
	// yamlFile, err := ioutil.ReadFile(filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	defer f.Close()
	b := bufio.NewReader(f)
	r := yaml.NewYAMLReader(b)

	doc, err := r.Read()
	if err == io.EOF {
		log.Fatalf("Empty Yaml file: %s\n%s", filename, err)
	}
	if err != nil {
		log.Fatal(err)
	}
	d := scheme.Codecs.UniversalDeserializer()
	obj, _, err := d.Decode(doc, nil, nil)
	if err != nil {
		log.Fatalf("could not decode yaml: %s\n%s", filename, err)
	}
	fmt.Println(obj)
	fmt.Println("---------------")
	clientset := createClientSet()
	deployPod(clientset, obj)
}

func createClientSet() *kubernetes.Clientset {
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" { // check if machine has home directory.
	// 	// read kubeconfig flag. if not provided use config file $HOME/.kube/config
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	// build configuration from the config file.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	// create kubernetes clientset. this clientset can be used to create,delete,patch,list etc for the kubernetes resources
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func deployPod(clientset *kubernetes.Clientset, obj runtime.Object) {
	// now create the pod in kubernetes cluster using the clientset
	podobj := obj.(*core.Pod)
	_, err := clientset.CoreV1().Pods("default").Create(context.Background(), podobj,
		metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Pod created successfully...")
}
