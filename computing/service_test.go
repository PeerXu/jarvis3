package computing

import (
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	c "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	"github.com/PeerXu/jarvis3/project"
	"github.com/PeerXu/jarvis3/repository"
	"github.com/PeerXu/jarvis3/signing"
	"github.com/PeerXu/jarvis3/user"
)

func TestComputingService(t *testing.T) {
	users := repository.NewUserRepository()
	projects := repository.NewProjectRepository()
	logger := log.NewLogfmtLogger(os.Stderr)
	ss := signing.NewService(logger, users)
	cs := NewService(logger, projects)
	ctx := context.Background()

	at, _ := ss.Login(ctx, "admin", "admin")
	jctx := jcontext.NewContext(&user.User{Username: "admin"}, at)
	ctx = context.WithValue(ctx, "JarvisContext", jctx)

	c.Convey("create an executor", t, func() {
		_, err := cs.CreateExecutor(ctx, "demo", "github.com/PeerXu/jarvis3/examples/executors/demo", nil)
		c.So(err, c.ShouldBeNil)

		_, err = cs.GetExecutor(ctx, "demo")
		c.So(err, c.ShouldBeNil)

		es, err := cs.ListExecutors(ctx)
		c.So(err, c.ShouldBeNil)
		c.So(len(es), c.ShouldEqual, 1)

		c.Convey("create a project", func() {
			_, err := cs.CreateProject(ctx, "project1")
			c.So(err, c.ShouldBeNil)

			_, err = cs.GetProject(ctx, "project1")
			c.So(err, c.ShouldBeNil)

			ps, err := cs.ListProjects(ctx)
			c.So(err, c.ShouldBeNil)
			c.So(len(ps), c.ShouldEqual, 1)

			c.Convey("create a job", func() {
				_, err := cs.CreateJob(ctx, "job1", "project1", "demo", nil)
				c.So(err, c.ShouldBeNil)

				p, err := cs.GetProject(ctx, "project1")
				c.So(err, c.ShouldBeNil)
				c.So(len(p.Jobs), c.ShouldEqual, 1)

				c.Convey("update an existed job", func() {
					j := &project.Job{Status: project.JobStatus_Stop}
					_, err = cs.UpdateJob(ctx, "job1", "project1", j)
					c.So(err, c.ShouldBeNil)
				})
			})

			c.Convey("delete an existed project", func() {
				err = cs.DeleteProject(ctx, "project1")
				c.So(err, c.ShouldBeNil)
			})
		})

		c.Convey("delete an existed executor", func() {
			err := cs.DeleteExecutor(ctx, "demo")
			c.So(err, c.ShouldBeNil)
		})
	})
}
