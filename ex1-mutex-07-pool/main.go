package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Overview: the use of sync.Pool is to set a constrain on the creation of expensive resources
// e.g. DB / Network connections and memory
// A fix instance of resources is created and reused using snyc.Pool, rather than create new ones
// The caller that requires the resources will call sync.Pool.Get() whcih checks whether there are existing instance
// to the resource and uses them, otherwise new instances are created
// When fininshed with the usage, the caller calls sync.Pool.Get to release the resources

var bufPool = sync.Pool{
	//sync.Pool has a New field that hold a reference to the funcation that
	// will be called if there are no instances available in the pool
	// the anonymous function allocates and returns a new instance of a bytes.Buffer
	New: func() interface{} {
		fmt.Println("allocate new buffer of bytes")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {

	// The line below is not used because, for every call to the log() function,
	// it will create a new btye.buffer, that results in alot of stale memory and
	// the application will become slow, thus affecting the performance
	// var b bytes.Buffer

	b := bufPool.Get().(*bytes.Buffer) // 2. Using a sync.Pool, get the instance from buffer pool by passing a ptr of bytes.Buffer
	b.Reset()                          // 3. Next, we reset the buffer

	b.WriteString(time.Now().Format("15:04:05")) // 3. Now, we use the buffer
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
	bufPool.Put(b) // 4. Finally, we put back the buffer back to the buffer pool
}

func main() {
	log(os.Stdout, "debug-string-1")
	log(os.Stdout, "debug-string-2") //For the second example, we are using the same buffer pool used in the 1st example
}
