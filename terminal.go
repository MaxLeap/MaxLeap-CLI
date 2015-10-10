package main
import(
	"github.com/chzyer/readline"

)

func start(session Session) {
	rl, err := readline.New("las>")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	cmds:=newCommands()
	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			ctx:=NewContext(line)
			cmd:=cmds.getCommand(ctx.GetCMD())
			cmd.Action(ctx)
		}
		println(line)
	}
}
