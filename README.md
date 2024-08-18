# insights
A slow query log analyzer tool based on llm.

```mermaid
block-beta
columns 1

ollama

prompt<["prompt"]>(up)  

Collector
space
block:logs
loki
logfile["log file"]
end

Collector -- "collect" --> loki
Collector -- "collect" --> logfile

log<["log"]>(up)  

block:Databases
MongoDB[("MongoDB")]
MySQL[("&nbsp;&nbsp;&nbsp;MySQL&nbsp;&nbsp;")]    OtherDatabases[("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;...&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;")]
end

```

## TODO

- [x] Command Line Integration
- [ ] Ollama Integration
- [ ] Add support for MongoDB
- [ ] Add support for MySQL
- [ ] Loki Integration


# 🤝 Thanks
This tool is built with the help of the following projects:
- [ollama](https://github.com/ollama/ollama)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [cobra](github.com/spf13/cobra)
