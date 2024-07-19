<div align="center">
  <h1>Auto Updater</h1>
  <p>
    A simple little process replacement demo
  </p>
</div>
<br />

<!-- Table of Contents -->
# Table of Contents

- [About the Project](#about-the-project)
  * [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Roadmap](#roadmap)

<!-- About the Project -->
## About the Project
A small demo of a self updating app. At present, it monitors the folder from which the binary is executed from and upon the binary being replaced that matches its name it replaces the running process by executing the new binary. There is currently no validation of the new binary aside from a simple name check, nor are there any tests.
<br>
<br>
This is my first real Go project, and I had a genuinely great time working on it. That does mean that the project likely does not follow best Go practices and perhaps has some quirks in the coding/architecture. While choosing Go for this project was not the most efficient route for achieving a broad, "production ready", feature set, I believe it presents a good example of my capacity to pick up a new tech stack.  


<!-- TechStack -->
### Tech Stack

<details>
  <summary>Libraries</summary>
  <ul>
    <li><a href="https://github.com/fsnotify/fsnotify">FSNotify</a></li>
  </ul>
</details>


<!-- Getting Started -->
## Getting Started

<!-- Prerequisites -->
### Prerequisites

This project is written in Golang, please make sure you have Go installed


<!-- Run Locally -->
### Run Locally

Clone the project

```bash
  git clone https://github.com/Louis3797/awesome-readme-template.git
```

Go to the project directory

```bash
  cd SelfUpdate
```


<!-- Usage -->
## Usage

```bash
  go run main.go
```

Once the application is running, you may replace built binary with another of the same name to replace the running process

<!-- Roadmap -->
## Roadmap

* [ ] Generate hash on building binary as output artifact
* [ ] Validate hash before executing new binary to ensure binary is from a valid update source
* [ ] Tests
* [ ] Store keys in env and pull in on build
* [ ] Build pipeline
* [ ] Log errors to file, or external resource
* [ ] Remote deployment


<!-- Acknowledgments -->
## Acknowledgements
 - [Awesome README](https://github.com/matiassingers/awesome-readme)
 - [Readme Template](https://github.com/othneildrew/Best-README-Template)
