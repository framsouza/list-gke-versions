package listClusters

import (
	"fmt"
	"os"

	"text/tabwriter"

	container "google.golang.org/api/container/v1"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func ListClusters(svc *container.Service, projectID, zone string) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, ' ', 0)
	defer w.Flush()

	// retriving cluster list
	list, err := svc.Projects.Zones.Clusters.List(projectID, zone).Do()
	if err != nil {
		return fmt.Errorf("failed to list clusters: %v", err)
	}

	// header
	fmt.Fprintf(w, "%s\t\t%s\t\t%s", "CLUSTER NAME", "MASTER VERSION", "NODE VERSION")

	for _, v := range list.Clusters {
		fmt.Fprintf(w, "\n%s\t\t", v.Name)
		fmt.Fprintf(w, "%s\t\t", v.CurrentMasterVersion)
		fmt.Fprintf(w, "%s\t\t", v.CurrentNodeVersion)

	}

	return nil
}
