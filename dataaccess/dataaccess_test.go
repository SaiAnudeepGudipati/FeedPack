package dataaccess

import (
	"enterpret/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataStore_FetchFeedbacks(t *testing.T) {
	ds := NewDataStore()
	assert.NotNil(t, ds)

	err := ds.Store(models.FeedbackIngest{
		Meta: models.Meta{Tenant: "tenant", Source: "source", ID: "123"},
	})
	assert.Nil(t, err)

	_, err = ds.FetchFeedbacks("tenant", 1, 10, "source")
	assert.Nil(t, err)

	_, err = ds.FetchFeedbacks("", 1, 10, "source")
	assert.NotNil(t, err)
}

func TestDataStore_Store(t *testing.T) {
	ds := NewDataStore()
	assert.NotNil(t, ds)

	err := ds.Store(models.FeedbackIngest{
		Meta: models.Meta{Tenant: "tenant", Source: "source", ID: "123"},
	})
	assert.Nil(t, err)

	err = ds.Store(models.FeedbackIngest{
		Meta: models.Meta{Tenant: "tenant", Source: "source", ID: "123"},
	})
	assert.Nil(t, err)
}

func TestNewDataStore(t *testing.T) {
	ds := NewDataStore()
	assert.NotNil(t, ds)
}
