package models

import (
	"fmt"
	"strconv"
	"time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var DefaultTaskList *TaskManager

type Task struct {
	Id        int64
	ID        int64  // Unique identifier
	Title     string // Description
	Done      bool   // Is this task done?
	Content   string `orm:"sie(100)"`
	UserName  string
	UserId    string
	State     int64     //0表示计划任务，1表示进行中任务，2表示完成任务
	StartTime time.Time `orm:"null;auto_now_add;type(datetime)"`
	EndTime   time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created   time.Time `orm:"null;auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"null;auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Task))
}

// func DeleteTask(tid string) error {

// }

func AddTask(title, content string, starttime, endtime time.Time) error {
	// tidNum, err := strconv.ParseInt(tid, 10, 64)
	// if err != nil {
	// 	return err
	// }
	task := &Task{
		Title:     title,
		Content:   content,
		StartTime: starttime,
		EndTime:   endtime,
		State:     0,
		Created:   time.Now(),
	}
	o := orm.NewOrm()
	_, err := o.Insert(task)
	if err != nil {
		return err
	}
	return err
}

func GetAllTasks() (tasks, tasks1, tasks2 []*Task, err error) {
	tasks = make([]*Task, 0)
	tasks1 = make([]*Task, 0)
	tasks2 = make([]*Task, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("task")

	_, err = qs.OrderBy("-created").Filter("state", 0).All(&tasks)
	_, err = qs.OrderBy("-created").Filter("state", 1).All(&tasks1)
	_, err = qs.OrderBy("-created").Filter("state", 2).All(&tasks2)
	return tasks, tasks1, tasks2, err
}

func GetDetails(tid string) (task []*Task, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	// task = &Task{Id: tidNum}
	o := orm.NewOrm()
	qs := o.QueryTable("task")
	// var oldtidNum int64
	err = qs.Filter("Id", tidNum).One(&task) //这里Id区分大小写，奇怪，大概是因为struct有2个id的缘故吧？？
	return task, err
}

func DeleteTask(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	// var oldtidNum int64
	task := &Task{Id: tidNum}
	if o.Read(task) == nil {
		// oldtidNum = task.Id
		_, err = o.Delete(task)
		if err != nil {
			return err
		}
	}
	return err
}
func UpdateTaskstate(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	task := &Task{Id: tidNum}

	if o.Read(task) == nil {
		task.State = task.State + 1
		task.Updated = time.Now()
		_, err = o.Update(task)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateTaskstate1(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	task := &Task{Id: tidNum}

	if o.Read(task) == nil {
		task.State = task.State - 1
		task.Updated = time.Now()
		_, err = o.Update(task)
		if err != nil {
			return err
		}
	}
	return nil
}

// var oldCateId int64
// o := orm.NewOrm()
// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
// user := User{Name: "slene"}
// err = o.Read(&user, "Name")

// topic := Topic{Id: tidNum}
// if o.Read(&topic) == nil {
// 	oldCateId = topic.CategoryId
// 	_, err = o.Delete(&topic)
// 	if err != nil {
// 		return err
// 	}
// }
// NewTask creates a new task given a title, that can't be empty.
// func NewTask(title string) (*Task, error) {
// 	if title == "" {
// 		return nil, fmt.Errorf("empty title")
// 	}
// 	return &Task{0, title, false}, nil
// }

// TaskManager manages a list of tasks in memory.
type TaskManager struct {
	tasks  []*Task
	lastID int64
}

// NewTaskManager returns an empty TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Save(task *Task) error {
	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		m.tasks = append(m.tasks, cloneTask(task))
		return nil
	}

	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

// All returns the list of all the Tasks in the TaskManager.
func (m *TaskManager) All() []*Task {
	return m.tasks
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Find(ID int64) (*Task, bool) {
	for _, t := range m.tasks {
		if t.ID == ID {
			return t, true
		}
	}
	return nil, false
}

func init() {
	DefaultTaskList = NewTaskManager()
}
