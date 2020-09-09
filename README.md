# MTR-ArriveTime-Backend
a CLI app with [cobra](https://github.com/spf13/cobra) which can check the time of arrival of the coming 3 trains

# Installation
`go get -u github.com/spf13/cobra/cobra`

# Usage
### Build your app
`go build -o <your_app> main.go` <br>
### Enter the root
`./your_app`
### List the available lines
`./your_app line`
### List the available stations of a specific line
`./your_app station -l <line_acronym>` / `./your_app station --line <line_acronym>`
### List the estimated arrive time
`./your_app time -l <line_acronym> -s <station_acronym>` / `./your_app time --line <line_acronym> --station <station_acronym>`
### Help
`./your_app -h` / `./your_app --help` <br>
`./your_app line -h` / `./your_app line --help`<br>
`etc.`

# Example
`go build -o mtr main.go`<br>
`./mtr`<br>
`./mtr line`<br>
`./mtr station -l AEL` / `./mtr station --line AEL`<br>
`./mtr time -l AEL -s KOW` / `./mtr time --line AEL --station KOW` <br>
`./mtr -h` / `./mtr --help`<br>
`./mtr line -h` / `./mtr line --help`<br>

# Source
[MTR API](https://data.gov.hk/en-data/dataset/mtr-data2-nexttrain-data) <br>
[JSON specification](https://opendata.mtr.com.hk/doc/Next_Train_DataDictionary_v1.1.pdf)
