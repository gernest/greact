# goss
write css in go

# Example

-
```go
	s, _ := goss.ParseCSS(
		"",
		goss.CSS{
			"background": "blue",
		},
	)
	fmt.Println(goss.ToCSS(s, &goss.Options{}))
```
```css
background: blue;
```

- 
```go
	s, _ := goss.ParseCSS(
		"a",
		goss.CSS{
			"float": "left",
			"width": "1px",
		},
	)
    fmt.Println(goss.ToCSS(s, &goss.Options{}))
```
```css
a {
  float: left;
  width: 1px;
}
```

-

```go
	s, _ := goss.ParseCSS(
		"a",
		goss.CSS{
			"border": []string{"1px solid red", "1px solid blue"},
		},
    )
```

```css
a {
  border: 1px solid red, 1px solid blue;
}
```

### Related  projects

- [jss](https://github.com/cssinjs/jss)