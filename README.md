GoPDQ is a persistent, disk-backed queue written in go.

Basic properties:
* fast - try to max out an SSD - try with both small bits of data and large
* lightweight
* persistent
* safe
* (relatively) easy to inspect
* optional wait for confirmation of writes (send after write+sync)
* flexible
  - maximum delay between writes
  - maximum entries between writes
  - demo how to 'subclass'
  - optional compression

use:

    q := gopdq.New()
    q.Push(val)
    q.Wait() // wait for a flush to occur
    q.Flush() // force a flush
    i := q.Pop() // maybe this?
    i.Value.(type)?
    i.Done() // mark as deletable