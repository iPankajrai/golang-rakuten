- What is a channel?
    Do not communicate by sharing memory; instead, share memory by communicating.

- Why channel?
  To establish communication between goroutines


- There is a saperate memory for communication purpose between goroutines. That is nothting but a channel
- 1. There are two kinds of channels. Buffered and Unbuffered channel
- 2. There is a sender and there is a receier.
- For Unbuffered channel
- 3. The sender gorountine is blocked until the receiver goroutine receives the value
- 4. The receier goroutine is blocked untile the sender sends the value
- 5. Either the sender or a receiver is blocked it is called deadlock.

- For Buffered channel
- 6. The sender gorountine is blocked(only when it is full) until the receiver goroutine receives the value
- 7. The receier goroutine is blocked untile the sender sends the value
- 8. Either the sender or a receiver is blocked it is called deadlock.
- 9. For buffered channel if the buffer is not full ,then neither the sender nor receiver is blocked

- <- is towards the channel then it is a sender. ch <- 100
- <- is from the channel then it is a receiver. v:= <- ch
