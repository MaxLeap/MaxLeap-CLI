package main
import (
	"strings"
	"strconv"
)
type Context struct {
	cmd string
	args Args
	flags map[string]string
}

type Args struct {
	args []string
}
type Flag interface {
	getName() string
	getUsage() string
	getValue()
}

type StringFlag struct {
	Name string
	Usage string
	Value string
}
func (flag StringFlag)getValue()string  {
	return flag.Value
}
func (flag StringFlag)getName()string  {
return flag.Name
}
func (flag StringFlag)getUsage()string  {
	return flag.Usage
}
type IntFlag struct {
	Name string
	Usage string
	Value int
}
func (flag IntFlag)getValue() int {
	return flag.Value
}
func (flag IntFlag)getName()string  {
return flag.Name
}
func (flag IntFlag)getUsage()string  {
	return flag.Usage
}

func NewContext(s string) *Context{
	args:=strings.Split(s," ");
	ctx:=Context{"",Args{make([]string,0)},make(map[string]interface{},0)}
	length:=len(args)
	if length>0{
		ctx.cmd=args[0]
		if length>1{
			lastFlag:=""
			for _,item :=range(args[1:]){
				if strings.Index(item,"-")==0{
					lastFlag=item[1:]
					ctx.flags[lastFlag]=nil
				}else {
					if lastFlag==""{
						ctx.args=append(ctx.args,item)
					}else {
						ctx.flags[lastFlag]=item
						lastFlag=""
					}

				}
			}
		}
	}
	return &ctx
}
func (context Context)GetCMD() string {
	return context.cmd
}
func(context Context) Args() Args  {
	
	return context.args
}
func(context Context)String(f string)string{
	return context.flags[f]
}
func (context Context)Int(f string) (i int) {
	flag:=context.flags(f)
	if(flag!=nil){
		if iflag,err:=strconv.Atoi(flag);err==nil{
			i=iflag
		}

	}
	return nil
}
func( args Args) First() string{
	if len(args)>0{
		return args[0]
	}
	return ""
}