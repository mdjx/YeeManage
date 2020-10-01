# Yeelight Management

Windows and Linux command line tool for controlling Yeelight colour smart globes. 

## Binaries

Downloadable binaries can be found on the releases page. 

## Usage

```
yeelight.exe -ip <ip> 
    [-hsvHue <0-359>] 
    [-hsvSat <0-100>] 
    [-css <red, blue, aqua, violet, etc>] 
    [-rgb <000000-FFFFFF>] 
    [-effect <sudden|smooth>]
    [-duration <30+>]
    [-power <on|off|toggle>]
```

### Examples

Turn on a globe

`yeelight.exe -ip 192.168.1.55 -power on`

Toggle the power

`yeelight.exe --ip=192.168.1.55 --power=toggle`

Multiple commands can be used. Turn on a globe, set the color to RGB value #9400D3, and set the brightness to max (100). 

`yeelight.exe -ip 10.1.1.15 -power on -rgb 9400d3 -brightness 100`

Set colour using CSS named colours. See [here](https://css-tricks.com/snippets/css/named-colors-and-hex-equivalents/) for a list of valid values. 

`yeelight.exe -ip 10.1.1.15 -css royalblue`

Linux example showing HSV usage

`./yeelight-linux --ip=10.250.1.117 --hsvHue=120 --hsvSat=100 --effect=sudden`

## Release Notes

`1.0`
- Initial release

`1.1`
- Added CSS Named colour flag (`-css`). 