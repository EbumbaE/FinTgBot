package cache

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/report"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

func countExpirationByPeriod(now time.Time, period string) (expiration int64, err error) {
	_, end, err := report.GetPeriod(period)
	if err != nil {
		return 0, err
	}
	return end.Unix() - now.Unix(), nil
}

func concatKey(userID int64, period string) string {
	return strconv.FormatInt(userID, 10) + ":" + period
}

func (c *Cache) AddReportInCache(userID int64, period string, addedReport report.ReportFormat) (err error) {
	expiration, err := countExpirationByPeriod(time.Now(), period)
	if err != nil {
		return err
	}

	byteReport, err := json.Marshal(addedReport)
	if err != nil {
		return
	}

	item := &memcache.Item{
		Key:        concatKey(userID, period),
		Value:      byteReport,
		Expiration: int32(expiration),
	}
	return c.Add(item)
}

func (c *Cache) GetReportFromCache(userID int64, period string) (getReport report.ReportFormat, err error) {
	key := concatKey(userID, period)
	item, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(item.Value, &getReport)
	return
}

func (c *Cache) DeleteReportFromCache(userID int64, period string) error {
	key := concatKey(userID, period)
	return c.client.Delete(key)
}

func (c *Cache) AddNoteInCacheReports(userID int64, timeNote time.Time, note diary.Note) error {
	periods, err := report.DeterminePeriod(timeNote, time.Now())
	if err != nil {
		return err
	}
	for _, period := range periods {
		report, err := c.GetReportFromCache(userID, period)
		if err != nil {
			logger.Info("get report", zap.Error(err))
			continue
		}
		report[note.Category] += note.Sum
		if err := c.DeleteReportFromCache(userID, period); err != nil {
			logger.Error("delete report from cache", zap.Error(err))
			return err
		}
		if err := c.AddReportInCache(userID, period, report); err != nil {
			logger.Error("add report in cache", zap.Error(err))
			return err
		}
	}
	return nil
}
