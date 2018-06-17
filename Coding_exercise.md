# VAT Test

## Background
Utility Warehouse offer gas, electricity, broadband, mobile phone contracts, landline contracts and cashback cards to their customers. It's possible they may expand to offer other services in future.
When pricing usage we need to select the **correct** VAT rate for any one of these services **when** the utility was consumed. For this reason we maintain a RESTful refdata service that will return the standard
and reduced VAT rates over time. e.g.

```
GET /api/1.0/views/uk-vat-rates
```

```
[
  {
    "start":"2011-01-04T00:00:00Z",
    "end":null,
    "data":{
      "standard":0.2,
      "reduced":0.05
    }
  },
  {
    "start":"2010-01-01T00:00:00Z",
    "end":"2011-01-04T00:00:00Z",
    "data":{
      "standard":0.175,
      "reduced":0.05
    }
  },
  {
    "start":"2008-12-01T00:00:00Z",
    "end":"2010-01-01T00:00:00Z",
    "data":{
      "standard":0.15,
      "reduced":0.05
    }
  },
  {
    "start":"1997-08-31T23:00:00Z",
    "end":"2008-12-01T00:00:00Z",
    "data":{
      "standard":0.175,
      "reduced":0.05
    }
  },
  {
    "start":"1994-04-01T00:00:00Z",
    "end":"1997-08-31T23:00:00Z",
    "data":{
      "standard":0.175,
      "reduced":0.08
    }
  },
  {
    "start":"1991-03-19T00:00:00Z",
    "end":"1994-04-01T00:00:00Z",
    "data":{
      "standard":0.175,
      "reduced":0.00
    }
  },
  {
    "start":"1979-06-18T00:00:00Z",
    "end":"1991-03-19T00:00:00Z",
    "data":{
      "standard":0.15,
      "reduced":0.0
    }
  }
]
```
## The Exercise
Write a GO application which given a date/time and VAT type (standard or reduced) will output the correct VAT rate, e.g.

```
$ go run vat.go standard 2017-01-01T00:00:00Z
The standard rate of VAT for 2017-01-01 00:00:00 +0000 UTC is 20%

$ go run vat.go reduced 2017-01-01T00:00:00Z
The reduced rate of VAT for 2017-01-01 00:00:00 +0000 UTC is 5%
```

You can run the refdata service on your pairing station by typing...
```
docker run -p8080:8080 catalinilea/vat
```
or use the instance running in development
```
http://localhost:8080/api/1.0/views/uk-vat-rates
```