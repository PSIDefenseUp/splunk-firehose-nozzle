package eventrouter

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry-community/splunk-firehose-nozzle/cache"
	fevents "github.com/cloudfoundry-community/splunk-firehose-nozzle/events"
	"github.com/cloudfoundry-community/splunk-firehose-nozzle/eventsink"
	"github.com/cloudfoundry/sonde-go/events"
)

type OrgSplunkMappingConfig struct {
	Mappings []OrgSplunkMapping `yaml:"mappings"`
}

type OrgSplunkMapping struct {
	Org string `yaml:"org"`
	Spaces []string `yaml:"spaces"`
	DestinationIndex string `yaml:"destination_index"`
}

type Config struct {
	SelectedEvents        string
	OrgIndexMappings      []OrgSplunkMapping
	BlacklistNonAppEvents bool
}

type router struct {
	appCache              cache.Cache
	sink                  eventsink.Sink
	selectedEvents        map[string]bool
	orgIndexMappings      []OrgSplunkMapping
	blacklistNonAppEvents bool
}

func New(appCache cache.Cache, sink eventsink.Sink, config *Config) (Router, error) {
	selectedEvents, err := fevents.ParseSelectedEvents(config.SelectedEvents)

	if err != nil {
		return nil, err
	}

	return &router{
		appCache:       appCache,
		sink:           sink,
		selectedEvents: selectedEvents,
		orgIndexMappings: config.OrgIndexMappings,
		blacklistNonAppEvents: config.BlacklistNonAppEvents,
	}, nil
}

func (r *router) Route(msg *events.Envelope) error {
	eventType := msg.GetEventType()

	if _, ok := r.selectedEvents[eventType.String()]; !ok {
		// Ignore this event since we are not interested
		return nil
	}

	var event *fevents.Event
	switch eventType {
	case events.Envelope_HttpStartStop:
		event = fevents.HttpStartStop(msg)
	case events.Envelope_LogMessage:
		event = fevents.LogMessage(msg)
	case events.Envelope_ValueMetric:
		event = fevents.ValueMetric(msg)
	case events.Envelope_CounterEvent:
		event = fevents.CounterEvent(msg)
	case events.Envelope_Error:
		event = fevents.ErrorEvent(msg)
	case events.Envelope_ContainerMetric:
		event = fevents.ContainerMetric(msg)

	default:
		return fmt.Errorf("Unsupported event type: %s", eventType.String())
	}

	event.AnnotateWithEnvelopeData(msg)
	event.AnnotateWithCFMetaData()

	if _, hasAppId := event.Fields["cf_app_id"]; hasAppId {
		event.AnnotateWithAppData(r.appCache)
	}

	var appIdStr string
	hasAppIdStr := false
	if appId, hasAppId := event.Fields["cf_app_id"]; hasAppId {
		appIdStr, hasAppIdStr = appId.(string)
	}

	if hasAppIdStr && appIdStr != "" {
		event.AnnotateWithAppData(r.appCache)
	} else if r.blacklistNonAppEvents{
		return nil
	}

	if r.orgIndexMappings != nil && len(r.orgIndexMappings) > 0 {
		var orgName string
		hasOrgNameString := false
		org, hasOrgName := event.Fields["cf_org_name"]
		if (hasOrgName) {
			orgName, hasOrgNameString = org.(string)
		}

		var spaceName string
		hasSpaceNameString := false
		space, hasSpaceName := event.Fields["cf_space_name"]
		if (hasSpaceName) {
			spaceName, hasSpaceNameString = space.(string)
		}

		if (hasOrgNameString && hasSpaceNameString) {
			mappedSplunkIndex := r.getOrgSplunkIndex(orgName, spaceName)
			if mappedSplunkIndex == nil {
				return nil
			}

			event.Fields["info_splunk_index"] = *mappedSplunkIndex
		}
	}

	if ignored, ok := event.Fields["cf_ignored_app"]; ok {
		if ignoreApp, ok := ignored.(bool); ok && ignoreApp {
			// Ignore events from this app since end user tag to ignore this app
			return nil
		}
	}

	err := r.sink.Write(event.Fields, event.Msg)
	if err != nil {
		fields := map[string]interface{}{"err": fmt.Sprintf("%s", err)}
		r.sink.Write(fields, "Failed to write events")
	}
	return err
}

func (r *router) getOrgSplunkIndex(org string, space string) *string {
	for _, mapping := range r.orgIndexMappings {
		if strings.EqualFold(org, mapping.Org) {
			if mapping.Spaces == nil || len(mapping.Spaces) == 0 {
				return &mapping.DestinationIndex
			}

			for _, mappingSpace := range mapping.Spaces {
				if strings.EqualFold(space, mappingSpace) {
					return &mapping.DestinationIndex
				}
			}
		}
	}

	return nil
}
