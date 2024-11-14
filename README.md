# Dual Server

This project demonstrates running two HTTP servers concurrently with Go routines. Each server sends periodic status messages, which are logged by the main application.

<div style="display: flex; align-items: flex-start;">
    <a href="https://blog.stackademic.com/what-are-goroutines-learn-in-3-minutes-038da2f0e200">
        <img src="https://miro.medium.com/v2/resize:fit:720/format:webp/1*eq5OHymD52p-fV_9g-Trgw.png" alt="Goroutine Schedule" style="align-self: flex-start;height: 210px;" />
    </a>
    <a href="https://blog.stackademic.com/what-are-goroutines-learn-in-3-minutes-038da2f0e200">
        <img src="https://go101.org/article/res/goroutine-schedule.png" alt="Goroutine Schedule" style="align-self: flex-end; height: 210px;" />
    </a>
</div>

## Prerequisites

- Go installed: `go version`

### Install & Run

```bash
git clone git@github.com:Stei-ITstudents/DualServer.git
cd DualServer
make run
```

- Server 1: `http://localhost:8081/server1`
- Server 2: `http://localhost:8082/server2`

#### Graceful Shutdown

To gracefully stop the servers, use Ctrl+C or send a termination signal. The application will clean up and wait for both servers to finish before exiting.

### License

MIT License
