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

->

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

|              | **HTML**  | **HTLX**  | **Differens** |
| :----------- | :-------- | :-------- | :------------ |
| _Lines_      | 13 lines  | 10 lines  | ~23%          |
| _Characters_ | 221 char  | 154 char  | ~30%          |
| _File Size_  | 194 bytes | 164 bytes | ~15%          |

> **Result:**
>
> _Using HTLX will result in:_
>
> - Cleaner code
> - More readable code
> - ~23% less lines of code
> - ~30% less characters
