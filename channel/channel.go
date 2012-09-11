package channel

import "fmt"

type Channel interface {
    Each(func(Data))
    Map(func(Data) Data) DataChannel
    String() string
    Any(func(Data) bool) bool
    All(func(Data) bool) bool
}

type Data interface{}

type DataChannel chan Data

func (c DataChannel) Each(f func(Data)) {
    for v := range c {
        f(v)
    }
}

func (c DataChannel) Map(f func(Data) Data) DataChannel {
    r := make(DataChannel)
    go func() {
        for v := range c {
            c <- f(v)
        }
        close(r)
    }()
    return r
}

func (c DataChannel) String() string {
    s = ""
    for v := range c {
        s += fmt.Sprint(v)
    }
    return s
}

type Number int64

type NumberChannel chan Number