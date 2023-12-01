# POC code - deadlock in `database/sql` package

Tu run, simply do: 
```
docker-compose up
```
which will spin up postgres, initialize a simple db and will also build and start our PoC go code.

When you want to have a clear state, just do 
```
docker-compose down -v
```
to delete all disk states.

If you need to rebuild the docker image after a code change, you may run 
```
docker-compose build
```

### Example output

```
...
go-sql-deadlock-postgres-1  | The files belonging to this database system will be owned by user "postgres".
...
go-sql-deadlock-postgres-1  | 2023-12-01 13:05:36.227 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
go-sql-deadlock-postgres-1  | 2023-12-01 13:05:36.227 UTC [1] LOG:  listening on IPv6 address "::", port 5432
go-sql-deadlock-postgres-1  | 2023-12-01 13:05:36.228 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
go-sql-deadlock-postgres-1  | 2023-12-01 13:05:36.229 UTC [66] LOG:  database system was shut down at 2023-12-01 13:05:36 UTC
go-sql-deadlock-postgres-1  | 2023-12-01 13:05:36.231 UTC [1] LOG:  database system is ready to accept connections
go-sql-deadlock-gotest-1    | uuid [48 48 48 48 48 48 48 48 45 48 48 48 48 45 48 48 48 48 45 48 48 48 48 45 48 48 48 48 48 48 48 48 48 48 48 48]
go-sql-deadlock-gotest-1    | sleep begin
go-sql-deadlock-gotest-1    | sleep end
go-sql-deadlock-gotest-1    | err context canceled
go-sql-deadlock-gotest-1    | successful exit
go-sql-deadlock-gotest-1    | uuid [48 48 48 48 48 48 48 48 45 48 48 48 48 45 48 48 48 48 45 48 48 48 48 45 48 48 48 48 48 48 48 48 48 48 48 48]
go-sql-deadlock-gotest-1    | sleep begin
go-sql-deadlock-gotest-1    | sleep end
```

In the above example, the process was successful at the first run, but started to hang at the second run attempt. 

You may open another terminal and print the stack trace of the go process like: 
```
# docker exec -it ea9717c40426 bash
root@ea9717c40426:/go/src/app# ps aux
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0   2316  1280 ?        Ss   13:09   0:00 sh -c while true; do ./test; done
root          13  0.0  0.0 1603836 8064 ?        Sl   13:09   0:00 ./test
root          19  0.0  0.0   2316  1280 pts/0    Ss   13:10   0:00 sh
root          25  0.1  0.0   4056  3200 pts/0    S    13:10   0:00 bash
root          27  0.0  0.0   8040  3712 pts/0    R+   13:10   0:00 ps aux
root@ea9717c40426:/go/src/app# kill -s QUIT 13
```
Replace `ea9717c40426` with your container id, and `13` with your PID instance of the running `./test` process. The `SIGQUIT` signal would print the stacktrace to the stdout which you can see the in the docker-compose output. 
