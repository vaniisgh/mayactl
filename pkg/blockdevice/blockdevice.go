/*
Copyright 2020-2021 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package blockdevice

import (
	"github.com/dustin/go-humanize"
	"github.com/openebs/api/v2/pkg/apis/openebs.io/v1alpha1"
	"github.com/openebs/openebsctl/pkg/client"
	"github.com/openebs/openebsctl/pkg/util"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/printers"
)

const (
	firstElemPrefix = `├─`
	lastElemPrefix  = `└─`
)

// Get manages various implementations of blockdevice listing
func Get(bds []string, openebsNS string) error {
	// TODO: Prefer passing the client from outside
	k, _ := client.NewK8sClient(openebsNS)
	err := createTreeByNode(k, bds)
	if err != nil {
		return err
	}
	return nil
}

// createTreeByNode uses the [node <- list of bds on the node] and creates a tree like output,
// also showing the relevant details to the bds.
func createTreeByNode(k *client.K8sClient, bds []string) error {
	// 1. Get a list of the BlockDevices
	var bdList *v1alpha1.BlockDeviceList
	bdList, err := k.GetBDs(bds, "")
	if err != nil {
		return err
	}
	// 2. Create a map out of the list of bds, by their node names.
	var nodeBDlistMap = map[string][]v1alpha1.BlockDevice{}
	for _, bd := range bdList.Items {
		if _, ok := nodeBDlistMap[bd.Spec.NodeAttributes.NodeName]; ok {
			// Append to the node if key exists
			nodeBDlistMap[bd.Spec.NodeAttributes.NodeName] = append(nodeBDlistMap[bd.Spec.NodeAttributes.NodeName], bd)
		} else {
			// Create new key with node name and add the bd, if node does not exist
			nodeBDlistMap[bd.Spec.NodeAttributes.NodeName] = []v1alpha1.BlockDevice{bd}
		}
	}
	var rows []metav1.TableRow
	if len(nodeBDlistMap) == 0 {
		// If there are no block devices show error
		return errors.New("no blockdevices found in the " + k.Ns + " namespace")
	} else {
		for key, value := range nodeBDlistMap {
			// Create the root, which contains only the node-name
			rows = append(rows, metav1.TableRow{Cells: []interface{}{key, "", "", "", "", "", ""}})
			for i, bd := range value {
				// If the bd is the last bd in the list, or the list has only one bd
				// append lastElementPrefix before bd name
				if i == len(value)-1 {
					rows = append(rows, metav1.TableRow{
						Cells: []interface{}{
							lastElemPrefix + bd.Name,
							bd.Spec.Path,
							humanize.IBytes(bd.Spec.Capacity.Storage),
							bd.Status.ClaimState,
							bd.Status.State,
							bd.Spec.FileSystem.Type,
							bd.Spec.FileSystem.Mountpoint,
						}})
				} else {
					// If the bd is the not last bd in the list append firstElementPrefix before
					// bd name which signifies there are more to append in the tree.
					rows = append(rows, metav1.TableRow{
						Cells: []interface{}{
							firstElemPrefix + bd.Name,
							bd.Spec.Path,
							humanize.IBytes(bd.Spec.Capacity.Storage),
							bd.Status.ClaimState,
							bd.Status.State,
							bd.Spec.FileSystem.Type,
							bd.Spec.FileSystem.Mountpoint,
						}})
				}
			}
			// Add an empty row so that the tree looks neat
			rows = append(rows, metav1.TableRow{Cells: []interface{}{"", "", "", "", "", "", ""}})
		}
	}
	// Show the output using cli-runtime
	util.TablePrinter(util.BDTreeListColumnDefinations, rows, printers.PrintOptions{Wide: true})
	return nil
}
