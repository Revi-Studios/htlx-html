# htlx

A cleaner and more readable way of writing html

## Idea:

The htlx tool will convert htlx into normal html. When building web apps primaraly on the server, this tool can convert the files on the fly

```htlx
button> Hello
div cors=true>
    li> World
    flow st=margin:20px>
        p> Things can be
        p> Good!
    stack>
        input>
        p> D
        input>
```

Converts to ->

```html
<button>Hello</button>
<div cors="true">
    <li>World</li>
    <flow style="margin:20px">
        <p>Things can be</p>
        <p>Good!</p>
    </flow>
    <stack>
        <input />
        <p>D</p>
        <input />
    </stack>
</div>
```

**Result:**

- Cleaner code
- More readable code
- ~23% less lines of code
- ~30% less characters
