package main

import (
	"fmt"
	"container/list"
	"github.com/kataras/iris/middleware/recover"
	c "content"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"strconv"
	"execution"
	"reflect"
	"github.com/kataras/iris/context"
	"deepcopy"
	"github.com/ulule/deepcopier"
)

var gameState = new(c.GameState)
var undoHistory = list.New()

func saveForUndo() {
	var tmpThreadCtx = []execution.ThreadContext{}

	for _, tc := range gameState.ThreadContexts {
		ctx := &execution.ThreadContext{}
		deepcopier.Copy(tc).To(ctx)
		tmpThreadCtx = append(tmpThreadCtx, *ctx)
	}

	var history = c.History{tmpThreadCtx, deepcopy.Copy(*gameState.GlobalState).(execution.GlobalContext)}
	undoHistory.PushBack(history)
}

//-1失败、0正常、1成功
var runState = 0
//TODO
func startLevel(levelId int) {
	//清空
	for e := undoHistory.Front(); e != nil; e = e.Next() {
		undoHistory.Remove(e)
	}
	gameState = new(c.GameState)
	// 开始
	var level = c.GetLevel(levelId)
	gameState.GlobalState = level.GlobalContext
	gameState.ThreadContexts = level.ThreadContexts
	gameState.Level = *level

}

func areAllThreadsBlocked() bool {
	for _, ctx := range gameState.ThreadContexts {
		if !IsThreadBlocked(ctx.Id) {
			return false
		}
	}
	return true
}

func areAllThreadsFinished() bool {
	finished := true
	for _, ctx := range gameState.ThreadContexts {
		finished = finished && IsThreadFinished(ctx.Id)
	}
	return finished
}

func checkForVictoryConditions() {
	var howManyCriticalSections = 0

	for threadId, t := range gameState.Level.ThreadContexts {

		if IsThreadFinished(threadId) {
			continue
		}
		var thread = gameState.Level.ThreadContexts[threadId]
		var instructions = *thread.Instructions
		var threadState = gameState.ThreadContexts[threadId]
		var programCounter = threadState.ProgramCounter
		var currentInstruction = instructions[programCounter]
		fmt.Print("CurIns:", currentInstruction, ";Reflect:", reflect.TypeOf(currentInstruction))
		//TODO
		if reflect.TypeOf(currentInstruction) == reflect.TypeOf(&execution.CriticalSectionExpression{}) {
			howManyCriticalSections++
		}
		fmt.Println(t)
	}

	if howManyCriticalSections >= 2 {
		win("More than one threads were in a critical section at the same time.")
		return
	}

	if gameState.GlobalState.IsPanic {
		win("Program entered panic instruction.")
		return
	}

	if areAllThreadsBlocked() {
		win("A deadlock occurred - all threads were blocked simultaneously.")
		return
	}

	if areAllThreadsFinished() {
		lose("All threads of the program ran to the end, so the program was successful. Try to sabotage the program before it finishes.")
		return
	}
}

func win(winInfo string) {
	runState = 1
	fmt.Print(winInfo)
}

func lose(loseInfo string) {
	runState = -1
	fmt.Print(loseInfo)
}

func expandThread(threadId int) {
	saveForUndo()
	gameState.ThreadContexts[threadId].Expanded = true
}

func undo() {

	for e := undoHistory.Front(); e != nil; e = e.Next() {
		if e.Next() == nil {
			temp := e.Value.(c.History).GlobalContext
			gameState.GlobalState = &temp
			var tmpThreadCtx = []*execution.ThreadContext{}
			for _, tx := range e.Value.(c.History).ThreadContext {
				tv := &execution.ThreadContext{}
				deepcopier.Copy(tx).To(tv)
				tmpThreadCtx = append(tmpThreadCtx, tv)
			}

			gameState.ThreadContexts = tmpThreadCtx
			undoHistory.Remove(e)
		}
	}
}

func stepThread(thread int) {
	//if IsLevelPristine() {
	////第一关
	////sendEvent('Gameplay', 'level-first-step', gameState.getLevelId());
	//}
	var program = gameState.GetProgramOfThread(thread)
	var threadState = gameState.ThreadContexts[thread]
	var pc = threadState.ProgramCounter

	if !IsThreadFinished(thread) {
		saveForUndo()
		if threadState.Expanded {
			//展开了的情况
			program[pc].GetExpandInstructions()[threadState.ExpProgramCounter].Execute(gameState.GlobalState, threadState)
		} else {
			program[pc].Execute(gameState.GlobalState, threadState)
		}
		checkForVictoryConditions()

	}
	//fmt.Println(program, threadState, pc)
	fmt.Println(gameState.GlobalState)
}

func IsThreadBlocked(threadId int) bool {
	if IsThreadFinished(threadId) {
		return false
	}

	var threadState = gameState.ThreadContexts[threadId]
	var currentIns = (*threadState.Instructions)[threadState.ProgramCounter]

	return currentIns.IsBlocking(gameState.GlobalState, threadState)
}

func IsThreadFinished(threadId int) bool {
	program := gameState.GetProgramOfThread(threadId)
	var maxInstructions = len(program)
	var threadState = gameState.ThreadContexts[threadId]
	var pc = threadState.ProgramCounter

	return pc >= maxInstructions
}

func IsLevelPristine() bool {
	//return undoHistory.length == 0;
	return true
}

// View Model
type LevelInfo struct {
	Title       string        `json:"title"`
	Label       string        `json:"label"`
	Description string        `json:"description"`
	Programs    []Program     `json:"programs"`
	GameStatus  int           `json:"gameStatus"`
	Context     []ContextType `json:"context"`
	CanUndo     bool          `json:"canUndo"`
	VictoryCond string        `json:"victoryCond"`
}

