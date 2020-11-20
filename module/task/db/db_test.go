package db

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/ericpai/please/database"
	"github.com/ericpai/please/module/task"
	"github.com/ericpai/please/test"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/dig"
)

func TestTaskDB(t *testing.T) {

	Convey("test defaultDB1", t, func() {
		c := dig.New()
		database.Init(c)
		testDB := New(test.InitDB()).(*defaultDB)
		ctx := context.Background()
		testObj := []task.PO{
			{
				Address:     "192.168.1.1:4444",
				SourcePath:  "E:\\logs\\host.log",
				DestPath:    "F:\\logs\\host.log",
				Backend:     task.BackendWindows,
				CreatedTime: time.Now().UTC(),
			},
			{
				Address:     "192.168.1.1:4444",
				SourcePath:  "/var/log/host.log",
				DestPath:    "F:\\logs\\host.log",
				CreatedTime: time.Now().UTC(),
			},
		}
		var outputObj []task.PO

		for _, v := range testObj {
			res, err := testDB.Insert(ctx, v)
			So(err, ShouldBeNil)
			So(res.Address, ShouldEqual, v.Address)
			So(res.Backend, ShouldEqual, v.Backend)
			So(res.SourcePath, ShouldEqual, v.SourcePath)
			So(res.DestPath, ShouldEqual, v.DestPath)
			So(res.CreatedTime, ShouldEqual, v.CreatedTime)
			So(res.UpdatedTime, ShouldEqual, v.UpdatedTime)
			So(res.Succeed, ShouldEqual, v.Succeed)
			So(res.Enabled, ShouldEqual, v.Enabled)
			So(res.ID, ShouldNotBeZeroValue)
			outputObj = append(outputObj, res)
		}

		oss, err := testDB.SelectAll(ctx)
		So(err, ShouldBeNil)
		sort.Slice(oss, func(i, j int) bool {
			return oss[i].ID < oss[j].ID
		})
		for i, v := range outputObj {
			res := oss[i]
			So(res.Address, ShouldEqual, v.Address)
			So(res.Backend, ShouldEqual, v.Backend)
			So(res.SourcePath, ShouldEqual, v.SourcePath)
			So(res.DestPath, ShouldEqual, v.DestPath)
			So(res.CreatedTime, ShouldEqual, v.CreatedTime)
			So(res.UpdatedTime, ShouldEqual, v.UpdatedTime)
			So(res.Succeed, ShouldEqual, v.Succeed)
			So(res.Enabled, ShouldEqual, v.Enabled)
			So(res.Enabled, ShouldEqual, v.Enabled)
			So(res.ID, ShouldEqual, v.ID)
		}

		v := outputObj[0]
		v.Succeed = true
		v.UpdatedTime = time.Now().UTC()
		newV, err := testDB.Update(ctx, v)
		So(err, ShouldBeNil)
		So(newV.Address, ShouldEqual, v.Address)
		So(newV.Backend, ShouldEqual, v.Backend)
		So(newV.SourcePath, ShouldEqual, v.SourcePath)
		So(newV.DestPath, ShouldEqual, v.DestPath)
		So(newV.CreatedTime, ShouldEqual, v.CreatedTime)
		So(newV.UpdatedTime, ShouldEqual, v.UpdatedTime)
		So(newV.Succeed, ShouldEqual, v.Succeed)
		So(newV.Enabled, ShouldEqual, v.Enabled)
		So(newV.ID, ShouldEqual, v.ID)

		selectedNewV, err := testDB.SelectByID(ctx, newV.ID)
		So(err, ShouldBeNil)
		So(selectedNewV.Address, ShouldEqual, newV.Address)
		So(selectedNewV.Backend, ShouldEqual, newV.Backend)
		So(selectedNewV.SourcePath, ShouldEqual, newV.SourcePath)
		So(selectedNewV.DestPath, ShouldEqual, newV.DestPath)
		So(selectedNewV.CreatedTime, ShouldEqual, newV.CreatedTime)
		So(selectedNewV.UpdatedTime, ShouldEqual, newV.UpdatedTime)
		So(selectedNewV.Succeed, ShouldEqual, newV.Succeed)
		So(selectedNewV.Enabled, ShouldEqual, v.Enabled)
		So(selectedNewV.ID, ShouldEqual, v.ID)

		err = testDB.Delete(ctx, newV.ID)
		So(err, ShouldBeNil)
		oss, err = testDB.SelectAll(ctx)
		So(err, ShouldBeNil)
		So(oss, ShouldHaveLength, 1)
	})

}
