## Breyta

Breyta is a [json-xls](http://www.json-xls.com/api) REST client.
Breyta retries if json-xls server isn't responding with exponential jitter backoff strategy.

Breyta expects a [Mashape](https://market.mashape.com/) account and the json-xls API key.
You can obtain a key from [here](https://market.mashape.com/json-xls-com/json2xls/pricing) after
subscribing to one of the plans.

#### Installation

Using `go get`

```go
go get -u -v github.com/younisshah/breyta
```

Using `govendor`
```go
govendor fetch -v github.com/younisshah/breyta
```


#### Usage

To convert a JSON string to XLSX/XLS

```go
b := breyta.NewJSONClient(breyta.FormatXlsx, breyta.Both, breyta.LayoutAuto, breyta.Both, "MASHAPE_KEY")
convertedBytes, err := b.ConvertJSON(`[{"Name":"root1","children":[{"Name":"AAA","Age":"22","Job":"PPP"},{"Name":"BBB","Age":"25","Job":"QQQ"}]},{"Name":"root2","children":[{"Name":"CCC","Age":"38","Job":"RRR"}]},{"Name":"root3","children":[]}]`)
if err != nil {
    //
}
// Do something to convertedBytes. Maybe save to a file
```

To convert an XML string to XLSX/XLS

```go
b := breyta.NewXMLClient(breyta.FormatXlsx, breyta.Both, breyta.LayoutAuto, breyta.Both, "MASHAPE_KEY")
convertedBytes, err := b.ConvertXML(`<?xml version=\"1.0\" encoding=\"UTF-8\"?><note>  <to>Tove</to>  <from>Jani</from>  <heading>Reminder</heading>  <body>Do not forget me this weekend!</body></note>`)
if err != nil {
    //
}
// Do something to convertedBytes. Maybe save to a file
```


#### TODOs

1) Add support for `ConvertJsonFile` API
2) Add support for `ConvertXmlFile` API.

#### License

MIT