package db_test

import (
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {
	var pipeline db.Pipeline

	BeforeEach(func() {
		var (
			created bool
			err     error
		)

		pipeline, created, err = defaultTeam.SavePipeline(
			"pipeline-with-resources",
			atc.Config{
				Resources: atc.ResourceConfigs{
					{
						Name:    "some-resource",
						Type:    "registry-image",
						Source:  atc.Source{"some": "repository"},
						Version: atc.Version{"ref": "abcdef"},
					},
					{
						Name:   "some-other-resource",
						Type:   "git",
						Source: atc.Source{"some": "other-repository"},
					},
					{
						Name:   "some-secret-resource",
						Type:   "git",
						Source: atc.Source{"some": "((secret-repository))"},
					},
					{
						Name:         "some-resource-custom-check",
						Type:         "git",
						Source:       atc.Source{"some": "some-repository"},
						CheckEvery:   "10ms",
						CheckTimeout: "1m",
					},
				},
			},
			0,
			db.PipelineUnpaused,
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(created).To(BeTrue())
	})

	Describe("(Pipeline).Resources", func() {
		var resources []db.Resource

		JustBeforeEach(func() {
			var err error
			resources, err = pipeline.Resources()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns the resources", func() {
			Expect(resources).To(HaveLen(4))

			ids := map[int]struct{}{}

			for _, r := range resources {
				ids[r.ID()] = struct{}{}

				switch r.Name() {
				case "some-resource":
					Expect(r.Type()).To(Equal("registry-image"))
					Expect(r.Source()).To(Equal(atc.Source{"some": "repository"}))
					Expect(r.PinnedVersion()).To(Equal(atc.Version{"ref": "abcdef"}))
				case "some-other-resource":
					Expect(r.Type()).To(Equal("git"))
					Expect(r.Source()).To(Equal(atc.Source{"some": "other-repository"}))
				case "some-secret-resource":
					Expect(r.Type()).To(Equal("git"))
					Expect(r.Source()).To(Equal(atc.Source{"some": "((secret-repository))"}))
				case "some-resource-custom-check":
					Expect(r.Type()).To(Equal("git"))
					Expect(r.Source()).To(Equal(atc.Source{"some": "some-repository"}))
					Expect(r.CheckEvery()).To(Equal("10ms"))
					Expect(r.CheckTimeout()).To(Equal("1m"))
				}
			}
		})
	})

	Describe("(Pipeline).Resource", func() {
		var (
			err      error
			found    bool
			resource db.Resource
		)

		Context("when the resource exists", func() {
			JustBeforeEach(func() {
				resource, found, err = pipeline.Resource("some-resource")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns the resource", func() {
				Expect(found).To(BeTrue())
				Expect(resource.Name()).To(Equal("some-resource"))
				Expect(resource.Type()).To(Equal("registry-image"))
				Expect(resource.Source()).To(Equal(atc.Source{"some": "repository"}))
			})
		})

		Context("when the resource does not exist", func() {
			JustBeforeEach(func() {
				resource, found, err = pipeline.Resource("bonkers")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns nil", func() {
				Expect(found).To(BeFalse())
				Expect(resource).To(BeNil())
			})
		})
	})

	Describe("Pause", func() {
		var (
			resource db.Resource
			err      error
			found    bool
		)

		JustBeforeEach(func() {
			resource, found, err = pipeline.Resource("some-resource")
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(resource.Paused()).To(BeFalse())
		})

		It("pauses the resource", func() {
			err = resource.Pause()
			Expect(err).ToNot(HaveOccurred())

			found, err = resource.Reload()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(resource.Paused()).To(BeTrue())
		})
	})

	Describe("Unpause", func() {
		var (
			resource db.Resource
			err      error
			found    bool
		)

		JustBeforeEach(func() {
			resource, found, err = pipeline.Resource("some-resource")
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())

			err = resource.Pause()
			Expect(err).ToNot(HaveOccurred())

			found, err = resource.Reload()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(resource.Paused()).To(BeTrue())
		})

		It("pauses the resource", func() {
			err = resource.Unpause()
			Expect(err).ToNot(HaveOccurred())

			found, err = resource.Reload()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(resource.Paused()).To(BeFalse())
		})
	})

})