func checkCanUndo() bool {
	if undoHistory.Len() > 0 {
		return true
	} else {
		return false
	}
}

type Program struct {
	Current          []int   `json:"current"`
	CanStepNext      bool    `json:"canStepNext"`
	CanCurrentExpand bool    `json:"canCurrentExpand"`
	Code             []Coder `json:"code"`
}
type Coder struct {
	Name               string  `json:"name"`
	Code               string  `json:"code"`
	Indent             int     `json:"indent"`
	Description        string  `json:"description"`
	ExpandInstructions []Coder `json:"expandInstructions"`
	Expanded           bool    `json:"expanded"`
}

type BackResponse struct {
	Status int       `json:"status"`
	Msg    string    `json:"msg"`
	Data   LevelInfo `json:"data"`
}

type ContextType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func canBeExpanded(ins []execution.Instruction) bool {
	if len(ins) > 0 {
		return true
	} else {
		return false
	}

}

func getExpandInstructions(instruction []execution.Instruction) []Coder {
	cders := []Coder{}
	for _, v := range instruction {
		o := Coder{}
		o.Name = v.GetName()
		o.Code = v.GetCode()
		o.Expanded = canBeExpanded(v.GetExpandInstructions())
		o.ExpandInstructions = nil
		o.Indent = 0
		o.Description = v.GetDescription()
		cders = append(cders, o)
	}
	return cders
}

func packageData() LevelInfo {

	levelInfo := LevelInfo{}
	levelInfo.Title = gameState.Level.Title
	levelInfo.Label = gameState.Level.Label
	levelInfo.VictoryCond = gameState.Level.VictoryCondition

	levelInfo.CanUndo = checkCanUndo()
	levelInfo.Description = gameState.Level.Description
	levelInfo.GameStatus = runState
	for _, pro := range gameState.ThreadContexts {
		p := Program{}

		for _, co := range *pro.Instructions {
			coder := Coder{}
			coder.Code = co.GetCode()
			coder.Name = co.GetName()
			coder.Description = co.GetDescription()
			coder.Indent = 0
			coder.ExpandInstructions = getExpandInstructions(co.GetExpandInstructions())
			coder.Expanded = pro.Expanded

			p.Code = append(p.Code, coder)
		}

		var temp = (*pro.Instructions)[pro.ProgramCounter]

		p.CanCurrentExpand = temp != nil && len(temp.GetExpandInstructions()) > 0
		p.CanStepNext = !IsThreadFinished(pro.Id) && !IsThreadBlocked(pro.Id)
		p.Current = []int{pro.ProgramCounter, pro.ExpProgramCounter}

		levelInfo.Programs = append(levelInfo.Programs, p)
	}

	for _, v := range gameState.GlobalState.Values {
		ctx := ContextType{}
		ctx.Name = v.Name
		ctx.Value = fmt.Sprint(v.Value)
		levelInfo.Context = append(levelInfo.Context, ctx)
	}
	//IsThreadFinished
	return levelInfo
}
func loadLevel(id int, err error) BackResponse {
	backResponse := BackResponse{}
	if err == nil {
		backResponse.Data = packageData()
		backResponse.Status = 0
		backResponse.Msg = "success"
		//正确解析
	} else {
		backResponse.Data = packageData()
		backResponse.Status = 1
		backResponse.Msg = "关卡ID无法识别"
		// 不是数字
	}
	return backResponse
}

func main() {

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(context context.Context) {
		context.ServeFile("/home/novemser/Documents/Code/Hackathon/GoLangPingCAP/dist/index.html", true)
	})

	app.StaticServe("/home/novemser/Documents/Code/Hackathon/GoLangPingCAP/dist/static", "/static")

	// 加载关卡
	app.Get("/api/level/{id}", func(ctx iris.Context) {
		runState = 0
		levelIdStr := ctx.Params().Get("id")
		levelId, err := strconv.Atoi(levelIdStr)
		startLevel(levelId)

		ctx.JSON(loadLevel(levelId, err))
	})
	// 步进代码

	app.Get("/api/stepthread/{level}/{thread}/{currentLine}", func(ctx iris.Context) {
		thread := ctx.Params().Get("thread")
		levelIdStr := ctx.Params().Get("level")
		levelId, err := strconv.Atoi(levelIdStr)
		threadId, err := strconv.Atoi(thread)
		stepThread(threadId)
		ctx.JSON(loadLevel(levelId, err))
	})

	//展开代码
	app.Get("/api/expand/{level}/{thread}/{currentLine}", func(ctx iris.Context) {
		thread := ctx.Params().Get("thread")
		levelIdStr := ctx.Params().Get("level")
		levelId, err := strconv.Atoi(levelIdStr)
		threadId, err := strconv.Atoi(thread)
		expandThread(threadId)
		ctx.JSON(loadLevel(levelId, err))
	})

	//步进展开代码
	app.Get("/api/stepExpandedThread/{level}/{thread}/{currentLine}/{currentSubLine}", func(ctx iris.Context) {
		thread := ctx.Params().Get("thread")
		levelIdStr := ctx.Params().Get("level")
		levelId, err := strconv.Atoi(levelIdStr)
		threadId, err := strconv.Atoi(thread)
		stepThread(threadId)
		ctx.JSON(loadLevel(levelId, err))
	})

	//撤销
	app.Get("/api/undo/{level}", func(ctx iris.Context) {
		levelIdStr := ctx.Params().Get("level")
		levelId, err := strconv.Atoi(levelIdStr)
		undo()
		ctx.JSON(loadLevel(levelId, err))
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
