## Stack
### Last-In-First-Out
Core Operations:
- Push - Add element to the top
- Pop - Remove element from the top
- Peek/Top - View the top element without removing it
- IsEmpty - Check if stack is empty

```go
stack := []int{}
stack = append(stack, 5)                       // push
if len(stack) > 0 {                           
    top := stack[len(stack)-1]                 // peek
}
if len(stack) > 0 {                           
    stack = stack[:len(stack)-1]               // pop
}
isEmpty := len(stack) == 0                     // isEmpty
```
