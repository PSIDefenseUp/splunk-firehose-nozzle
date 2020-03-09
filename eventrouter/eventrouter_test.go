package eventrouter_test

import (
	"github.com/cloudfoundry-community/splunk-firehose-nozzle/eventrouter"
	. "github.com/cloudfoundry-community/splunk-firehose-nozzle/eventrouter"
	"github.com/cloudfoundry-community/splunk-firehose-nozzle/testing"
	"github.com/cloudfoundry/sonde-go/events"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("eventrouter", func() {

	var (
		r   Router
		err error

		origin        string
		deployment    string
		job           string
		jobIndex      string
		ip            string
		timestampNano int64
		msg           *events.Envelope
		eventType     events.Envelope_EventType

		memSink *testing.MemorySinkMock
		noCache *testing.MemoryCacheMock
	)

	BeforeEach(func() {
		noCache = testing.NewMemoryCacheMock()
		memSink = &testing.MemorySinkMock{}
		config := &Config{
			SelectedEvents: "LogMessage,HttpStart,HttpStop,HttpStartStop,ValueMetric,CounterEvent,Error,ContainerMetric",
		}
		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		timestampNano = 1467040874046121775
		deployment = "cf-warden"
		jobIndex = "85c9ff80-e99b-470b-a194-b397a6e73913"
		ip = "10.244.0.22"
		appId := "f964a41c-76ac-42c1-b2ba-663da3ec22d5"
		sourcetype := "testing"
		mtype := events.LogMessage_OUT
		logMsg := &events.LogMessage{
			Message:        []byte("testing"),
			MessageType:    &mtype,
			Timestamp:      &timestampNano,
			AppId:          &appId,
			SourceType:     &sourcetype,
			SourceInstance: &sourcetype,
		}

		msg = &events.Envelope{
			Origin:     &origin,
			EventType:  &eventType,
			Timestamp:  &timestampNano,
			Deployment: &deployment,
			Job:        &job,
			Index:      &jobIndex,
			Ip:         &ip,
			LogMessage: logMsg,
		}
	})

	It("Route valid message", func() {
		eventTypes := []events.Envelope_EventType{
			events.Envelope_LogMessage, events.Envelope_HttpStart,
			events.Envelope_HttpStop, events.Envelope_HttpStartStop,
			events.Envelope_ValueMetric, events.Envelope_CounterEvent,
			events.Envelope_Error, events.Envelope_ContainerMetric,
		}
		for i, eType := range eventTypes {
			eventType = eType
			err := r.Route(msg)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(len(memSink.Events)).To(Equal(i + 1))
			Expect(len(memSink.Messages)).To(Equal(i + 1))
		}
	})

	It("Route un-selected message", func() {
		config := &Config{
			SelectedEvents: "HttpStart",
		}
		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		eventType = events.Envelope_HttpStop
		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(0))
		Expect(len(memSink.Messages)).To(Equal(0))
	})

	It("Route default selected message", func() {
		config := &Config{
			SelectedEvents: "",
		}
		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		eventType = events.Envelope_LogMessage
		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(1))
		Expect(len(memSink.Messages)).To(Equal(1))
	})

	It("Route invalid message, no error", func() {
		eventType = events.Envelope_EventType(1000)
		err := r.Route(msg)
		// Since we will error out first
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(0))
		Expect(len(memSink.Messages)).To(Equal(0))
	})

	It("Route invalid message, error out", func() {
		invalid := events.Envelope_EventType(-1)

		// Update the map
		events.Envelope_EventType_value["invalid"] = int32(invalid)
		events.Envelope_EventType_name[int32(invalid)] = "invalid"

		config := &Config{
			SelectedEvents: "invalid",
		}
		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		eventType = invalid
		err := r.Route(msg)

		Ω(err).Should(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(0))
		Expect(len(memSink.Messages)).To(Equal(0))
	})

	It("Route ignore app", func() {
		noCache.SetIgnoreApp(true)
		eventType = events.Envelope_LogMessage
		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(0))
		Expect(len(memSink.Messages)).To(Equal(0))
	})

	It("Route whitelisted org", func() {
		orgName := "test-org-name"

		noCache.SetOrgName(orgName)

		orgIndexMappings := []eventrouter.OrgSplunkMapping {}
		orgIndexMappings = append(orgIndexMappings, eventrouter.OrgSplunkMapping { Org: orgName, DestinationIndex: "somewhere" })

		config := &Config{
			SelectedEvents: "LogMessage,HttpStart,HttpStop,HttpStartStop,ValueMetric,CounterEvent,Error,ContainerMetric",
			OrgIndexMappings: orgIndexMappings,
		}

		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(1))
		Expect(len(memSink.Messages)).To(Equal(1))
	})

	It("Route org log when no org-index mappings provided", func() {
		orgName := "test-org-name"

		noCache.SetOrgName(orgName)

		orgIndexMappings := []eventrouter.OrgSplunkMapping {}

		config := &Config{
			SelectedEvents: "LogMessage,HttpStart,HttpStop,HttpStartStop,ValueMetric,CounterEvent,Error,ContainerMetric",
			OrgIndexMappings: orgIndexMappings,
		}

		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(1))
		Expect(len(memSink.Messages)).To(Equal(1))
	})

	It("Route non-whitelisted org", func() {
		noCache.SetOrgName("some-other-org-name")

		orgIndexMappings := []eventrouter.OrgSplunkMapping {}
		orgIndexMapping := eventrouter.OrgSplunkMapping {
			Org: "someorg",
			DestinationIndex: "somewhere",
		}
		orgIndexMappings = append(orgIndexMappings, orgIndexMapping)

		config := &Config{
			SelectedEvents: "LogMessage,HttpStart,HttpStop,HttpStartStop,ValueMetric,CounterEvent,Error,ContainerMetric",
			OrgIndexMappings: orgIndexMappings,
		}

		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(0))
		Expect(len(memSink.Messages)).To(Equal(0))
	})

	It("Route whitelisted org and space", func() {
		orgName := "test-org-name"
		spaceName := "test-space-name"

		noCache.SetOrgName(orgName)
		noCache.SetSpaceName(spaceName)

		orgIndexMappings := []eventrouter.OrgSplunkMapping {}
		orgIndexMapping := eventrouter.OrgSplunkMapping {
			Org: orgName,
			DestinationIndex: "somewhere",
			Spaces: []string { "someTestSpace", spaceName, "someOtherTestSpace" },
		}
		orgIndexMappings = append(orgIndexMappings, orgIndexMapping)

		config := &Config{
			SelectedEvents: "LogMessage,HttpStart,HttpStop,HttpStartStop,ValueMetric,CounterEvent,Error,ContainerMetric",
			OrgIndexMappings: orgIndexMappings,
		}

		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(1))
		Expect(len(memSink.Messages)).To(Equal(1))
	})

	It("Route whitelisted org and non-whitelisted space", func() {
		orgName := "test-org-name"
		spaceName := "test-space-name"

		noCache.SetOrgName(orgName)
		noCache.SetSpaceName(spaceName)

		orgIndexMappings := []eventrouter.OrgSplunkMapping {}
		orgIndexMapping := eventrouter.OrgSplunkMapping {
			Org: orgName,
			DestinationIndex: "somewhere",
			Spaces: []string { "someTestSpace", "someOtherTestSpace" },
		}
		orgIndexMappings = append(orgIndexMappings, orgIndexMapping)

		config := &Config{
			SelectedEvents: "LogMessage,HttpStart,HttpStop,HttpStartStop,ValueMetric,CounterEvent,Error,ContainerMetric",
			OrgIndexMappings: orgIndexMappings,
		}

		r, err = New(noCache, memSink, config)
		Ω(err).ShouldNot(HaveOccurred())

		err := r.Route(msg)
		Ω(err).ShouldNot(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(0))
		Expect(len(memSink.Messages)).To(Equal(0))
	})

	It("Route sink error", func() {
		memSink.ReturnErr = true
		eventType = events.Envelope_LogMessage
		err := r.Route(msg)
		Ω(err).Should(HaveOccurred())
		Expect(len(memSink.Events)).To(Equal(0))
		Expect(len(memSink.Messages)).To(Equal(0))
	})

	It("Invalid event", func() {
		config := &Config{
			SelectedEvents: "invalid-event",
		}
		_, err = New(noCache, memSink, config)
		Ω(err).Should(HaveOccurred())
	})
})
