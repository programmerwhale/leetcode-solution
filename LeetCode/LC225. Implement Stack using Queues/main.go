package LC225__Implement_Stack_using_Queues

func main() {

}

type MyStack struct {
	queue1,queue2 []int
}


func Constructor() MyStack {
	return MyStack{}
}


func (this *MyStack) Push(x int)  {
	this.queue1=append(this.queue1,x)
}


func (this *MyStack) Pop() (x int) {
	if len(this.queue1)>0{
		x=this.queue1[len(this.queue1)-1]
		this.queue2=this.queue1[:len(this.queue1)-1]
	}
	this.queue1=this.queue2
	return
}


func (this *MyStack) Top() (x int) {
	if len(this.queue1)>0{
		x= this.queue1[len(this.queue1)-1]
	}
	return
}


func (this *MyStack) Empty() bool {
	if len(this.queue1)==0&&len(this.queue2)==0{
		return true
	}else{
		return false
	}
}