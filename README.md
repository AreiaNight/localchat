# localchat
![panic](https://github.com/AreiaNight/localchat/blob/main/vmware_Y8HtpjJBJj.png)

A 1-to-1 terminal chat between two computers on the same local network, no internet required. Built in Go as a learning project.

Links and text are sent as plain text, no hyperlinks, no risk of accidentally clicking anything.

## How it works

One machine runs the server and waits for a connection. The other runs the client and connects to it. Once connected, both sides can send and receive messages at the same time.

```
Machine A (server)             Machine B (client)
  localchat.go        ◄──────►   client.go
  listens on :6969               connects to A's IP
```

## Requirements

- Go installed on both machines
- Both connected to the same local network (WiFi or Ethernet)

## Usage

**1. Clone the repository on both machines**
```bash
git clone https://github.com/your-username/localchat
cd localchat
```

**2. Find the server machine's IP address**
```bash
# Windows
ipconfig

# Mac / Linux
ip addr
```
Look for an IP starting with `19x.x.x.x`

**3. Machine A — run the server**
```bash
go run localchat.go
```

**4. Machine B — edit `client.go` with A's IP address, then run**
```bash
go run client.go
```

## Project structure

```
localchat/
├── localchat.go   # server: waits for incoming connection
├── client.go      # client: connects to the server
├── go.mod
└── README.md
```

## Notes

- Default port is `6969`, can be changed in both files (I had to make the joke)
- On Windows you may need to open the port in the firewall:
```powershell
New-NetFirewallRule -DisplayName "localchat" -Direction Inbound -Protocol TCP -LocalPort 6969 -Action Allow
```
