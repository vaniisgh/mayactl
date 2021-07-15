# Release v0.2.0
<img width="300" align="right" alt="OpenEBS Logo" src="https://raw.githubusercontent.com/cncf/artwork/master/projects/openebs/stacked/color/openebs-stacked-color.png" xmlns="http://www.w3.org/1999/html">

## Notes
Users are encouraged to install this tool, use it and help us know what can be better.<br/>
Thank you to all that contributed with flushing out issues with OpenEBS-CTL!<br/>
You can checkout the [documentation](https://github.com/openebs/openebsctl#readme) for more information.<br/>

## Notable Changes
We have added support for `Jiva`, `LocalPV-LVM`, `LocalPV-ZFS` storage engine.<br/>
* Jiva `volume` listing and describing commands.
* LocalPV-LVM `volume` and `volumegroups` listing commands.
* LocalPV-ZFS `volume` and `pools` listing commands.
* BlockDevices listing by Nodes
* PersistentVolumeClaim describe now supports jiva as well.

## What's Next
* Adding support for Mayastor
* Adding support for managing Storage Devices
* Adding support for performing sanity checks and flagging discrepancies like listing stale volumes or over-utilized pools.
* Adding support for getting overall status (like kubectl openebs cluster-info)
* Ability to generate raise GitHub issues with required troubleshooting information.
* Ability to tell why a cStor volume is not ready.
* Adding support for upgrading pool and volumes.

## Resolved Bugs

+ [[Issue 40]](https://github.com/openebs/openebsctl/issues/40) Install openebsctl using a script.
+ [[Issue 35]](https://github.com/openebs/openebsctl/issues/35) Support for Jiva storage engine.
+ [[Issue 32]](https://github.com/openebs/openebsctl/issues/32) Refactoring of the code and cleanup.

---

# Release v0.1.0
<img width="300" align="right" alt="OpenEBS Logo" src="https://raw.githubusercontent.com/cncf/artwork/master/projects/openebs/stacked/color/openebs-stacked-color.png" xmlns="http://www.w3.org/1999/html">

## Notes
OpenEBS-CTL v0.1.0 is a first feature release.<br/> 
Users are encouraged to install this tool, use it and help us know what can be better.<br/>
Thank you to all that contributed with flushing out issues with OpenEBS-CTL!<br/>
You can checkout the [documentation](https://github.com/openebs/openebsctl#readme) for more information.<br/>

## Notable Changes
We have added support for `cStor` storage engine.<br/>
* cStor `volume` and `pools` listing commands
* cStor `volume` and `pools` describing commands
* `PersistentVolumeClaim` describe command

## What's Next
* Restructuring of code, to add supoort for other Storage Engines
* Adding support for Jiva
* Adding support for Local PV
* Adding support for Mayastor
* Adding support for managing Storage Devices
* Adding support for performing sanity checks and flagging discrepancies like listing stale volumes or over-utilized pools.
* Adding support for getting overall status (like kubectl openebs cluster-info)
* Ability to generate raise GitHub issues with required troubleshooting information.

## Resolved Bugs

+ [[Issue 24]](https://github.com/openebs/openebsctl/issues/24) Convert object blocks into object list
+ [[Issue 18]](https://github.com/openebs/openebsctl/issues/18) Determine openebs ns from the CLI.
+ [[Issue 15]](https://github.com/openebs/openebsctl/issues/15) Support command openebsctl pool describe
+ [[Issue 8]](https://github.com/openebs/openebsctl/issues/8) Add the ability to specify KUBECONFIG variable in openebsctl enhancement good first issue
+ [[Issue 1]](https://github.com/openebs/openebsctl/issues/1) Missed case for listing cStor volumes cStor volume command
