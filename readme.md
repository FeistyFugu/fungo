# Fungo

Since Go still has no generics, classical loops like Apply, Reduce, Min or Max need to be re-implemented every time you need them for a new type. Fungo makes this a little less painful. It provides fast and simple implementations of many typical functional constructs and uses 'go generate' to create type safe versions of each algorithm.

Please note this is a work in progress. All constructs PROBABLY work but a lot more testing is needed.

### Currently supported constructs

- All
- Apply
- Compare
- Contains
- FanOut
- Filter
- FindFirst
- FindLast
- GroupBy
- Max
- Min
- Reduce

### How to use

First, install Fungo:

`go get github.com/FeistyFugu/fungo`

Then make sure your Go bin folder is in your path. If it isn't use:

 `PATH=$PATH:[your GOPATH]/bin`
 
Now, at the top of any Go source file, add the required go generate statement. Here's an example:
 
`//go:generate fungo -template=Apply -fileName=apply -t1=string -t2=int`

Finally, to generate the code, open a terminal, cd to your project's root folder and enter:

`go generate ./...`

### Command line arguments:

- template: Name of the construct you want to use. Required.
- functionName: Name of the file to generate. By default, it will be the name of the template.
- packageName: Name of the package the new function will be part of. By default, it will be the name of the directory where the file is located.
- fileName: Name of the source file to generate. By default, it will be the same as the function.
- t1: Name of the first type to use. By default, it's string.
- t2: Name of the second type to use. By default, it's also string.

### Current limitations

Right now, Fungo only works with slices. It might make sense in the future to add some support for maps.