# gron
gron, inspired by clockwork.




## Design Goal

- No delay job. Next schedule for a particular job shall not be delayed for its processing time. For example: Every(1).Minute() ensures tasks run at 10:01, 10:02, 10:03, but not 10:01+processing.
- every - second,minute,hour,day,week
- at (dow hh:mm)


- take away second, its too risky to have diff(job) <=1second
-https://github.com/robfig/cron/blob/32d9c273155a0506d27cf73dd1246e86a470997e/constantdelay_test.go //follow this.




## Design Specs

### CRON
An ADT that maintains a queue of entries/jobs by time (earliest).

- **Run()**. Core functionality, run indefinitely(go-routine), multiplex different channels.
- **AddSchedule(entry)**. Adds entry to queue, signals `add`.
- **Stop()**. Signals `stop` to halt cron processing.
- **Clear()**. Clear all entries from queue.


### Entry
An ADT that keep tracks of the following states: `next`, `prev`, and `job`.

-**Schedule(time)**. To schedule next run, referenced from input `time`.
-**Run()**. To run the given job (go-routine), recoverable.

### DelayedEntry
DelayedEntry is a type of entry which implements Entry.

-**Every(period)**. Create an entry with time referenced to `time.Now()`.
-**At(period)**. ignore time.Now(), create new date.