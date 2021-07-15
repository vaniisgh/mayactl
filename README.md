
<img width="300" align="right" alt="OpenEBS Logo" src="https://raw.githubusercontent.com/cncf/artwork/master/projects/openebs/stacked/color/openebs-stacked-color.png" xmlns="http://www.w3.org/1999/html">

## OpenEBSCTL


[![Go Report Card](https://goreportcard.com/badge/github.com/openebs/openebsctl?)](https://goreportcard.com/report/github.com/openebs/openebsctl)
[![Contributors](https://img.shields.io/github/contributors/openebs/openebsctl)](https://github.com/openebs/openebsctl/graphs/contributors)
[![release](https://img.shields.io/github/release-pre/openebs/openebsctl.svg)](https://github.com/openebs/openebsctl/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mum4k/termdash/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/downloads/openebs/openebsctl/total.svg)](https://github.com//openebs/openebsctl/releases)



OpenEBSCTL is a kubectl plugin to manage OpenEBS storage components.


### Project Status

**Alpha**. Under active development and seeking [contributions from the community](#contributing).
The CLI currently supports managing `cStor`, `Jiva`, `LocalPV-LVM`, `LocalPV-ZFS` Cas-Engines.

### Table of Contents
* [Installation](#installation)
* [Build](#build)
* [Flags](#flags)
* [Usage](#usage)
  * [cStor](#cstor)
    * [Get cStor volumes](#get-cstor-volumes)
    * [Get cStor pools](#get-cstor-pools)
    * [Describe cStor volumes](#describe-cstor-volumes)
    * [Describe cStor pool](#describe-cstor-pool)
  * [Jiva](#jiva)
    * [Get Jiva volumes](#get-jiva-volumes)
    * [Describe Jiva volumes](#describe-jiva-volumes)
  * [LocalPV-LVM](#localpv-lvm)
    * [Get LocalPV-LVM volumes](#get-localpv-lvm-volumes)
    * [Get LocalPV-LVM VolumeGroups](#get-localpv-lvm-volumegroups)
  * [LocalPV-ZFS](#localpv-zfs)
    * [Get LocalPV-ZFS volumes](#get-localpv-zfs-volumes)
    * [Get LocalPV-ZFS Pools](#get-localpv-zfs-pools)
  * [BlockDevice](#blockdevice)
    * [Get BlockDevices by Nodes](#get-blockdevices-by-nodes)
  * [PersistentVolumeClaims](#persistentvolumeclaims)
    * [Describe pvcs](#describe-pvcs)
* [Contributing](#contributing)


### Installation

OpenEBSCTL is available on Linux, macOS and Windows platforms.

* Binaries for Linux, Mac and Windows are available as tarballs and zip in the [release](https://github.com/openebs/openebsctl/releases) page.
* For Linux, download the respective tarball from [release](https://github.com/openebs/openebsctl/releases) page and :-
   ```shell
   tar -xvf kubectl-openebs_v0.1.0_Linux_x86_64.tar.gz
   cd kubectl-openebs_v0.1.0_Linux_x86_64
   sudo mv kubectl-openebs /usr/local/bin/
   ```
  Or, download the `debian` package from the [release](https://github.com/openebs/openebsctl/releases) page and double click it launch the installer if using GUI.<br/><br/>
  Or, we can also use script to install the latest version :-
   ```shell
   wget https://raw.githubusercontent.com/openebs/openebsctl/develop/scripts/install-latest.sh -O - | bash
   ```
* For Mac, download the respective tarball from [release](https://github.com/openebs/openebsctl/releases) page and :-
  ```shell
   tar -xzvf kubectl-openebs_v0.1.0_Darwin_x86_64.tar.gz
   cd kubectl-openebs_v0.1.0_Darwin_x86_64
   sudo mv kubectl-openebs /usr/local/bin/
   ```
* For Windows, download the respective zip from [release](https://github.com/openebs/openebsctl/releases) page and :-
    - Extract the zip, copy the `path` of the folder the contents are in.
    - Add the `path` to the `PATH` environment variable.

### Build

- Clone this repo to your system. `git clone https://github.com/openebs/openebsctl`
- `cd openebsctl`
- Run `make openebsctl`
- Run `kubectl openebs [get|describe] [resource]` to use the plugin

### Flags

* `--openebs-namespace` :- to override the determination of `namespace` where storage engine is installed with the provided value.
* `--namespace, -n` :- to pass the namespace, if the resource is namespaced, like `pvc` etc.
* `--cas-type` :- to pass the cas-type, like cstor, jiva.

### Usage
* #### `cStor`
  * #### Get `cStor` volumes
    ```bash
    $ kubectl openebs get volumes --cas-type=cstor
    NAMESPACE   NAME                                       STATUS    VERSION    CAPACITY   STORAGE CLASS         ATTACHED   ACCESS MODE      ATTACHED NODE
    cstor       pvc-193844d7-3bef-45a3-8b7d-ed3991391b45   Healthy   2.9.0      5.0 GiB    cstor-csi-sc          Bound      ReadWriteOnce    N/A
    cstor       pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc   Healthy   2.0.0      20 GiB     common-storageclass   Bound      ReadWriteOnce    node1-virtual-machine
    ```
    Note: For volumes not attached to any application, the `ATTACH NODE` would be shown as `N/A`.
  * #### Get `cStor` pools
    ```bash
    $ kubectl openebs get storage --cas-type=cstor
    NAME                      HOSTNAME                FREE     CAPACITY   READ ONLY   PROVISIONED REPLICAS   HEALTHY REPLICAS   STATUS    AGE
    cstor-storage-k5c2        node1-virtual-machine   45 GiB   45 GiB     false       1                      0                  ONLINE    10d2h
    default-cstor-disk-dcrm   node1-virtual-machine   73 GiB   90 GiB     false       7                      7                  ONLINE    27d2h
    default-cstor-disk-fp6v   node2-virtual-machine   73 GiB   90 GiB     false       7                      7                  ONLINE    27d2h
    default-cstor-disk-rhwj   node1-virtual-machine   73 GiB   90 GiB     false       7                      4                  OFFLINE   27d2h
    ```
  * #### Describe `cStor` volumes
    ```bash
    $ kubectl openebs describe volume pvc-193844d7-3bef-45a3-8b7d-ed3991391b45
  
    pvc-193844d7-3bef-45a3-8b7d-ed3991391b45 Details :
    -----------------
    NAME            : pvc-193844d7-3bef-45a3-8b7d-ed3991391b45
    ACCESS MODE     : ReadWriteOnce
    CSI DRIVER      : cstor.csi.openebs.io
    STORAGE CLASS   : cstor-csi
    VOLUME PHASE    : Released
    VERSION         : 2.9.0
    CSPC            : cstor-storage
    SIZE            : 5.0 GiB
    STATUS          : Init
    REPLICA COUNT	  : 1
    
    
    Portal Details :
    ------------------
    IQN              :  iqn.2016-09.com.openebs.cstor:pvc-193844d7-3bef-45a3-8b7d-ed3991391b45
    VOLUME NAME      :  pvc-193844d7-3bef-45a3-8b7d-ed3991391b45
    TARGET NODE NAME :  node1-virtual-machine
    PORTAL           :  10.106.27.10:3260
    TARGET IP        :  10.106.27.10
    
    
    Replica Details :
    -----------------
    NAME                                                          TOTAL    USED      STATUS    AGE
    pvc-193844d7-3bef-45a3-8b7d-ed3991391b45-cstor-storage-k5c2   72 KiB   4.8 MiB   Healthy   10d3h
    
    Cstor Completed Backup Details :
    -------------------------------
    NAME                                               BACKUP NAME   VOLUME NAME                                LAST SNAP NAME
    backup4-pvc-b026cde1-28d9-40ff-ba95-2f3a6c1d5668   backup4       pvc-193844d7-3bef-45a3-8b7d-ed3991391b45   backup4
    
    Cstor Restores Details :
    -----------------------
    NAME                                           RESTORE NAME   VOLUME NAME                                RESTORE SOURCE       STORAGE CLASS   STATUS
    backup4-3cc0839b-8428-4361-8b12-eb8509208871   backup4        pvc-193844d7-3bef-45a3-8b7d-ed3991391b45   192.168.1.165:9000   cstor-csi       0
    ```
  * #### Describe `cStor` pool
    ```bash
    $ kubectl openebs describe storage default-cstor-disk-fp6v --openebs-namespace=openebs
    
    default-cstor-disk-fp6v Details :
    ----------------
    NAME             : default-cstor-disk-fp6v
    HOSTNAME         : node1-virtual-machine
    SIZE             : 90 GiB
    FREE CAPACITY    : 73 GiB
    READ ONLY STATUS : false
    STATUS	         : ONLINE
    RAID TYPE        : stripe
    
    Blockdevice details :
    ---------------------
    NAME                                           CAPACITY   STATE
    blockdevice-8a5b69d8a2b23276f8daeac3c8179f9d   100 GiB    Active
    
    Replica Details :
    -----------------
    NAME                                                               PVC NAME   SIZE      STATE
    pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc-default-cstor-disk-fp6v   mongo      992 MiB   Healthy
    ```
* #### `Jiva`
  * #### Get `Jiva` volumes
    ```bash
    $ kubectl openebs get volumes --cas-type=jiva
    NAMESPACE   NAME                                       STATUS   VERSION   CAPACITY   STORAGE CLASS         ATTACHED   ACCESS MODE     ATTACHED NODE
    openebs     pvc-478a8329-f02d-47e5-8288-0c28b582be25   RW       2.9.0     4Gi        openebs-jiva-csi-sc   Released   ReadWriteOnce   minikube-2
    ```
    Note: For volumes not attached to any application, the `ATTACH NODE` would be shown as `N/A`.
  * #### Describe `Jiva` volumes
    ```bash
    $ kubectl openebs describe volume pvc-478a8329-f02d-47e5-8288-0c28b582be25

    pvc-478a8329-f02d-47e5-8288-0c28b582be25 Details :
    -----------------
    NAME            : pvc-478a8329-f02d-47e5-8288-0c28b582be25
    ACCESS MODE     : ReadWriteOnce
    CSI DRIVER      : jiva.csi.openebs.io
    STORAGE CLASS   : openebs-jiva-csi-sc
    VOLUME PHASE    : Released
    VERSION         : 2.9.0
    JVP             : jivavolumepolicy
    SIZE            : 4.0 GiB
    STATUS          : RW
    REPLICA COUNT	: 1
      
    Portal Details :
    ------------------
    IQN              :  iqn.2016-09.com.openebs.jiva:pvc-478a8329-f02d-47e5-8288-0c28b582be25
    VOLUME NAME      :  pvc-478a8329-f02d-47e5-8288-0c28b582be25
    TARGET NODE NAME :  node2-virtual-machine
    PORTAL           :  10.101.211.252:3260
  
    Replica Details :
    -----------------
    NAME                                                          STATUS   VOLUME                                     CAPACITY   STORAGECLASS       AGE
    openebs-pvc-478a8329-f02d-47e5-8288-0c28b582be25-jiva-rep-0   Lost     pvc-d94c2500-6ed4-44c2-ba5d-bc6aecd5cff7   4.0 GiB    openebs-hostpath   8d21h

    ```
* #### `LocalPV-LVM`
  * #### Get `LocalPV-LVM` volumes
    ```bash
    $ kubectl openebs get volumes --cas-type=lvmlocalpv
    NAMESPACE   NAME                                       STATUS   VERSION   CAPACITY   STORAGE CLASS   ATTACHED   ACCESS MODE     ATTACHED NODE
    openebs     pvc-04c2d4ea-f072-4e17-9e0a-db0fde0b2550   Ready    ci        1Gi        lvmpv-sc        Bound      ReadWriteOnce   worker-sh1
    openebs     pvc-1ec1c9b7-b74e-4742-901d-2af4558d6636   Ready    ci        1Gi        openebs-lvmpv   Bound      ReadWriteOnce   worker-sh1
    openebs     pvc-9999274f-ad01-48bc-9b21-7c51b47a870c   Ready    ci        4Gi        openebs-lvmpv   Bound      ReadWriteOnce   worker-sh1
    ```
    Note: For volumes not attached to any application, the `ATTACH NODE` would be shown as `N/A`.
  * #### Get `LocalPV-LVM` VolumeGroups
    ```bash
    $ kubectl openebs get storage --cas-type=lvmlocalpv
    NAME         FREESIZE   TOTALSIZE
    worker-sh1              
    └─lvmvg      1018 GiB   1024 GiB
    
    worker-sh2              
    └─lvmvg-1    46.7 GiB   50 GiB
    ```
* #### `LocalPV-ZFS`
  * #### Get `LocalPV-ZFS` volumes
    ```bash
    $ kubectl openebs get volumes --cas-type=zfslocalpv
    NAMESPACE   NAME                                       STATUS   VERSION   CAPACITY   STORAGE CLASS   ATTACHED   ACCESS MODE     ATTACHED NODE
    openebs     pvc-43fcbc72-a45a-49d5-9ec3-e383fcb91452   Ready    1.9.0     4Gi        openebs-zfspv   Bound      ReadWriteOnce   worker-sh1
    ```
    Note: For volumes not attached to any application, the `ATTACH NODE` would be shown as `N/A`.
  * #### Get `LocalPV-ZFS` Pools
    ```bash
    $ kubectl openebs get storage --cas-type=zfslocalpv
    NAME              FREESIZE
    node1         
    └─zfs-test-pool   32 GiB
    
    node2         
    └─zfs-test-pool   36 GiB
    ```
* #### `BlockDevice`
  * #### Get `BlockDevices` by Nodes
    ```bash
    $ kubectl openebs get bd
    NAME                                             PATH            SIZE      CLAIMSTATE   STATUS     FSTYPE       MOUNTPOINT
    minikube-2                                                                                                      
    ├─blockdevice-94312c16fb24476c3a155c34f0c211c3   /dev/sdb1       50 GiB    Unclaimed    Inactive   ext4         /var/lib/kubelet/mntpt
    └─blockdevice-94312c16fb24476c3a155c34f0c2143c   /dev/sdb1       50 GiB    Claimed      Active
    
    minikube-1                                                                                                      
    ├─blockdevice-94312c16fb24476c3a155c34f0c6153a   /dev/sdb1       50 GiB    Claimed      Inactive   zfs_member   /var/openebs/zfsvol
    ├─blockdevice-8a5b69d8a2b23276f8daeac3c8179f9d   /dev/nvme2n1    100 GiB   Claimed      Active                  
    └─blockdevice-e5a1c3c1b66c864588a66d0a7ff8ca58   /dev/nvme10n1   100 GiB   Claimed      Active
    
    minikube-3                                                                                                      
    └─blockdevice-94312c16fb24476c3a155c34f0c6199k   /dev/sdb1       50 GiB    Claimed      Active               
    ```
* #### `PersistentVolumeClaims`
  * #### Describe pvcs
    Describe any PVC using this command, it will determine the cas engine and show details accordingly.
    ```bash
    $ kubectl openebs describe pvc mongo
  
    mongo Details :
    ------------------
    NAME             : mongo
    NAMESPACE        : default
    CAS TYPE         : cstor
    BOUND VOLUME     : pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc
    ATTACHED TO NODE : node1-virtual-machine
    POOL             : default-cstor-disk
    STORAGE CLASS    : common-storageclass
    SIZE             : 20 GiB
    USED             : 1.1 GiB
    PV STATUS	       : Healthy
    
    Target Details :
    ----------------
    NAMESPACE   NAME                                                              READY   STATUS    AGE      IP           NODE
    openebs     pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc-target-7487cbc8bc5ttzl   3/3     Running   26d22h   172.17.0.7   node1-virtual-machine
    
    Replica Details :
    -----------------
    NAME                                                               TOTAL     USED      STATUS    AGE
    pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc-default-cstor-disk-dcrm   992 MiB   1.1 GiB   Healthy   26d23h
    pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc-default-cstor-disk-fp6v   992 MiB   1.1 GiB   Healthy   26d23h
    pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc-default-cstor-disk-rhwj   682 MiB   832 MiB   Offline   26d23h
    
    Additional Details from CVC :
    -----------------------------
    NAME          : pvc-b84f60ae-3f26-4110-a85d-bce7ec00dacc
    REPLICA COUNT : 3
    POOL INFO     : [default-cstor-disk-dcrm default-cstor-disk-fp6v default-cstor-disk-rhwj]
    VERSION       : 2.1.0
    UPGRADING     : true
    ```

### Contributing

OpenEBS welcomes your feedback and contributions in any form possible.

- [Join OpenEBS community on Kubernetes Slack](https://kubernetes.slack.com)
    - Already signed up? Head to our discussions at [#openebs](https://kubernetes.slack.com/messages/openebs/)
- Want to raise an issue or help with fixes and features?
    - See [open issues](https://github.com/openebs/openebs/issues)
    - See [contributing guide](./CONTRIBUTING.md)
    - See [Project Roadmap](https://github.com/openebs/openebsctl/projects/1)
    - Checkout our existing [adopters](https://github.com/openebs/openebs/tree/master/adopters) and their [feedbacks](https://github.com/openebs/openebs/issues/2719).
    - Want to join our contributor community meetings, [check this out](https://hackmd.io/mfG78r7MS86oMx8oyaV8Iw?view).
- Join our OpenEBS CNCF Mailing lists
    - For OpenEBS project updates, subscribe to [OpenEBS Announcements](https://lists.cncf.io/g/cncf-openebs-announcements)
    - For interacting with other OpenEBS users, subscribe to [OpenEBS Users](https://lists.cncf.io/g/cncf-openebs-users)


For more details checkout [CONTRIBUTING.md](./CONTRIBUTING.md).

