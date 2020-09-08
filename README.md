# MTR-ArriveTime-Backend
a CLI app with [cobra](https://github.com/spf13/cobra) which can check the time of arrival of the coming 3 trains

# Usage
`go build -o <your_app> main.go` <br>
### Enter the root
`./your_app`
### List the available lines
`./your_app line`
### List the available stations of a specific line
`./your_app station -l <line_acronym>` / `./your_app station --line <line_acronym>`
### List the estimated arrive time
`./your_app time -l <line_acronym> -s <station_acronym>` / `./your_app time --line <line_acronym> --station <station_acronym>`

# Example
`go build -o mtr main.go`<br>
`./mtr`<br>
`./mtr line`<br>
`./mtr station -l AEL` / `./mtr station --line AEL`<br>
`./mtr time -l AEL -s KOW` / `./mtr time --line AEL --station KOW` <br>

# Source
[MTR API](https://data.gov.hk/en-data/dataset/mtr-data2-nexttrain-data) <br>
[JSON specification](https://opendata.mtr.com.hk/doc/Next_Train_DataDictionary_v1.1.pdf)
