package go_koans

import (
	"bytes"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
        out.WriteString(in.String())

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

        str, _ := in.ReadString(byte(' '))
        out.Write([]byte(str[:len(str)-1]))
        // fmt.Printf("-%s-", out.String())

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
