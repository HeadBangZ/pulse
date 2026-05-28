# Pulse

Pulse is a gRPC based log ingestor designed to securely receive and queue logs from distributed services (That is the plan!)

## Prerequisites
To work on this project, ensure you have the following installed:
- Go (Version 1.25.4 or higher)
- Protocol Buffers Compiler (protoc)
- Task (The task runner for building and managing the project)

## Setup
1. Clone the repository
```bash
git clone https://github.com/HeadBangZ/pulse.git
cd pulse
```

2. Initialize dependencies
Run the following command to install required Go modules
```bash
task init
```

## Task Commands
The project uses ```Taskfile.yaml``` to manage build, generation, and execution workflows.
| Command | Description |
| :--- | :--- |
| `task init` | Installs dependencies and runs `go mod tidy`. |
| `task proto` | Generates Go source code from the `.proto` contract. |
| `task build:windows` | Compiles the Windows binary (`.exe`). |
| `task build:linux` | Compiles the Linux binary. |
| `task build:all` | Compiles binaries for both Windows and Linux. |
| `task run:windows` | Builds and executes the ingestor on Windows. |
| `task run:linux` | Builds and executes the ingestor on Linux. |

## Project Structure
- ```/api```: Contains the ```.proto``` definitions and generated gRPC/protobuf code.
- ```/cmd```: Contains the entry point for the ingestor service.

