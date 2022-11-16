[![CircleCI](https://circleci.com/gh/giantswarm/capg-garbage-collector.svg?style=shield)](https://circleci.com/gh/giantswarm/capg-garbage-collector)

# capg-garbage-collector

Clean up leftover resources caused by bugs or unexpected situations

| Resource Name      | Reason |
| :---        |    :----:   |
| MachinePool      | !machinePool.DeletionTimestamp.IsZero()       |
