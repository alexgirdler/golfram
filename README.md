# GOLFRAM
A Wolfram Alpha API lib wrote in go.

## Instalation
Simply Import "github.com/alexgirdler/golfram"

## Example Usage
```go
c := golfram.NewClient("APP-KEY")
result, err := c.NewQuery("plot sinx")
if err != nil {
  c.GetPlot(result, 0, "")
}
```

This example will query Wolfram Alpha with input "plot sinx" and will save the graph to "imgs/plot_sinx"(default save path, can be changed by passing path to the third argument in GetPlot)

## Contributing
Please feel free to contribute simply checkout a feature branch commit your changes and open a pull request.

## License
Copyright (c) 2015 Alex Girdler

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
