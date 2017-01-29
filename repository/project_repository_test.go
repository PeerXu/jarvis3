package repository

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/PeerXu/jarvis3/project"
)

func TestProjectRepository(t *testing.T) {
	Convey("Testing Project Repository", t, func() {
		ur := NewUserRepository()
		u, err := ur.LookupUserByUsername("admin")
		So(err, ShouldBeNil)

		userID := u.ID

		pr := NewProjectRepository()

		Convey("should create an executor", func() {
			e := project.NewExecutor(userID, "test", "example.com/peerxu/jarvis3", nil)
			exec, err := pr.CreateExecutor(e)
			So(err, ShouldBeNil)

			execID := exec.ID
			_, err = pr.GetExecutorByID(execID)
			So(err, ShouldBeNil)

			executors, err := pr.ListExecutors(userID)
			So(err, ShouldBeNil)
			So(len(executors), ShouldEqual, 1)

			Convey("should delete an existed executor", func() {
				err := pr.DeleteExecutorByID(execID)
				So(err, ShouldBeNil)
			})

			Convey("should create a project", func() {
				p := project.NewProject(userID, "proj0")
				proj, err := pr.CreateProject(p)
				So(err, ShouldBeNil)

				projID := proj.ID

				_, err = pr.GetProjectByID(projID)
				So(err, ShouldBeNil)

				projs, err := pr.ListProjects(userID)
				So(err, ShouldBeNil)
				So(len(projs), ShouldEqual, 1)

				Convey("should delete an existed project", func() {
					err := pr.DeleteProjectByID(projID)
					So(err, ShouldBeNil)

					_, err = pr.GetProjectByID(projID)
					So(err, ShouldEqual, project.ErrProjectNotFound)
				})

				Convey("should create a task", func() {
					t := project.NewTask(projID, execID, "t1", nil)
					task, err := pr.CreateTask(t)
					So(err, ShouldBeNil)

					taskID := task.ID

					tasks, err := pr.ListTasksByProjectID(projID)
					So(err, ShouldBeNil)
					So(len(tasks), ShouldEqual, 1)

					Convey("should get an existed task", func() {
						_, err = pr.GetTaskByID(taskID)
						So(err, ShouldBeNil)
					})

					SkipConvey("should update an existed task", func() {
						t := &project.Task{Status: project.TaskStatus_Stop}
						_, err := pr.UpdateTaskByID(taskID, t)
						So(err, ShouldBeNil)
					})

					SkipConvey("should delete an existed task", func() {
						err := pr.DeleteTaskByID(taskID)
						So(err, ShouldBeNil)

						_, err = pr.GetTaskByID(taskID)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})
}
