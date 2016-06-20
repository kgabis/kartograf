##About
Kartograf is a tool to convert pngs into maps for use in games. If a pixel in map's image matches a color in map's format json then a value corresponding to this color is added to output's tiles array with (x, y) values corresponding to pixel's (x, y) values. See example directory to see how map's format json looks like. 

##Installing
```
go get github.com/kgabis/kartograf
```

##Usage
```
kartograf map.png mapFormat.json 

```

##Example
From kartograf's root:
```
kartograf example/map.png example/mapFormat.json > map.json
```


##License
[The MIT License (MIT)](http://opensource.org/licenses/mit-license.php)