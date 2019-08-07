package fetchers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"code.cloudfoundry.org/log-cache/pkg/rpc/logcache_v1"
)

type CurrentUsageFetcher struct {
	client LogCacheClient
}

func NewCurrentUsageFetcher(client LogCacheClient) *CurrentUsageFetcher {
	return &CurrentUsageFetcher{client: client}
}

func (f CurrentUsageFetcher) FetchInstanceData(appGUID string) (map[int][]InstanceData, error) {
	res, err := f.client.PromQL(
		context.Background(),
		fmt.Sprintf(`idelta(absolute_usage{source_id="%s"}[1m]) / idelta(absolute_entitlement{source_id="%s"}[1m])`, appGUID, appGUID),
	)
	if err != nil {
		return nil, err
	}

	return parseCurrentUsage(res), nil
}

func parseCurrentUsage(res *logcache_v1.PromQL_InstantQueryResult) map[int][]InstanceData {
	usagePerInstance := map[int][]InstanceData{}
	for _, sample := range res.GetVector().GetSamples() {
		instanceID, err := strconv.Atoi(sample.GetMetric()["instance_id"])
		if err != nil {
			continue
		}
		timestamp, err := strconv.ParseFloat(sample.GetPoint().GetTime(), 64)
		if err != nil {
			continue
		}

		usagePerInstance[instanceID] = []InstanceData{{
			InstanceID: instanceID,
			Time:       time.Unix(int64(timestamp), 0),
			Value:      sample.GetPoint().GetValue(),
		}}
	}

	return usagePerInstance
}
