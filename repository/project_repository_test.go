package repository

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"

	"github.com/PeerXu/jarvis3/project"
)

func TestProjectRepository(t *testing.T) {
	c.Convey("Testing Project Repository", t, func() {
		r := NewProjectRepository()

		c.Convey("should create an executor", func() {
			e := project.NewExecutor("test", "example.com/peerxu/jarvis3", "admin", nil)
			_, err := r.CreateExecutor(e)
			c.So(err, c.ShouldBeNil)
			_, err = r.GetExecutor("admin", "test")
			c.So(err, c.ShouldBeNil)
			executors, err := r.ListExecutors("admin")
			c.So(err, c.ShouldBeNil)
			c.So(len(executors), c.ShouldEqual, 1)

			c.Convey("should delete an existed executor", func() {
				err := r.DeleteExecutor("admin", "test")
				c.So(err, c.ShouldBeNil)
			})

			c.Convey("should create a project", func() {
				p := project.NewProject("proj0", "admin")
				_, err := r.CreateProject(p)
				c.So(err, c.ShouldBeNil)
				_, err = r.GetProject("admin", "proj0")
				c.So(err, c.ShouldBeNil)
				projs, err := r.ListProjects("admin")
				c.So(err, c.ShouldBeNil)
				c.So(len(projs), c.ShouldEqual, 1)

				c.Convey("should delete an existed project", func() {
					err := r.DeleteProject("admin", "proj0")
					c.So(err, c.ShouldBeNil)

					_, err = r.GetProject("admin", "proj0")
					c.So(err, c.ShouldEqual, project.ErrProjectNotFound)
				})

				c.Convey("should create a job", func() {
					j := project.NewJob("j1", e, nil)
					_, err := r.CreateJob("admin", "proj0", j)
					c.So(err, c.ShouldBeNil)

					jobs, err := r.ListJobs("admin", "proj0")
					c.So(err, c.ShouldBeNil)
					c.So(len(jobs), c.ShouldEqual, 1)

					c.Convey("should update an existed job", func() {
						j := &project.Job{Status: project.JobStatus_Stop}
						_, err := r.UpdateJob("admin", "j1", "proj0", j)
						c.So(err, c.ShouldBeNil)
					})
				})
			})
		})
	})
}
