[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/capg-garbage-collector/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/capg-garbage-collector/tree/main)

# capi-garbage-collector

Clean up leftover resources caused by bugs or unexpected situations

| Resource Name      | Reason |
| :---        |    :----:   |
| MachinePool      | !machinePool.DeletionTimestamp.IsZero()       |
