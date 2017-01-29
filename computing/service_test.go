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
)

func TestComputingService(t *testing.T) {
	c.Convey("create an executor", t, func() {
		users := repository.NewUserRepository()
		projects := repository.NewProjectRepository()
		logger := log.NewLogfmtLogger(os.Stderr)
		ss := signing.NewService(logger, users)
		cs := NewService(logger, projects)
		ctx := context.Background()

		u, _ := ss.Login(ctx, "admin", "admin")
		at := u.AccessTokens[0]
		jctx := jcontext.NewContext(u, at)
		ctx = context.WithValue(ctx, "JarvisContext", jctx)

		exec, err := cs.CreateExecutor(ctx, "demo", "github.com/PeerXu/jarvis3/examples/executors/demo", nil)
		c.So(err, c.ShouldBeNil)

		execID := exec.ID
		_, err = cs.GetExecutorByID(ctx, execID)
		c.So(err, c.ShouldBeNil)

		es, err := cs.ListExecutors(ctx)
		c.So(err, c.ShouldBeNil)
		c.So(len(es), c.ShouldEqual, 1)

		c.Convey("create a project", func() {
			p, err := cs.CreateProject(ctx, "project1")
			c.So(err, c.ShouldBeNil)

			projID := p.ID

			_, err = cs.GetProjectByID(ctx, projID)
			c.So(err, c.ShouldBeNil)

			ps, err := cs.ListProjects(ctx)
			c.So(err, c.ShouldBeNil)
			c.So(len(ps), c.ShouldEqual, 1)

			c.Convey("create a task", func() {
				t, err := cs.CreateTask(ctx, projID, execID, "task1", nil)
				c.So(err, c.ShouldBeNil)

				taskID := t.ID

				t, err = cs.GetTaskByID(ctx, taskID)
				c.So(err, c.ShouldBeNil)

				p, err := cs.GetProjectByID(ctx, projID)
				c.So(err, c.ShouldBeNil)
				c.So(len(p.Tasks), c.ShouldEqual, 1)

				c.Convey("pop a ready task", func() {
					_, err := cs.PopReadyTask(ctx)
					c.So(err, c.ShouldBeNil)

					c.Convey("update an existed task", func() {
						t := &project.Task{Status: project.TaskStatus_Stop}
						_, err = cs.UpdateTaskByID(ctx, taskID, t)
						c.So(err, c.ShouldBeNil)
					})
				})

				c.Convey("get an existed task", func() {
					_, err := cs.GetTaskByID(ctx, taskID)
					c.So(err, c.ShouldBeNil)
				})
			})

			c.Convey("delete an existed project", func() {
				err = cs.DeleteProjectByID(ctx, projID)
				c.So(err, c.ShouldBeNil)
			})
		})

		c.Convey("delete an existed executor", func() {
			err := cs.DeleteExecutorByID(ctx, execID)
			c.So(err, c.ShouldBeNil)
		})
	})
}
