package testing

import (
	"github.com/cloudfoundry-community/splunk-firehose-nozzle/cache"
)

type MemoryCacheMock struct {
	ignoreApp bool
	orgName string
	spaceName string
}

func NewMemoryCacheMock() *MemoryCacheMock {
	return &MemoryCacheMock {
		orgName: "testing-org",
		spaceName: "testing-space",
	}
}

func (c *MemoryCacheMock) Open() error {
	return nil
}

func (c *MemoryCacheMock) Close() error {
	return nil
}

func (c *MemoryCacheMock) GetAllApps() (map[string]*cache.App, error) {
	return nil, nil
}

func (c *MemoryCacheMock) GetApp(appGuid string) (*cache.App, error) {
	app := &cache.App{
		Name:       "testing-app",
		Guid:       "f964a41c-76ac-42c1-b2ba-663da3ec22d5",
		SpaceName:  c.spaceName,
		SpaceGuid:  "f964a41c-76ac-42c1-b2ba-663da3ec22d6",
		OrgName:    c.orgName,
		OrgGuid:    "f964a41c-76ac-42c1-b2ba-663da3ec22d7",
		IgnoredApp: c.ignoreApp,
	}

	return app, nil
}

func (c *MemoryCacheMock) SetIgnoreApp(ignore bool) {
	c.ignoreApp = ignore
}

func (c *MemoryCacheMock) SetOrgName(orgName string) {
	c.orgName = orgName
}

func (c *MemoryCacheMock) SetSpaceName(spaceName string) {
	c.spaceName = spaceName
}